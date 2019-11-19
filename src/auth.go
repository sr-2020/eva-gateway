package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

type AuthUser struct {
	Id        int         `json:"id"`
	Amount    int         `json:"amount"`
	Followers []int       `json:"followers"`
	Following []int       `json:"following"`
	Admin     bool        `json:"admin"`
	Status    string      `json:"status"`
	Role      string      `json:"role"`
	Items     interface{} `json:"items"`
	Name      string      `json:"name"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
}

type ProfileUser struct {
	Id                int               `json:"id"`
	Amount            int               `json:"amount"`
	Followers         []int             `json:"followers"`
	Following         []int             `json:"following"`
	Admin             bool              `json:"admin"`
	Status            string            `json:"status"`
	Role              string            `json:"role"`
	Items             interface{}       `json:"items"`
	Name              string            `json:"name"`
	CreatedAt         string            `json:"created_at"`
	UpdatedAt         string            `json:"updated_at"`
	LocationUpdatedAt string            `json:"location_updated_at"`
	LocationId        *int              `json:"location_id"`
	Location          *PositionLocation `json:"location"`
}

type AuthLogin struct {
	Email         string `json:"email"`
	Password      string `json:"password"`
	FirebaseToken string `json:"firebase_token"`
}

type AuthUserToken struct {
	Id     int    `json:"id"`
	ApiKey string `json:"api_key"`
}

type PushToken struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

func (pu *ProfileUser) join(u AuthUser, p PositionUser) {
	pu.Id = u.Id
	pu.Admin = u.Admin
	pu.Status = u.Status
	pu.Name = u.Name
	pu.CreatedAt = u.CreatedAt
	pu.UpdatedAt = u.UpdatedAt

	pu.Amount = u.Amount
	pu.Followers = u.Followers
	pu.Following = u.Following
	pu.Role = u.Role
	pu.Items = u.Items

	pu.LocationUpdatedAt = p.UpdatedAt
	pu.LocationId = p.LocationId
	pu.Location = p.Location
}

func AuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if err := Auth(r); err != nil {
			log.Println(err)
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func LoginMiddleware(next http.Handler) http.Handler {
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
				Id:    authToken.Id,
				Token: authLogin.FirebaseToken,
			}

			if err := setInterfaceToBody(token, &r.Body); err != nil {
				log.Println(err)
			}

			req, err := http.NewRequest(http.MethodPut, "", nil)
			if err != nil {
				log.Println(err)
			}

			req.Header = r.Header
			req.Body = r.Body
			var pushResp interface{}
			if err := ProxyOld(req, cfg.Push+"/save_token", &pushResp); err != nil {
				log.Println(err)
			}
		}

		w.WriteHeader(res.StatusCode)
		responseBody, err := json.Marshal(authToken)
		if err != nil {
			ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		fmt.Fprint(w, string(responseBody))
	}

	return http.HandlerFunc(fn)
}

func Auth(request *http.Request) error {
	return AuthRequest(request, cfg.Auth+"/api/v1/profile", nil)
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

func GetProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var positionUser PositionUser
	var authUser AuthUser

	if err := Auth(r); err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err := ProxyOld(r, cfg.Auth+"/api/v1/profile", &authUser); err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err := ProxyOld(r, cfg.Position+"/api/v1/users/"+strconv.Itoa(authUser.Id), &positionUser); err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	temp := ProfileUser{}
	temp.join(authUser, positionUser)

	response, err := json.Marshal(temp)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(response))
}
