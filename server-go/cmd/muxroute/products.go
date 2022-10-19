package muxroute

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Products(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["name"]
	var incoming struct {
		Cmd     string
		Tabname string
	}
	json.NewDecoder(r.Body).Decode(&incoming)
	res2B, _ := json.Marshal(incoming)
	log.Println(params, string(res2B))
}
