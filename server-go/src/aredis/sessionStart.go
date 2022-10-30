package aredis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"os"
)

func StartSession(req *http.Request, client redis.UniversalClient) (*RedisStore, error) {
	//client := redis.NewUniversalClient(&redis.UniversalOptions{
	// client := redis.NewClient(&redis.Options{
	// 	Addr: os.Getenv("Redis_host") + ":" + os.Getenv("Redis_port"),
	// })
	// New default RedisStore
	store, err := NewRedisStore(context.Background(), client)
	if err != nil {
		log.Fatal("failed to create redis store: ", err)
		return nil, err
	}

	// Установим префикс ключ сессии, префикс для Redis
	//store.KeyPrefix(os.Getenv("Redis_ProjectPrefix") + ":sess:")

	// Запросим сессию
	session, err := store.Get(req, os.Getenv("COOKIE_NAME")) // cookie name BreadAndTandoor
	if err != nil {
		log.Println("не удалось получить сессию: ", err)
		return nil, err
	}
	//log.Println("sesss", session)
	// Получим сессию и положим ее в контекст
	//if nodeSession, ok := session.Values["Data"].(redisstore.SessionJson); ok {
	// засунем сессию в контекст
	var sessKey SessionKey = "sessKey" // необходим свой тип string
	ctx := req.Context()
	r := req.WithContext(context.WithValue(ctx, sessKey, session)) // в контекст суем указатель
	*req = *r
	// session.User.Email = "Arkadii@yandex.ru" // TODO тест изменения, по ссылке
	return store, nil
	//}

	//return errors.New("не удалось прочитать сессию")
}
