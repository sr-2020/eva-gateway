package middleware

import (
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/sr-2020/eva-gateway/app/adapter/service"
	"github.com/sr-2020/eva-gateway/app/adapter/support"
	"github.com/sr-2020/eva-gateway/app/entity"
	"log"
	"net/http"
)

func CharacterModelMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ps := context.Get(r, "params").(httprouter.Params)
		userId := r.Header.Get("X-User-Id")
		r.URL.Path = ps.ByName("path") + "/" + userId

		if r.Method == "POST" {
			modelsManagerLocation := entity.ModelsManagerLocation{
				Id: 1,
				ManaLevel: 0,
			}

			positionService := service.Services["position"]

			var positionUser entity.PositionUser
			req, err := http.NewRequest("GET", positionService.Host, nil)
			if err != nil {
				log.Fatal(err)
			}

			req.Header = r.Header
			if err := positionService.Client.ProxyOld(req, positionService.Host + "/api/v1/users/" + userId, &positionUser); err != nil {
				log.Printf("Error: %v", err)
			}

			var event entity.ModelsManagerEvent
			if err := support.GetBodyToInterface(&r.Body, &event); err != nil {
				log.Println(err)
			}
			if positionUser.Location != nil && positionUser.Location.Id != 0 {
				modelsManagerLocation.Id = positionUser.Location.Id
				if v, ok := positionUser.Location.Options["manaLevel"]; ok {
					if v, ok := v.(float64); ok {
						modelsManagerLocation.ManaLevel = int(v)
					}
				}
			}

			event.Data["location"] = modelsManagerLocation
			if err := support.SetInterfaceToBody(event, &r.Body); err != nil {
				log.Println(err)
			}
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
