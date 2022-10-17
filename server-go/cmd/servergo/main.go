package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ark-go/bake-go/cmd/redisstore"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var clientRedis *redis.Client

func init() {
	// hsh := hmac.New(sha256.New, []byte("Bread&Tandoor%bakehouse_516")) // secret
	// hsh.Write([]byte("uoTdPV7vvEkXmv-3NGcsFN4KXA0DWiFC"))              // есть в куке
	// b64 := base64.StdEncoding.EncodeToString(hsh.Sum(nil))             // выдаст хеш  совпадает с JS но обнулить пару сиволовов - https://github.com/tj/node-cookie-signature/blob/master/index.js
	// sampleRegexp := regexp.MustCompile(`\=+$`)
	// result := sampleRegexp.ReplaceAllString(b64, "") //  совпадает с JS !!!
	// log.Println("b64:", result)
	//log.Println(hex.EncodeToString(hsh.Sum(nil)))
}

// spaHandler реализует интерфейс http.Handler, поэтому мы можем его использовать
// для ответа на HTTP-запросы. Путь к статической директории и
// путь к индексному файлу в этом статическом каталоге используется для
// обслуживаем SPA в заданном статическом каталоге.
type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP проверяет URL-адрес, чтобы найти файл в статическом каталоге
// в обработчике SPA. Если файл найден, он будет обслуживаться. Если нет, то
// файл, расположенный по пути индекса, в продолжении staticPath, в обработчике SPA, будет обслуживаться. Этот
// подходит для обслуживания SPA (одностраничного приложения).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// получить абсолютный путь, чтобы предотвратить обход каталога
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// если нам не удалось получить абсолютный путь, отвечаем 400 неверным запросом
		// и стоп
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// добавьте к пути путь к статическому каталогу
	path = filepath.Join(h.staticPath, path)
	// проверить, существует ли файл по указанному пути
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// Файла не существует, подставим индексный файл
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// если мы получили ошибку (не то, что файл не существует) с указанием
		// файл, вернуть внутреннюю ошибку сервера 500 и остановить
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// в противном случае используйте http.FileServer для обслуживания статического каталога
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

// Пересылам через прокси на другой сервер
func reverseProxy(w http.ResponseWriter, r *http.Request) {
	// для прокси указываем только путь к прокси, остальной url подставляется из запроса
	targ := fmt.Sprintf("http://%s", "127.0.0.1:8877")
	url, _ := url.Parse(targ) // преобразуем в тип url.URL
	//log.Println("в проксю:", url)
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(w, r)
}

// Поиск совпадений начала Uri со строками в массивe
// если есть совпадение вернет true
func containsUri(str string, s []string) bool {
	for _, v := range s {
		if strings.HasPrefix(str, v) {
			return true
		}
	}
	return false
}

// Здесь запрашиваем сессию
func sessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var store *redisstore.RedisStore = nil // подготовим переменную
		var err error
		// пропускаем запуск сессии там где нам не надо
		if !containsUri(r.RequestURI, []string{"/socket.io/", "/assets/", "/images/"}) {
			store, err = startSession(r, clientRedis)
			if err != nil {
				log.Println(err.Error())
			}
		}

		// Вызовите следующий обработчик, которым может быть другое ПО промежуточного слоя в цепочке, или последний обработчик.
		next.ServeHTTP(w, r)
		if store != nil {
			// если был запущен Store то сохраним при выходе
			store.Save(r, w)
		}
		//log.Println("Выход............")
	})
}
func main() {
	// TODO я делаю клиента общим !?
	clientRedis = redis.NewClient(&redis.Options{
		Addr: os.Getenv("Redis_host") + ":" + os.Getenv("Redis_port"),
	})
	router := mux.NewRouter()
	// Запуск сессии
	router.Use(sessionMiddleware)
	// Роуты
	router.HandleFunc("/api/{.*}", func(w http.ResponseWriter, r *http.Request) {
		// проксируем эти запросы на Node Js
		sess := r.Context().Value(redisstore.SessionKey("sessKey")).(*redisstore.SessionJson) // получаем указатель на структуру сессии
		log.Println(r.RequestURI, "\t>>", sess.User.Email, sess.IPAddress, sess.UserBrowser.Timezone)
		reverseProxy(w, r) // проксируем запрос на NodeJs
	})
	// socket.io пересылаем молча
	router.HandleFunc("/socket.io/", func(w http.ResponseWriter, r *http.Request) {
		// json.NewEncoder(w).Encode(map[string]bool{"ok": true})
		reverseProxy(w, r)
	})
	// объявим пути для файлов SPA если файла нет в каталоге, то выдается index.html
	spa := spaHandler{staticPath: "/home/arkadii/Projects/bake/quasar/dist/spa", indexPath: "/index.html"}
	router.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler: router,
		Addr:    ":8018",
		//Хорошая практика: устанавливайте тайм-ауты для серверов, которые вы создаете!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Старт сервера:", srv.Addr)
	// запуск сервера
	log.Fatal(srv.ListenAndServe())

}
