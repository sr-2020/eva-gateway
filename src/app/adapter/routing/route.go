package routing

import (
	"errors"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/sr-2020/eva-gateway/app/adapter/client"
	"github.com/sr-2020/eva-gateway/app/adapter/middleware"
	"github.com/sr-2020/eva-gateway/app/adapter/presenter"
	"github.com/sr-2020/eva-gateway/app/adapter/service"
	"log"
	"net/http"
	"net/url"
	"time"
)

func InitRoute(prefix string, router *httprouter.Router, services map[string]service.Service) {
	pr := presenter.NewJson()

	router.GET(prefix + "/version", GetVersion(pr))
	router.GET(prefix + "/users", GetUsers(pr, services))
	router.GET(prefix + "/profile", GetProfile(pr, services))
	router.POST(prefix + "/positions", PostPositions(pr, services))

	ServiceRouter(router, pr, services, middleware.ServiceMiddleware{
		Global: nil,
		Route: map[string][]middleware.Middleware{
			"/profile": {
				middleware.AuthMiddleware,
			},
			"/login": {
				middleware.LoginMiddleware,
			},
		},
	}, prefix + "/auth/*path", "auth")

	ServiceRouter(router, pr, services, middleware.ServiceMiddleware{
		Global: []middleware.Middleware{
			middleware.AuthMiddleware,
		},
		Route: map[string][]middleware.Middleware{
			"/balance": {
				middleware.AccountInfoMiddleware,
			},
			"/transfer": {
				middleware.TransferMiddleware,
			},
			"/transfers": {
				middleware.GetTransfersMiddleware,
			},
		},
	}, prefix + "/billing/*path", "billing")

	ServiceRouter(router, pr, services, middleware.ServiceMiddleware{
		Global: []middleware.Middleware{
			middleware.AuthWithAnonymousMiddleware,
		},
		Route: nil,
	},prefix + "/position/*path", "position")

	ServiceRouter(router, pr, services, middleware.ServiceMiddleware{
		Global: []middleware.Middleware{
			middleware.AuthMiddleware,
		},
		Route: nil,
	},prefix + "/push/*path", "push")

	ServiceRouter(router, pr, services, middleware.ServiceMiddleware{
		Global: []middleware.Middleware{
			middleware.AuthMiddleware,
		},
		Route: nil,
	},prefix + "/model-engine/*path", "model-engine")

	ServiceRouter(router, pr, services, middleware.ServiceMiddleware{
		Global: []middleware.Middleware{
			middleware.AuthMiddleware,
		},
		Route: map[string][]middleware.Middleware{
			"/character/model": {
				middleware.CharacterModelMiddleware,
			},
		},
	},prefix + "/models-manager/*path", "models-manager")
}

func ServiceRouter(router *httprouter.Router, pr presenter.Interface, ss map[string]service.Service, sm middleware.ServiceMiddleware, path string, serviceName string) {
	if serv, ok := ss[serviceName]; ok {
		handle := WrapHandler(ServiceRegister(pr, serv.Host + serv.Path, sm))

		router.GET(path, handle)
		router.POST(path, handle)
		router.PUT(path, handle)
		router.DELETE(path, handle)
	} else {
		log.Fatal(errors.New("Unknown service " + serviceName))
	}
}

func ServiceRegister(pr presenter.Interface, path string, middlewares middleware.ServiceMiddleware) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		ps := context.Get(r, "params").(httprouter.Params)
		p := ps.ByName("path")

		urlPath, err := url.Parse(path + p + "?" + r.URL.RawQuery)
		if err != nil {
			log.Fatal(err)
			return
		}
		r.URL = urlPath

		middlewaresList := middlewares.Global
		if routeMiddlewares, ok := middlewares.Route[p]; ok {
			middlewaresList = append(middlewaresList, routeMiddlewares...)
		}

		handler := alice.New(middlewaresList...).ThenFunc(ProxyHandler(pr))
		handler.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func WrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		context.Set(r, "params", ps)
		h.ServeHTTP(w, r)
	}
}

func ProxyHandler(pr presenter.Interface) http.HandlerFunc {
	httpClient := client.NewHttpClient(&http.Client{
		Timeout: time.Second * 10,
	})

	return func(w http.ResponseWriter, r *http.Request) {
		var data interface{}
		res, err := httpClient.ProxyData(r, &data)
		if err != nil {
			_ = pr.Write(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_ = pr.Write(w, data, res.StatusCode)
	}
}