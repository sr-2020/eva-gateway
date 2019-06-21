package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type AuthUser struct {
	Id        int    `json:"id"`
	Admin     bool   `json:"admin"`
	Status    string `json:"status"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AuthLogin struct {
	Email         string `json:"email"`
	Password      string `json:"password"`
	FirebaseToken string `json:"firebase_token"`
}

type AuthUserToken struct {
	Id        int    `json:"id"`
	ApiKey    string `json:"api_key"`
}

type  PushToken struct {
	Id        int    `json:"id"`
	Token     string  `json:"token"`
}

func AuthMiddleware(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (*http.Response, error) {
	Auth(r)
	return nil, nil
}

func LoginMiddleware(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (*http.Response, error) {

	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var authLogin AuthLogin
	jsonErr := json.Unmarshal(body, &authLogin)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	r.Body = ioutil.NopCloser(strings.NewReader(string(body)))
	if authLogin.FirebaseToken == "" {
		return nil, nil
	}

	var resp interface{}
	res, err := Proxy(r, &resp)
	if err != nil {
		log.Fatal(err)
	}

	bodyResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	res.Body = ioutil.NopCloser(strings.NewReader(string(bodyResp)))

	if res.StatusCode == 200 {
		var authToken AuthUserToken
		Decode(&authToken, resp)

		token := PushToken{
			Id: authToken.Id,
			Token: authLogin.FirebaseToken,
		}

		bodyRequest, err := json.Marshal(token)
		if err != nil {
			log.Fatal(err)
		}

		t := strings.NewReader(string(bodyRequest))
		rc := ioutil.NopCloser(t)
		r.Body = rc

		r.Method = "PUT"
		var pushResp interface{}
		ProxyOld(r, cfg.Push + "/save_token", &pushResp, nil)
	}

	return res, nil
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
