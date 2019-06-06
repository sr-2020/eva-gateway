package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type AuthUser struct {
	Id        int    `json:"id"`
	Admin     bool   `json:"admin"`
	Status    string `json:"status"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func AuthMiddleware(w http.ResponseWriter, r *http.Request, ps httprouter.Params) string {
	Auth(r)
	return ps.ByName("path")
}

func Auth(request *http.Request) error {
	return AuthRequest(request, cfg.Auth + "/api/v1/profile", nil)
}

func AuthRequest(request *http.Request, url string, data interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", request.Header.Get("Authorization"))

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var response interface{}
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Print(jsonErr)
		return jsonErr
	}

	var authUser AuthUser
	Decode(&authUser, response)

	request.Header.Set("X-User-Id", strconv.Itoa(authUser.Id))

	return nil
}
