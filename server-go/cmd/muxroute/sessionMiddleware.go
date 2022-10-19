package muxroute

import (
	"log"
	"net/http"
	"strings"

	"github.com/ark-go/bake-go/cmd/aredis"
)

func sessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var store *aredis.RedisStore = nil // подготовим переменную
		var err error
		// пропускаем запуск сессии там где нам не надо
		if !containsUri(r.RequestURI, []string{"/socket.io/", "/assets/", "/images/"}) {
			store, err = aredis.StartSession(r, aredis.ClientRedis)
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
func containsUri(str string, s []string) bool {
	for _, v := range s {
		if strings.HasPrefix(str, v) {
			return true
		}
	}
	return false
}
