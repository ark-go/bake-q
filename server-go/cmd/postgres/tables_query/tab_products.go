package query

import (
	"errors"

	"github.com/ark-go/bake-go/cmd/iface"
	"github.com/ark-go/bake-go/cmd/postgres/tables_query/products"
	"github.com/ark-go/bake-go/cmd/postgres/tables_query/productvid"
)

type Tab_products iface.TableContrlT

func (t *Tab_products) CreateQuery() (*iface.ReturnQueryT, error) {
	cmd, exists := (t.InParams)["cmd"]
	if !exists {
		return nil, errors.New("не правильный запрос")
	}
	if cmd == "add" {
		cmd = t.GetIdAddOrUpdate() // проверка наличия ID если есть то это Update
	}
	switch cmd {
	case `load`:
		pr := &products.Load{Session: t.Session, InParams: t.InParams}
		otv := pr.MakeQuery(t.InParams) // подготовка запроса, проверка параметров уточнение
		return otv, nil
	case `update`:
		pr := &products.Update{Session: t.Session, InParams: t.InParams}
		otv := pr.MakeQuery(t.InParams)
		result, err := t.ExecQuery(otv) // выполнить запрос update
		if err != nil {
			return nil, err
		} else {
			// запрашиваем load после обновления //TODO придумать как работать с одной строкой без чтения всего заново это во front
			pr := &products.Load{Session: t.Session, InParams: t.InParams}
			t.ReturningQueryToMap(result, &pr.RowReturning) // поместим результат в структуру, для передачи предыдущего запроса точнее ID
			otv := pr.MakeQuery(t.InParams)                 // подготовка запроса, проверка параметров уточнение
			//result, err := pr.ExecQuery(otv.TextSql, otv.ValuesSql) // выполнение запроса вернет []byte запроса из структуры
			return otv, nil
		}
	case `add`:
		pr := &products.Add{Session: t.Session, InParams: t.InParams}
		otv := pr.MakeQuery(t.InParams)
		result, err := t.ExecQuery(otv) // выполнить запрос update
		if err != nil {
			// возвращаем ошибку клиенту
			return nil, err
		} else {
			// запрашиваем load после добавления //TODO придумать как работать с одной строкой без чтения всего заново это во front
			pr := &products.Load{Session: t.Session, InParams: t.InParams}
			t.ReturningQueryToMap(result, &pr.RowReturning) // поместим результат в структуру, для передачи предыдущего запроса точнее ID
			otv := pr.MakeQuery(t.InParams)                 // подготовка запроса, проверка параметров уточнение
			return otv, nil
			// result, err := p.ExecQuery(textSql, valuesSql) // выполнение запроса вернет []byte запроса из структуры
			// p.SendResponse(w, result, err)                 // отправка результата клиенту, ошибки зависят от запроса ранее

		}
	case `allSprav`:
		pr := &productvid.Load{Session: t.Session, InParams: t.InParams}
		otv := pr.MakeQuery(t.InParams) // подготовка запроса, проверка параметров уточнение
		return otv, nil
	default:
		return nil, errors.New("запрос не определен")
	}

}
