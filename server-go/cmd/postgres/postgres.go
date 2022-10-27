package postgres

import (
	"context"
	"log"
	"os"

	"github.com/ark-go/bake-go/cmd/utils"
	"github.com/fatih/color"
	"github.com/jackc/pgx/v5/pgxpool"
)

var PG Pg

type Pg struct {
	Pool *pgxpool.Pool
}

// размер базы
func (pg *Pg) sizeDB(nameDB string) int64 {
	var sz int64
	err := pg.Pool.QueryRow(context.Background(), "select "+"pg_database_size('"+nameDB+"')").Scan(&sz)
	if err != nil {
		log.Println("QueryRow failed: ", err)
		return 0
	}
	return sz
}

// подключение к Pg
func StartPostgres() {
	PG = Pg{}
	//log.Println(os.Getenv("PG_DatabaseStr"))

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("PG_DatabaseStr"))
	if err != nil {
		log.Println("Не удалось создать пул соединений: ", err.Error())
		os.Exit(1)
	}
	PG.Pool = dbpool

	sizeDB := PG.sizeDB(dbpool.Config().ConnConfig.Database)
	color.Set(color.FgGreen)
	log.Println("PG подключились, DB Name:", dbpool.Config().ConnConfig.Config.User, utils.ByteCountSI(sizeDB))
	color.Unset()
}

// закроем подключение к Pg
func (pg *Pg) Close() {
	pg.Pool.Close()
}
