package postgres

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func init() {
	PG = Pg{}
}

var PG Pg

type Pg struct {
	Pool *pgxpool.Pool
}

// размер базы
func (pg *Pg) sizeDB(nameDB string) string {
	var sz string
	err := pg.Pool.QueryRow(context.Background(), "select "+"pg_database_size('"+nameDB+"')").Scan(&sz)
	if err != nil {
		log.Println("QueryRow failed: ", err)
		return ""
	}
	return sz
}

// подключение к Pg
func (pg *Pg) Connect() {
	log.Println(os.Getenv("PG_DatabaseStr"))

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("PG_DatabaseStr"))
	if err != nil {
		log.Println("Не удалось создать пул соединений: ", err.Error())
		os.Exit(1)
	}
	pg.Pool = dbpool

	sizeDB := pg.sizeDB(dbpool.Config().ConnConfig.Database)
	//
	log.Println("PG подключились", dbpool.Config().ConnConfig.Config.User, sizeDB)

}

// закроем подключение к Pg
func (pg *Pg) Close() {
	pg.Pool.Close()
}
