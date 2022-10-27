package products

import (
	//	"github.com/ark-go/bake-go/cmd/aredis"
	"github.com/ark-go/bake-go/cmd/iface"
)

type Load iface.TabQueryT

func (p *Load) MakeQuery(param map[string]any) (otv *iface.ReturnQueryT) {
	otv = &iface.ReturnQueryT{}
	where := ""
	otv.ValuesSql = []any{p.Session.Timezone}
	if len(p.RowReturning) > 0 { // если требуется изменить запрос в зависимости от предыдущих запросов
		if id, ok := p.RowReturning["id"]; ok { // если есть требуемый параметр
			otv.ValuesSql = append(otv.ValuesSql, id) // добавляем параметр
			where = "WHERE products.id = $2"          // нас интересует только одна строка
		} else {
			where = ""
		}
	}

	otv.TextSql = `--sql
	SELECT
	products.id,
	products.name,
	products.massa,
	-- products.unit_id,
	-- unit.name AS unit_name,
	unit.name AS unit_name,
	producttype.prefix AS producttype_prefix,
	(select count(*) from productingred where productingred.products_id = products.id) AS count_ingredients,
	products.productvid_id,
	concat(productassortment.name,' ',productvid.name,' ',productvid.nameext) AS productvid_name,
	productassortment.prefix AS prefix,
	products.description,
	products.document_num, -- TTK
	-- products.document_date,
	to_char(products.document_date,  'DD.MM.YYYY') as document_date,
	products.article_buh,
	'PR-' ||products.article AS article,
	-- -- -- -- 
	users.email AS "user_email",
	to_char(products.user_date at time zone $1,  'DD.MM.YYYY HH12:MI:SS') as "user_date",
	products.user_date as "user_date_test",
	products.meta
	FROM products
	LEFT JOIN  users ON users.id = products.user_id
	
	LEFT JOIN  productvid ON productvid.id = products.productvid_id
	LEFT JOIN  productassortment ON productassortment.id = productvid.productassortment_id
	LEFT JOIN  producttype ON producttype.id = productassortment.producttype_id
	LEFT JOIN  unit ON unit.id = productvid.unit_id
	-- LEFT JOIN  productassortment ON productassortment.id = productvid.productassortment_id
	` + where + `
	ORDER BY productassortment.name, productvid.name,productvid.nameext, products.name
	
	`
	return

}
