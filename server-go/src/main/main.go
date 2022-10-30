package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ark-go/bake-go/src/aredis"
	"github.com/ark-go/bake-go/src/muxroute"
	"github.com/ark-go/bake-go/src/postgres"
)

func main() {
	// подключаемся к базе
	aredis.StartRedis()
	Pg := &postgres.Pg{}
	Pg.StartPostgres()
	defer Pg.Close()

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
	// log.Fatal(srv.ListenAndServe())
	// Запуск сервера в горутине
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println("Сервер: ", err)
		}
	}()

	c := make(chan os.Signal, 1)
	// сработает на  SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) Не сработает.
	signal.Notify(c, os.Interrupt)

	//Блокировать, пока мы не получим наш сигнал.
	<-c
	var wait time.Duration = time.Second * 15 //
	// Установите крайний срок для ожидания.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Не блокирует, если нет подключений, в противном случае будет ждать
	// до истечения срока ожидания.
	if err := srv.Shutdown(ctx); err != nil { // TODO не ловим ошибку ?
		log.Println("Error Shutdown", err.Error())
	}
	// При желании вы можете запустить srv.Shutdown в горутине и заблокировать
	// <-ctx.Done(), если ваше приложение должно ждать других служб
	// для завершения на основе отмены контекста.
	log.Println("Сервер остановлен")
	os.Exit(0)

}
