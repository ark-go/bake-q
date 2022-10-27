package iface

import (
	"log"
	"net/http"

	"github.com/ark-go/bake-go/cmd/aredis"
)

// интерфейс для сборки запроса
type CmdRequest interface {
	MakeQuery(map[string]any) (string, []any) // собираем вщзвращаем запрос для БД
}
type IDbMakeResponse interface {
	MakeQuery(param CmdRequest) (textSql string, valuesSql []any)
	ExecQuery(otv *ReturnQueryT) ([]byte, error)
	SendResponse(w http.ResponseWriter, textResult []byte, err error) error
	DecodeInParams(r *http.Request) error
}

type TabQueryT struct {
	Session      *aredis.SessionJson // текущая сессия , User Timezone и возможно роли
	InParams     map[string]any      // входные параметры для полного запроса
	RowReturning map[string]any      // для результата предыдущего запроса RETURNING если требуется для одной строки после Update
	// ExecQuery            func(otv *ReturnQueryT) ([]byte, error)
	// ReturningQueryToMap func(result []byte, target *map[string]any) error
	// GetIdAddOrUpdate     func() string
}
type TableContrlT struct {
	Session  *aredis.SessionJson // текущая сессия , User Timezone и возможно роли
	InParams map[string]any      // входные параметры для полного запроса
	//	RowReturning map[string]any      // для результата предыдущего запроса RETURNING если требуется для одной строки после Update
	// выполнить запрос к базе данных
	ExecQuery           func(otv *ReturnQueryT) ([]byte, error)
	ReturningQueryToMap func(result []byte, target *map[string]any) error
	GetIdAddOrUpdate    func() string
}

func (tq *TabQueryT) Test(r string) {
	log.Println("testttesttesttest", r)
}

type ReturnQueryT struct {
	TextSql       string
	ValuesSql     []any
	AddPrefixJson string
}
