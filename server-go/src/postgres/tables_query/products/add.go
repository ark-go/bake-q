package products

import (
	"github.com/ark-go/bake-go/src/iface"
)

type Add iface.TabQueryT

// type Add struct {
// 	Session  *aredis.SessionJson
// 	InParams map[string]any // входные параметры для запроса
// }

func (p *Add) MakeQuery(param map[string]any) (otv *iface.ReturnQueryT) {
	otv = &iface.ReturnQueryT{}
	otv.TextSql = `--sql
INSERT INTO products (name, massa,productvid_id,description,document_num,document_date,article_buh, user_id, user_date)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,CURRENT_TIMESTAMP)
RETURNING *;
`
	otv.ValuesSql = []any{
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
