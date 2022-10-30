package postgres

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/ark-go/bake-go/src/aredis"
	"github.com/ark-go/bake-go/src/iface"
	"github.com/ark-go/bake-go/src/postgres/tables_query"
	"github.com/jackc/pgx/v5"
)

type ApiHandler struct{}

type DbMakeResponse struct {
	InParams     *map[string]any  // входные параметры для запроса
	RowReturning []map[string]any // для результата RETURNING
	session      *aredis.SessionJson
	timeZone     string
}

// декодируем входные параметры
func (q *DbMakeResponse) DecodeInParams(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(&q.InParams); err == nil { // входные параметры
		log.Println("Входные параметры:")
		for l, m := range *q.InParams {
			log.Println(l, " : ", m)
		}
	} else {
		return err
	}

	return nil
}

// Выводим результат Returning в картy target
func (q *DbMakeResponse) ReturningQueryToMap(result []byte, target *map[string]any) error {
	m := []map[string]any{}
	err := json.Unmarshal(result, &m) // входные параметры
	*target = m[0]
	//log.Println(">>>>>>", target, ">>>>", string(result))
	return err
}

// проверка наличия ID если есть то это Update
func (p *DbMakeResponse) GetIdAddOrUpdate() string {
	_, ok := (*p.InParams)["id"] // проверяем есть ли ключ
	if ok {
		return "update"
	} else {
		return "add"
	}
}

// инициализация структуры для таблиц
func (p *DbMakeResponse) InitTableController() *iface.TableContrlT {

	t := &iface.TableContrlT{
		Session:             p.session,
		InParams:            *p.InParams,
		ExecQuery:           p.ExecQuery,
		GetIdAddOrUpdate:    p.GetIdAddOrUpdate,
		ReturningQueryToMap: p.ReturningQueryToMap,
	}
	if t.Session.Timezone == "" {
		t.Session.Timezone = "Europe/Moscow"
	}
	return t
}

// вход хандлера запроса, т.е. старт
func (pTmp *ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := &DbMakeResponse{}                             // создаем структуру, вроде только для этого потока ServeHTTP как бы старт, заход в поток
	w.Header().Set("Content-Type", "application/json") // Для ответа, отвечаем всегда и ошибки свои тоже
	//log.Println("wwwwwwwwwwwwwwww")
	p.getSession(w, r)                          // Запросим сессию
	if err := p.DecodeInParams(r); err != nil { // получим входные параметры
		p.SendResponse(w, nil, errors.New("не распознать параметры (go)"))
		return // не распознали входные параметры
	}
	// считаем что CMD есть всегда и в нем есть команда
	_, exists := (*p.InParams)["cmd"]            // определим наличие cmd
	tabname, exists2 := (*p.InParams)["tabname"] // определим tabname
	if !exists || !exists2 {
		// нет ожидаемого ключа cmd? должны вернуть ошибку
		p.SendResponse(w, nil, errors.New("не хватает данных в запросе (go)"))
		return

	}

	switch tabname {

	case `products`:
		tab_products := (*query.Tab_products)(p.InitTableController()) // создаем структуру для таблицы
		if resQuery, err := tab_products.CreateQuery(); err == nil {   // создаем SQL запрос
			result, err := p.ExecQuery(resQuery) // выполнение запроса вернет []byte запроса из структуры
			//log.Println("<><>", string(result))
			p.SendResponse(w, result, err) // отправка результата клиенту, ошибки зависят от запроса ранее
		} else {
			p.SendResponse(w, nil, err)
		}
	}
	//p.SendResponse(w, nil, errors.New("Нет командыы"))
}

// var d int = 8

// отправка клиенту результата или ошибки
func (p *DbMakeResponse) SendResponse(w http.ResponseWriter, textResult []byte, err error) {
	// TODO не возвращаем ничего?
	mess := ""
	if err != nil {
		mess = `{"error":"` + err.Error() + `"}` // TODO странно про кавычки ?? они тут нужны но почему?
	} else {
		mess = `{"result":` + string(textResult) + "}"
	}
	//w.WriteHeader(http.StatusOK)
	//log.Println("Отправка клиенту---->\n", mess)
	fmt.Fprintf(w, "%s", mess)
}

// вытаскиваем сессию из запроса // TODO организрвать проверку доступа
func (p *DbMakeResponse) getSession(w http.ResponseWriter, r *http.Request) {
	p.session = r.Context().Value(aredis.SessionKey("sessKey")).(*aredis.SessionJson) // получаем указатель на структуру сессии
	p.timeZone = p.session.Timezone
	log.Println("GO:", r.RequestURI+"\t>>", p.session.User.Email, p.session.IPAddress, p.session.UserBrowser.Timezone)

}

// отправка запроса на сборку SQL, возвращаем текст запроса и параметры SQL для запроса, на основе входщих параметров
func (p *DbMakeResponse) MakeQuery(param iface.CmdRequest) (textSql string, valuesSql []any) {
	textSql, valuesSql = param.MakeQuery(*p.InParams) //    вернуть параметры.. мы не используем структуру в этом коде.. только все делаем там а возвращаем через функции
	return
}

// выполнить запрос к базе данных
func (p *DbMakeResponse) ExecQuery(otv *iface.ReturnQueryT) ([]byte, error) {
	textSql := otv.TextSql
	valuesSql := otv.ValuesSql
	rows, err := PG.Pool.Query(context.Background(), textSql, valuesSql...)
	//defer rows.Close() // Обязательно не забывать!
	if err != nil {
		// TODO кончились слоты (default 100) в самом постгрессе мы вылетели тут.. ВАЖНО: оставшиеся слоты подключений зарезервированы для подключений суперпользователя
		log.Fatal("ошибка при выполнении запроса ", err.Error())
		return nil, err
	}
	defer rows.Close()
	result, err := pgx.CollectRows(rows, pgx.RowToMap)
	if err != nil {
		log.Println("ошибка преобразовния данных в map: ", err.Error())
		return nil, err
	}
	var b []byte
	var errMap error
	// если нам надо вставить ключ в JSON так чтобы полученный массив был по ключу
	if otv.AddPrefixJson != "" { // проверим был ли установлен ключ JSON для результата
		// вставляем наш ключ, и получаем новый map, старый идет значением ключа
		resultMap := map[string]any{otv.AddPrefixJson: result}
		// переводим в JSON из map
		b, errMap = json.Marshal(resultMap)
		if errMap != nil {
			log.Println("Ошибка map to json-1 ", errMap)
			return nil, errMap
		}
	} else {
		// чистый ответ массив map
		b, errMap = json.Marshal(result)
		if errMap != nil {
			log.Println("Ошибка map to json-2 ", errMap)
			return nil, errMap
		}
	}
	//log.Println(string(b))
	return b, nil
}
