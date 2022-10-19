package main

import (
	"github.com/ark-go/bake-go/cmd/muxroute"
	"github.com/ark-go/bake-go/cmd/postgres"
	"log"
	"net/http"
	"time"
)

func main() {
	// подключаемся к базе
	postgres.PG.Connect()
	defer postgres.PG.Close()
	postgres.PG.GetUnivirsal("recept")
	log.Println("--------------------------")
	postgres.PG.GetUnivirsal("region")
	muxroute.StartRouter()

	srv := &http.Server{
		Handler: muxroute.Router,
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
