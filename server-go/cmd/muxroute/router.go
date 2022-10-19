package muxroute

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/ark-go/bake-go/cmd/aredis"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

var (
	Router *mux.Router
)

func StartRouter() *mux.Router {
	Router = mux.NewRouter()
	Router.Use(sessionMiddleware)
	Router.HandleFunc("/api/products", Products).Methods("POST")
	Router.HandleFunc("/api/{.*}", func(w http.ResponseWriter, r *http.Request) {
		// проксируем эти запросы на Node Js
		sess := r.Context().Value(aredis.SessionKey("sessKey")).(*aredis.SessionJson) // получаем указатель на структуру сессии
		color.Set(color.FgYellow)
		log.Println(r.RequestURI+"\t>>", sess.User.Email, sess.IPAddress, sess.UserBrowser.Timezone)
		color.Unset()
		//color.Yellow(r.RequestURI+"\t>> %s %s %s", sess.User.Email, sess.IPAddress, sess.UserBrowser.Timezone)
		reverseProxy(w, r) // проксируем запрос на NodeJs
	}).Methods("POST")
	Router.HandleFunc("/socket.io/", func(w http.ResponseWriter, r *http.Request) {
		// json.NewEncoder(w).Encode(map[string]bool{"ok": true})
		reverseProxy(w, r)
	})
	// объявим пути для файлов SPA если файла нет в каталоге, то выдается index.html
	spa := spaHandler{staticPath: "/home/arkadii/Projects/bake/quasar/dist/spa", indexPath: "/index.html"}
	Router.PathPrefix("/").Handler(spa)
	return Router
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

// // spaHandler реализует интерфейс http.Handler, поэтому мы можем его использовать
// // для ответа на HTTP-запросы. Путь к статической директории и
// // путь к индексному файлу в этом статическом каталоге используется для
// // обслуживаем SPA в заданном статическом каталоге.
// type spaHandler struct {
// 	staticPath string
// 	indexPath  string
// }
