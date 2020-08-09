package middleware

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sr-2020/eva-gateway/app/adapter/client"
	"github.com/sr-2020/eva-gateway/app/adapter/config"
	"github.com/sr-2020/eva-gateway/app/adapter/presenter"
	"github.com/sr-2020/eva-gateway/app/adapter/service"
	"github.com/sr-2020/eva-gateway/app/adapter/support"
	"github.com/sr-2020/eva-gateway/app/entity"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func AuthMiddleware(next http.Handler) http.Handler {
	authService := service.Services["auth"]

	fn := func(w http.ResponseWriter, r *http.Request) {
		if err := AuthRequest(r, authService.Host + "/api/v1/profile", nil); err != nil {
			log.Println(err)

			// Set CORS headers
			support.SetCorsHeaders(w, r.Header.Get("origin"))

			pr := presenter.NewJson()
			pr.Write(w, struct{}{}, 401)
			return
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func AuthWithAnonymousMiddleware(next http.Handler) http.Handler {
	authService := service.Services["auth"]

	fn := func(w http.ResponseWriter, r *http.Request) {
		if err := AuthRequest(r, authService.Host + "/api/v1/profile", nil); err != nil {
			log.Println(err)
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func LoginMiddleware(next http.Handler) http.Handler {
	authService := service.Services["auth"]
	pushService := service.Services["push"]

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

func AuthRequest(request *http.Request, url string, data interface{}) error {
	httpClient := client.NewHttpClient(&http.Client{
		Timeout: time.Second * 10,
	})

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	authTokenHeader := strings.Split(request.Header.Get("Authorization"), " ")
	authToken := ""
	if len(authTokenHeader) == 2 {
		authToken = authTokenHeader[1]
	} else if len(authTokenHeader) == 1 {
		authToken = authTokenHeader[0]
	}

	if len(authToken) > 0 && len(authToken) <= 40 {
		req.Header.Set("Authorization", authToken)

		var authUser entity.AuthUser
		resp, err := httpClient.ProxyData(req, &authUser)
		if err != nil {
			log.Println(err)
			return err
		}
		defer resp.Body.Close()

		if authUser.Id == 0 {
			return errors.New("Unauthorize")
		}

		request.Header.Set("X-User-Id", strconv.Itoa(authUser.Id))
		return nil
	}

	if len(authToken) == 0 {
		authCookie, err := request.Cookie("Authorization")
		if err != nil {
			fmt.Println(err)
		}
		if authCookie == nil {
			return fmt.Errorf("Authorization cookie is invalid or empty.")
		}
		authToken = authCookie.Value
	}

	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		hmacSecret, err := base64.StdEncoding.DecodeString(config.Cfg.JwtSecret)
		if err != nil {
			return nil, fmt.Errorf("JWT Secret is invalid: %v", err)
		}

		return hmacSecret, nil
	})

	if token == nil {
		return fmt.Errorf("JWT Token is invalid")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if modelId, ok := claims["modelId"].(float64); ok {
			request.Header.Set("X-User-Id", strconv.Itoa(int(modelId)))
			return nil
		}
	}

	return fmt.Errorf("JWT Token is invalid")
}
