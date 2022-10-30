package postgres

import (
	"context"
	"log"
	"os"

	"github.com/ark-go/bake-go/src/utils"
	"github.com/fatih/color"
	"github.com/jackc/pgx/v5"
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
func (pg *Pg) StartPostgres() {

	config, err := pgxpool.ParseConfig(os.Getenv("PG_DatabaseStr"))
	if err != nil {
		log.Println("Ошибка парсинга конфига pgxpool")
		return
	}
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		// do something with every new connection
		log.Println("Коннект из пула : ", config.MinConns, config.MaxConns)
		return nil
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Println("Не удалось создать пул соединений: ", err.Error())
		os.Exit(1)

	}
	pg.Pool = dbpool
	PG = *pg // TODO  убрать тут
	//

	sizeDB := pg.sizeDB(dbpool.Config().ConnConfig.Database)
	color.Set(color.FgGreen)
	log.Println("Pg подключились, DB Name:", dbpool.Config().ConnConfig.Config.User, utils.ByteCountSI(sizeDB))
	color.Unset()
}

// закроем подключение к Pg
func (pg *Pg) Close() {
	pg.Pool.Close()
}
