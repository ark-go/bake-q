package postgres

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

func init() {

}

type recept struct {
	Id         *int64     `json:"id"`
	NomenclId  *int64     `json:"nomencl_id"`
	Article    *string    `json:"article"`
	Name       *string    `json:"name"`
	DateRecept *time.Time `json:"date_recept"`
	Text       *string    `json:"text"`
	Doc        *int64     `json:"doc"`
	UserId     *int64     `json:"user_id"`
	UserDate   *time.Time `json:"user_date"`
	Deleted    *bool      `json:"deleted"`
	Extend     *int64     `json:"extend"`
	Meta       *any       `json:"meta"`
}

func (pg *Pg) GetRecept() {
	query, value := getQuery1()
	rows, err := pg.Pool.Query(context.Background(), query, value...)
	if err != nil {
		log.Fatal("error while executing query ", err.Error())
	}
	defer rows.Close()
	recepts, err := pgx.CollectRows(rows, pgx.RowToStructByPos[recept])
	if err != nil {
		log.Println("CollectRows error: ", err.Error())
		return
	}

	for _, p := range recepts {
		log.Println(p.Name, p.DateRecept, p.UserDate)
	}
	// переводим в JSON NULL нормально проходит, за счет указателей в структуре
	b, err := json.Marshal(recepts)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(b))

	// for rows.Next() {
	// 	values, err := rows.Values()
	// 	if err != nil {
	// 		log.Fatal("error while iterating dataset")
	// 	}
	// 	log.Println("-->", values[0], values[4].(time.Time)) //, values[12], "id", values[17])
	// }
	// // Any errors encountered by rows.Next or rows.Scan will be returned here
	// if rows.Err() != nil {
	// 	log.Println("Err:", err)
	// }
}

func getQuery1() (query string, value []any) {
	query = /*sql*/ `
select * from recept  where id>$1
`
	value = []any{1}
	return
}
