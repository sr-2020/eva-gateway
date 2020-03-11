package routing

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sr-2020/eva-gateway/app/adapter/presenter"
	"github.com/sr-2020/eva-gateway/app/adapter/service"
	"github.com/sr-2020/eva-gateway/app/entity"
	"github.com/sr-2020/eva-gateway/app/usecases"
	"log"
	"net/http"
	"strconv"
)

func GetUsers(pr presenter.Interface, services map[string]service.Service) httprouter.Handle {
	positionService := services["position"]
	authService := services["auth"]

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var positionUsers []entity.PositionUser
		var authUsers []entity.AuthUser

		if err := positionService.Client.ProxyOld(r, positionService.Host + "/api/v1/users?limit=10000", &positionUsers); err != nil {
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
