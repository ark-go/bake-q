package productvid

import "github.com/ark-go/bake-go/src/iface"

type Load iface.TabQueryT

func (p *Load) MakeQuery(param map[string]any) (otv *iface.ReturnQueryT) {
	otv = &iface.ReturnQueryT{
		AddPrefixJson: "productvid", // Дополнтельный ключ для JSON
	}

	otv.TextSql = `--sql
	SELECT
	productvid.id,
	concat(productassortment.name,' ',productvid.name,' ',productvid.nameext) AS name
	--name
	FROM productvid
	LEFT JOIN  productassortment ON productassortment.id = productvid.productassortment_id
	ORDER BY name
	`
	otv.ValuesSql = nil
	return
}
