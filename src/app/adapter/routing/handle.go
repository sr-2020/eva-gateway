package routing

import (
	"github.com/go-redis/redis/v7"
	"github.com/julienschmidt/httprouter"
	"github.com/sr-2020/eva-gateway/app/adapter/presenter"
	"github.com/sr-2020/eva-gateway/app/adapter/service"
	"github.com/sr-2020/eva-gateway/app/adapter/support"
	"github.com/sr-2020/eva-gateway/app/entity"
	"github.com/sr-2020/eva-gateway/app/usecases"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func GetConfig(pr presenter.Interface, client *redis.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		support.SetCorsHeaders(w, r.Header.Get("origin"))

		key := ps.ByName("key")

		val, err := client.Get(key).Result()
		if err == redis.Nil {
			_ = pr.Write(w, struct{}{}, http.StatusNotFound)
			return
		} else if err != nil {
			_ = pr.Write(w, err, http.StatusInternalServerError)
			return
		}

		_ = pr.WriteRaw(w, []byte(val), http.StatusOK)
	}
}

func SetConfig(pr presenter.Interface, client *redis.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		support.SetCorsHeaders(w, r.Header.Get("origin"))

		key := ps.ByName("key")

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			_ = pr.Write(w, err, http.StatusBadRequest)
			return
		}

		value := string(body)
		if err := client.Set(key, value, 0).Err(); err != nil {
			_ = pr.Write(w, err, http.StatusInternalServerError)
			return
		}

		_ = pr.WriteRaw(w, []byte(value), http.StatusCreated)
	}
}

func GetVersion(pr presenter.Interface) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		_ = pr.Write(w, "0.0.2", http.StatusOK)
	}
}

func GetUsers(pr presenter.Interface, services map[string]service.Service) httprouter.Handle {
	positionService := services["position"]
	authService := services["auth"]

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var positionUsers []entity.PositionUser
		var authUsers []entity.AuthUser

		params, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			log.Printf("Error: %v", err)
			_ = pr.Write(w, err.Error(), http.StatusBadRequest)
			return
		}

		date := params.Get("date")
		if date == "" {
			now := time.Now()
			date = now.AddDate(0, -1, 0).Format("2006-01-02")
		}

		positionFilter := "&filterge[updated_at]=" + url.QueryEscape(date)
		if err := positionService.Client.ProxyOld(r, positionService.Host + "/api/v1/users?limit=10000" + positionFilter, &positionUsers); err != nil {
			log.Printf("Error: %v", err)
			_ = pr.Write(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := authService.Client.ProxyOld(r,  authService.Host+"/api/v1/users?limit=10000", &authUsers); err != nil {
			_ = pr.Write(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := usecases.JoinUsers(positionUsers, authUsers)
		_ = pr.Write(w, resp, http.StatusOK)
	}
}

func GetProfile(pr presenter.Interface, services map[string]service.Service) httprouter.Handle {
	positionService := services["position"]
	authService := services["auth"]

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var positionUser entity.PositionUser
		var authUser entity.AuthUser

		req, err := http.NewRequest("GET", authService.Host+"/api/v1/profile", nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header = r.Header

		respProfile, err := authService.Client.ProxyData(req, &authUser)
		if err != nil {
			_ = pr.Write(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer respProfile.Body.Close()

		if respProfile.StatusCode != http.StatusOK {
			_ = pr.Write(w, nil, http.StatusUnauthorized)
			return
		}

		if err := positionService.Client.ProxyOld(r, positionService.Host+"/api/v1/users/"+strconv.Itoa(authUser.Id), &positionUser); err != nil {
			//do nothing
		}

		resp := usecases.Profile(authUser, positionUser)
		_ = pr.Write(w, resp, http.StatusOK)
	}
}

func PostPositions(pr presenter.Interface, services map[string]service.Service) httprouter.Handle {
	positionService := services["position"]
	authService := services["auth"]

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		req, err := http.NewRequest("GET", authService.Host+"/api/v1/profile", nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header = r.Header
		respProfile, err := authService.Client.ProxyData(req, nil)
		if err != nil {
			_ = pr.Write(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer respProfile.Body.Close()

		if respProfile.StatusCode != http.StatusOK {
			_ = pr.Write(w, nil, http.StatusUnauthorized)
			return
		}

		var position entity.PositionUser
		if err := positionService.Client.ProxyOld(r, positionService.Host+"/api/v1/positions", &position); err != nil {
			_ = pr.Write(w, err.Error(), http.StatusBadRequest)
			return
		}

		_ = pr.Write(w, position, http.StatusOK)
	}
}
