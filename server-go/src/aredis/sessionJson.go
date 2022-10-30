package aredis

import (
	//"net/http"
	"time"
)

// https://mholt.github.io/json-to-go/
type SessionJson struct {
	IdRedis string   `json:"idRedis"`
	IsNew   bool     `json:"isNew"`
	Cookie  struct { // ark, обязательная структура
		OriginalMaxAge int       `json:"originalMaxAge"`
		Expires        time.Time `json:"expires"`
		Secure         bool      `json:"secure"`
		HTTPOnly       bool      `json:"httpOnly"`
		Path           string    `json:"path"`
		SameSite       bool      `json:"sameSite"`
	} `json:"cookie"`
	User struct {
		ID             string    `json:"id"`
		Email          string    `json:"email"`
		Active         bool      `json:"active"`
		UFam           string    `json:"u_fam"`
		UName          string    `json:"u_name"`
		UOtch          string    `json:"u_otch"`
		BlockedDate    time.Time `json:"blockedDate"`
		BlockedMessage string    `json:"blockedMessage"`
		Roles          []string  `json:"roles"`
		Permissions    struct {
		} `json:"permissions"`
		Meta struct {
		} `json:"meta"`
		Telegram    string    `json:"telegram"`
		CreateAt    time.Time `json:"createAt"`
		Description string    `json:"description"`
		TreeID      string    `json:"tree_id"`
		Rereg       bool      `json:"rereg"`
		UserDate    time.Time `json:"user_date"`
		UserID      string    `json:"user_id"`
		Status      string    `json:"status"`
		Fio         string    `json:"fio"`
		Sid         string    `json:"sid"`
	} `json:"user"`
	CreatedAt   int64  `json:"createdAt"`
	Timezone    string `json:"timezone"`
	UserBrowser struct {
		Name     string `json:"name"`
		Version  string `json:"version"`
		Os       string `json:"os"`
		Type     string `json:"type"`
		Timezone string `json:"timezone"`
	} `json:"userBrowser"`
	IPAddress string `json:"ipAddress"`
}

// func getSession(req *http.Request) SessionJson {
// 	session, err := store.Get(req, "BreadAndTandoor") // cookie name
// }
