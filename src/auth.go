package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type AuthUser struct {
	Id        int    `json:"id"`
	Admin     bool   `json:"admin"`
	Status    string `json:"status"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

var authUser AuthUser

func GetAuth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Proxy(r, "http://auth.evarun.ru/api/v1/profile", &authUser, nil)
}

func ProfileMiddleware(w http.ResponseWriter, r *http.Request, ps httprouter.Params) string {
	Auth(r)
	return "/api/v1" + ps.ByName("path")
}

func AuthService(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := "/api/v1" + ps.ByName("path")

	switch path {
	case "/profile":
		path = ProfileMiddleware(w, r, ps)
	default:
		//
	}

	var resp interface{}
	res, err := ProxyLite(r, cfg.Auth + path, &resp)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(res.StatusCode)
	response, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(response))
}
