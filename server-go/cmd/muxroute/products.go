package muxroute

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"

// 	"github.com/ark-go/bake-go/cmd/aredis"
// 	"github.com/ark-go/bake-go/cmd/postgres"
// 	//"github.com/gorilla/mux"
// )

// func Products(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	sess := r.Context().Value(aredis.SessionKey("sessKey")).(*aredis.SessionJson) // получаем указатель на структуру сессии
// 	log.Println("GO:", r.RequestURI+"\t>>", sess.User.Email, sess.IPAddress, sess.UserBrowser.Timezone)
// 	//postgres.Product = &postgres.PgDataStruct{}
// 	postgres.Product.Session = sess
// 	postgres.Product.Timezone = sess.Timezone
// 	json.NewDecoder(r.Body).Decode(&postgres.Product.InParams) // входные параметры
// 	for l, m := range postgres.Product.InParams {
// 		log.Println(l, " : ", m)
// 	}

// 	_, err := postgres.Product.LoadFromDatabase(w)
// 	if err != nil {
// 		var outcoming struct {
// 			Error string `json:"error"`
// 		}
// 		outcoming.Error = err.Error()
// 		err := json.NewEncoder(w).Encode(&outcoming)
// 		if err != nil {
// 			log.Fatalln("Произошла ошибка кодирования инициализированной структуры")
// 		}
// 	}
// 	//k := postgres.Pu
// 	//params := mux.Vars(r)["name"] // Vars возвращает переменные маршрута для текущего запроса, если они есть
// 	// структура которая должна прийти
// 	// var incoming struct {
// 	// 	Cmd     string `json:"cmd"`
// 	// 	Tabname string `json:"tabname"`
// 	// }
// 	// // структура к отправке
// 	// var outcoming struct {
// 	// 	Error string `json:"error"`
// 	// }
// 	// // // отправим структуру
// 	// outcoming.Error = "Го не настроен... " + sess.User.Email
// 	// // засунем запрос в структуру
// 	// err := json.NewDecoder(r.Body).Decode(&incoming)
// 	// if err != nil {
// 	// 	log.Println("ошибка получения данных, ")
// 	// }
// 	// resB, _ := json.Marshal(incoming)
// 	// log.Println(params, string(resB))

// 	// err := json.NewEncoder(w).Encode(&outcoming)
// 	// if err != nil {
// 	// 	log.Fatalln("Произошла ошибка кодирования инициализированной структуры")
// 	// }
// }

// type productJson struct {
// 	Id                int64       `json:"id"`
// 	Name              string      `json:"name"`
// 	Massa             string      `json:"massa"`
// 	UnitName          string      `json:"unit_name"`
// 	ProducttypePrefix string      `json:"producttype_prefix"`
// 	CountIngredients  float64     `json:"count_ingredients"`
// 	ProductvidId      int64       `json:"productvid_id"`
// 	ProductvidName    string      `json:"productvid_name"`
// 	Prefix            string      `json:"prefix"`
// 	Description       string      `json:"description"`
// 	DocumentNum       string      `json:"document_num"`
// 	DocumentDate      string      `json:"document_date"`
// 	ArticleBuh        string      `json:"article_buh"`
// 	Article           string      `json:"article"`
// 	UserEmail         string      `json:"user_email"`
// 	UserDate          string      `json:"user_date"`
// 	Meta              interface{} `json:"meta"`
// }
// type productArr []productJson

// let wher = idOne ? "WHERE " + tabname + ".id = $2" : "";
