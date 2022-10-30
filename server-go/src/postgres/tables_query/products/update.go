package products

import (
	"github.com/ark-go/bake-go/src/iface"
)

type Update iface.TabQueryT

// type Update struct {
// 	Session  *aredis.SessionJson
// 	InParams map[string]any // входные параметры для запроса
// }

func (p *Update) MakeQuery(param map[string]any) (otv *iface.ReturnQueryT) {
	otv = &iface.ReturnQueryT{}
	otv.TextSql = `--sql
	UPDATE products SET
	name = $2,
	massa = $3,
	productvid_id = $4,
	description = $5,
	document_num = $6,
	document_date = $7,
	article_buh = $8,
	user_id = $9,
	user_date = CURRENT_TIMESTAMP
	WHERE "id" = $1
	RETURNING *;
`
	otv.ValuesSql = []any{
		param["id"],
		param["name"],
		param["massa"],
		param["productvid_id"],
		param["description"],
		param["document_num"],
		param["document_date"],
		param["article_buh"],
		p.Session.User.ID,
	}
	return
}
