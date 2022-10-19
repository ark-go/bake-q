package postgres

import (
	"context"
	"log"
	"time"
)

func (pg *Pg) GetUsers() {
	query, value := getQuery()
	rows, err := pg.Pool.Query(context.Background(), query, value...)
	if err != nil {
		log.Fatal("error while executing query ", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
		log.Println("-->", values[0], values[1].(time.Time)) //, values[12], "id", values[17])
		// convert DB types to Go types
		// id := values[0].(int32)
		// firstName := values[1].(string)
		// lastName := values[2].(string)
		// dateOfBirth := values[3].(time.Time)

		// log.Println("[id:", id, ", first_name:", firstName, ", last_name:", lastName, ", date_of_birth:", dateOfBirth, "]")
	}
	// Any errors encountered by rows.Next or rows.Scan will be returned here
	if rows.Err() != nil {
		log.Println("Err:", err)
	}
}

func getQuery() (query string, value []any) {
	query = /*sql*/ `
select * from price  where id>$1
`
	value = []any{1}
	return
}
