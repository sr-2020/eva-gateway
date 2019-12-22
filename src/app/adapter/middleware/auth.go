package middleware

import (
	"github.com/sr-2020/eva-gateway/app/adapter/client"
	"github.com/sr-2020/eva-gateway/app/adapter/presenter"
	"github.com/sr-2020/eva-gateway/app/adapter/support"
	"github.com/sr-2020/eva-gateway/app/entity"
	"log"
	"net/http"
	"strconv"
	"time"
)

func AuthMiddleware(next http.Handler) http.Handler {
	authService := Services["auth"]

	fn := func(w http.ResponseWriter, r *http.Request) {
		if err := AuthRequest(r, authService.Host + "/api/v1/profile", nil); err != nil {
			log.Println(err)
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func LoginMiddleware(next http.Handler) http.Handler {
	authService := Services["auth"]
	pushService := Services["push"]

	fn := func(w http.ResponseWriter, r *http.Request) {
		var authLogin entity.AuthLogin
		if err := support.GetBodyToInterface(&r.Body, &authLogin); err != nil {
			log.Println(err)
		}

		if authLogin.FirebaseToken == "" {
		}

		var authToken entity.AuthUserToken
		res, err := authService.Client.ProxyData(r, &authToken)
		if err != nil {
			log.Println(err)
		}

		if err := support.SetInterfaceToBody(authToken, &res.Body); err != nil {
			log.Println(err)
		}

		if res.StatusCode == 200 {
			token := entity.PushToken{
				Id:    authToken.Id,
				Token: authLogin.FirebaseToken,
			}

			if err := support.SetInterfaceToBody(token, &r.Body); err != nil {
				log.Println(err)
			}

			req, err := http.NewRequest(http.MethodPut, "", nil)
			if err != nil {
				log.Println(err)
			}

			req.Header = r.Header
			req.Body = r.Body
			var pushResp interface{}
			if err := pushService.Client.ProxyOld(req, pushService.Host+"/save_token", &pushResp); err != nil {
				log.Println(err)
			}
		}

		pr := presenter.NewJson()
		pr.Write(w, authToken, res.StatusCode)
	}

	return http.HandlerFunc(fn)
}

//func Auth(request *http.Request) error {
//	return AuthRequest(request, cfg.Auth+"/api/v1/profile", nil)
//}

func AuthRequest(request *http.Request, url string, data interface{}) error {
	httpClient := client.NewHttpClient(&http.Client{
		Timeout: time.Second * 10,
	})

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	req.Header.Set("Authorization", request.Header.Get("Authorization"))

	var authUser entity.AuthUser
	_, err = httpClient.ProxyData(req, &authUser)
	if err != nil {
		log.Println(err)
		return err
	}

	request.Header.Set("X-User-Id", strconv.Itoa(authUser.Id))

	return nil
}

