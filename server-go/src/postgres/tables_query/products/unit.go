package products

import "github.com/ark-go/bake-go/src/aredis"

type Unit struct {
	Session  *aredis.SessionJson
	InParams map[string]any // входные параметры для запроса
}

func (p *Unit) MakeQuery(param map[string]any) (textSql string, valuesSql []any) {
	textSql = `--sql
	SELECT
	id,
	name
	FROM unit
	ORDER BY name
	`
	valuesSql = nil
	return
}
