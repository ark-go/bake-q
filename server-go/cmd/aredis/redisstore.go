package aredis

import (
	//"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"

	//"log"
	"net/http"
	"net/url"
	"time"

	_ "github.com/ark-go/bake-go/cmd/utils"
	"github.com/go-redis/redis/v8"
)

var ClientRedis *redis.Client

func init() {
	log.Println("Redis адрес: ", os.Getenv("Redis_host")+":"+os.Getenv("Redis_port"))
	ClientRedis = redis.NewClient(&redis.Options{
		Addr: os.Getenv("Redis_host") + ":" + os.Getenv("Redis_port"),
	})
}

type SessionKey string

// RedisStore stores gorilla sessions in Redis
type RedisStore struct {
	// client to connect to redis
	client redis.UniversalClient
	// key prefix with which the session will be stored
	keyPrefix string
	// key generator
	keyGen KeyGenFunc
	// session
	session *SessionJson
}

// KeyGenFunc defines a function used by store to generate a key
type KeyGenFunc func() (string, error)
type JsonSerializer struct{}

// NewRedis Store returns a new RedisStore with default configuration
func NewRedisStore(ctx context.Context, client redis.UniversalClient) (*RedisStore, error) {
	rs := &RedisStore{
		client:    client,
		keyPrefix: os.Getenv("Redis_ProjectPrefix") + ":sess:",
		keyGen:    generateRandomKey,
		session:   &SessionJson{},
	}
	return rs, rs.client.Ping(ctx).Err() // создаем, пингуем, отдаем
}

// получим, если есть, куку с браузером
func (s *RedisStore) GetBrouserCookie(r *http.Request, session *SessionJson) {
	c, err := r.Cookie("browser")
	if err != nil {
		return
	}
	br, err := url.QueryUnescape(c.Value) // возможные символы escape
	if err != nil {
		log.Println("Ошибка ID Escape из Cookie", err)
		return
	}
	err2 := json.Unmarshal([]byte(br), &session.UserBrowser) // Засунем в структуру
	if err2 != nil {
		log.Println("error:", err2)
		return
	}
	session.IPAddress = r.Header.Get("X-Forwarded-For")
	// log.Println("addddddd", r.Header.Get("X-Forwarded-For"))

}

// Get возвращает сеанс для данного имени после его добавления в реестр.
// либо возвращаем сессию по умолчанию - пустую
func (s *RedisStore) Get(r *http.Request, name string) (*SessionJson, error) {
	// создадим session пустую
	session := &SessionJson{}
	//s.GetBrouserCookie(r, session) // запросим куку браузера
	// запросим нашу куку
	c, err := r.Cookie(name)

	if err != nil {
		// нет куки, сеанс пустой
		// сгенерируем новый ключ для Redis? записи тут не должно быть
		key, err := s.keyGen()
		if err != nil {
			// не получилось сгенерировать ключ
			return session, err
		}
		session.IdRedis = key
		return session, nil
	}

	// запросим код из куки
	code, err := checkIdSession(c.Value)

	if err != nil {
		s.session.IdRedis = code
		if code == "BadSigned" {
			// Подпись подделана, надо сбросить удалить запись в Redis
			s.delete(r.Context(), session)
			return nil, err
		}
		if code == "BadCode" {
			// возможно код не наш
			return nil, err
		}
		return nil, errors.New("ошибка определения кода сессии из куки")
	}
	// поместим код сессии из куки без подписи в сессию
	session.IdRedis = code
	// код получили, запросим данные из redis, и заполним нашу сессию
	err = s.load(r.Context(), session)
	// данные или прочитаны из Redis, или осталась новая сессия
	s.GetBrouserCookie(r, session) // запросим куку браузера, в любом случае и засунем ее в сессию
	if err == nil {
		// укажем что это не новая сессия
		session.IsNew = false
	} else if err == redis.Nil { // Nil ключ не существует, ответ Redis-a
		// нет данных в Redis, говорим все хорошо, но сессия новая
		session.IsNew = true // TODO
		err = nil
	}

	// вернем сессию, либо пустую либо полученную из Redis
	return session, err
}

func (s *RedisStore) NewCookie(name, value string, session *SessionJson) *http.Cookie {
	// надо подписать здесь ID
	var valueSigned string
	var err error
	if value != "" {
		valueSigned, err = getSignedKey(value)
		if err != nil {
			log.Println("Ошибка при подписании ID сессии")
		}
	}

	cookie := &http.Cookie{
		Name:  name,
		Value: valueSigned,
		Path:  session.Cookie.Path,
		//Domain:   session.Cookie,
		MaxAge:   session.Cookie.OriginalMaxAge,
		Secure:   session.Cookie.Secure,
		HttpOnly: session.Cookie.HTTPOnly,
		SameSite: http.SameSite(http.SameSiteStrictMode),
	}
	if session.Cookie.OriginalMaxAge > 0 {
		d := time.Duration(session.Cookie.OriginalMaxAge) * time.Second
		cookie.Expires = time.Now().Add(d)
	} else if session.Cookie.OriginalMaxAge < 0 {
		// Set it to the past to expire now.
		cookie.Expires = time.Unix(1, 0)
	}
	return cookie
}

// Если Options.MaxAge сеанса <= 0, то файл сеанса будет
// удалено из магазина. С помощью этого процесса он обеспечивает правильное
// обработка файлов cookie сеанса, поэтому не нужно доверять управлению файлами cookie в
// веб-браузер.
func (s *RedisStore) Save(r *http.Request, w http.ResponseWriter) error {
	// получим текущую сессию в контексте если есть
	session := r.Context().Value(SessionKey("sessKey")).(*SessionJson)
	//log.Println("ses", session)
	if session.Cookie.OriginalMaxAge <= 0 {
		// удаляем куку и запись в Redis если max-age is <= 0
		if err := s.delete(r.Context(), session); err != nil {
			return err
		}
		// там эта кука поставится на удаление
		http.SetCookie(w, s.NewCookie(os.Getenv("COOKIE_NAME"), "", session))
		return nil
	}

	if session.IdRedis == "" {
		// если нет Id Redis
		id, err := s.keyGen()
		if err != nil {
			return errors.New("redisstore: не удалось сгенерировать идентификатор сеанса")
		}
		session.IdRedis = id
	}
	if err := s.saveRedis(r.Context(), session); err != nil {
		return err
	}

	http.SetCookie(w, s.NewCookie(os.Getenv("COOKIE_NAME"), session.IdRedis, session))
	return nil
}

// Key Prefix sets the key prefix to store session in Redis
func (s *RedisStore) KeyPrefix(keyPrefix string) {
	s.keyPrefix = keyPrefix
}

// KeyGen sets the key generator function
func (s *RedisStore) KeyGen(f KeyGenFunc) {
	s.keyGen = f
}

// Close closes the Redis store
func (s *RedisStore) Close() error {
	return s.client.Close()
}

// saveRedis writes session in Redis
func (s *RedisStore) saveRedis(ctx context.Context, session *SessionJson) error {

	b, err := s.Serialize(session)
	if err != nil {
		return err
	}
	return s.client.Set(ctx, s.keyPrefix+session.IdRedis, string(b), time.Duration(session.Cookie.OriginalMaxAge)*time.Second).Err()
}

// читаем из Redis
func (s *RedisStore) load(ctx context.Context, session *SessionJson) error {
	// прочитаем значение
	cmd := s.client.Get(ctx, s.keyPrefix+session.IdRedis)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	// получим значение Redis = Json-string
	b, err := cmd.Bytes()
	if err != nil {
		return err
	}
	// что расшифруем положим в сессию
	return s.Deserialize(b, session)
}

// удаляем из Redis
func (s *RedisStore) delete(ctx context.Context, session *SessionJson) error {
	return s.client.Del(ctx, s.keyPrefix+session.IdRedis).Err()
}

// Express Js хочет json в строке, сделаем тут
func (rd *RedisStore) Serialize(s *SessionJson) ([]byte, error) {
	// нашу структуру в Json - чтоб ее понимал Express Session
	b, err := json.Marshal(s)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	return b, nil
}

// разворачиваем данные в структуру session
func (rd *RedisStore) Deserialize(d []byte, s *SessionJson) error {
	err := json.Unmarshal(d, &s) // Засунем в структуру
	if err != nil {
		log.Println("error:", err)
		return err
	}
	return nil
}
