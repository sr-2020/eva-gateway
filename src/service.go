package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"net/url"
)

type Service struct {
	Host string
	Path string
	Middleware ServiceMiddleware
}

func InitService() {
	Services = map[string]Service{
		"auth": {
			Host: cfg.Auth,
			Path: "/api/v1",
			Middleware: MiddlewareMap.Auth,
		},
		"billing": {
			Host: cfg.Billing,
			Path: "",
			Middleware: MiddlewareMap.Billing,
		},
		"position": {
			Host: cfg.Position,
			Path: "/api/v1",
			Middleware: MiddlewareMap.Position,
		},
	}
}

func ServiceRouter(router *httprouter.Router, path string, serviceName string) {
	if service, ok := Services[serviceName]; ok {
		handle := ServiceRegister(service.Host + service.Path, service.Middleware)

		router.GET(path, handle)
		router.POST(path, handle)
		router.PUT(path, handle)
		router.DELETE(path, handle)
	} else {
		log.Fatal(errors.New("Unknown service " + serviceName))
	}
}

func ServiceRegister(path string, middlewares ServiceMiddleware) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		p := ps.ByName("path")
		url, err := url.Parse(path + p)
		if err != nil {
			log.Fatal(err)
		}
		r.URL = url

		for _, middleware := range middlewares.Global {
			if err := middleware(w, r, ps); err != nil {
				log.Fatal(err)
			}
		}

		if routeMiddlewares, ok := middlewares.Route[p]; ok {
			for _, middleware := range routeMiddlewares {
				if err := middleware(w, r, ps); err != nil {
					log.Fatal(err)
				}
			}
		}

		var resp interface{}
		res, err := Proxy(r, &resp)
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(res.StatusCode)
		response, err := json.Marshal(resp)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprint(w, string(response))
	}
}
