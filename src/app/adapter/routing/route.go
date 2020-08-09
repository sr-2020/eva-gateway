package routing

import (
	"errors"
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/sr-2020/eva-gateway/app/adapter/client"
	"github.com/sr-2020/eva-gateway/app/adapter/middleware"
	"github.com/sr-2020/eva-gateway/app/adapter/presenter"
	"github.com/sr-2020/eva-gateway/app/adapter/service"
	"github.com/sr-2020/eva-gateway/app/adapter/support"
	"log"
	"net/http"
	"net/url"
	"time"
)

func EnableCors(router *httprouter.Router) {
	router.HandleOPTIONS = true
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		w.WriteHeader(http.StatusNoContent)
	})
}

func InitRoute(prefix string, router *httprouter.Router, redisClient *redis.Client, services map[string]service.Service) {
	pr := presenter.NewJson()

	router.GET(prefix + "/config/:key", GetConfig(pr, redisClient))
	router.POST(prefix + "/config/:key", SetConfig(pr, redisClient))

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
			middleware.AuthWithAnonymousMiddleware,
			middleware.CorsMiddleware,
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
			middleware.CorsMiddleware,
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
		defer res.Body.Close()

		if data == nil {
			body, err := support.GetBody(&res.Body)
			if err != nil {
				_ = pr.Write(w, err.Error(), http.StatusInternalServerError)
				return
			}

			_ = pr.Write(w, body, res.StatusCode)
			return
		}

		_ = pr.Write(w, data, res.StatusCode)
	}
}