package main

import (
	"github.com/julienschmidt/httprouter"
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

//func NewAuthMiddleware(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	Auth(r)
//}

func NewAuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		Auth(r)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func AuthMiddleware(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (*http.Response, error) {
	Auth(r)
	return nil, nil
}

func NewLoginMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var authLogin AuthLogin
		if err := getBodyToInterface(&r.Body, &authLogin); err != nil {
			log.Println(err)
		}

		if authLogin.FirebaseToken == "" {
		}

		var authToken AuthUserToken
		res, err := ProxyData(r, &authToken)
		if err != nil {
			log.Println(err)
		}

		if err := setInterfaceToBody(authToken, &res.Body); err != nil {
			log.Println(err)
		}

		if res.StatusCode == 200 {
			token := PushToken{
				Id: authToken.Id,
				Token: authLogin.FirebaseToken,
			}

			if err := setInterfaceToBody(token, &r.Body); err != nil {
				log.Println(err)
			}

			r.Method = "PUT"
			var pushResp interface{}
			if err := ProxyOld(r, cfg.Push + "/save_token", &pushResp); err != nil {
				log.Println(err)
			}
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func LoginMiddleware(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (*http.Response, error) {

	var authLogin AuthLogin
	if err := getBodyToInterface(&r.Body, &authLogin); err != nil {
		log.Println(err)
		return nil, err
	}

	if authLogin.FirebaseToken == "" {
		return nil, nil
	}

	var authToken AuthUserToken
	res, err := ProxyData(r, &authToken)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err := setInterfaceToBody(authToken, &res.Body); err != nil {
		log.Println(err)
		return nil, err
	}

	if res.StatusCode == 200 {
		token := PushToken{
			Id: authToken.Id,
			Token: authLogin.FirebaseToken,
		}

		if err := setInterfaceToBody(token, &r.Body); err != nil {
			log.Println(err)
			return nil, err
		}

		r.Method = "PUT"
		var pushResp interface{}
		if err := ProxyOld(r, cfg.Push + "/save_token", &pushResp); err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return res, nil
}

func Auth(request *http.Request) error {
	return AuthRequest(request, cfg.Auth + "/api/v1/profile", nil)
}

func AuthRequest(request *http.Request, url string, data interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	req.Header.Set("Authorization", request.Header.Get("Authorization"))

	var authUser AuthUser
	_, err = ProxyData(req, &authUser)
	if err != nil {
		log.Println(err)
		return err
	}

	request.Header.Set("X-User-Id", strconv.Itoa(authUser.Id))

	return nil
}
