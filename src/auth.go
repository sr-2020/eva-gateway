package main

import (
	"github.com/julienschmidt/httprouter"
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
