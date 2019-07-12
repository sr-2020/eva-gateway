package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"log"
	"net/http"
	"net/url"
)

type Service struct {
	Host string
	Path string
	Middleware ServiceMiddleware
}

var Services map[string]Service

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
		"push": {
			Host: cfg.Push,
			Path: "",
			Middleware: MiddlewareMap.Push,
		},
	}
}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		context.Set(r, "params", ps)
		h.ServeHTTP(w, r)
	}
}

func ServiceRouter(router *httprouter.Router, path string, serviceName string) {
	commonHandlers := alice.New(context.ClearHandler, loggingHandler)
	if service, ok := Services[serviceName]; ok {
		serviceHandler := ServiceRegister(service.Host + service.Path, service.Middleware)
		handle := wrapHandler(commonHandlers.Then(serviceHandler))

		router.GET(path, handle)
		router.POST(path, handle)
		router.PUT(path, handle)
		router.DELETE(path, handle)
	} else {
		log.Fatal(errors.New("Unknown service " + serviceName))
	}
}

func ServiceRegister(path string, middlewares ServiceMiddleware) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		ps := context.Get(r, "params").(httprouter.Params)
		p := ps.ByName("path")

		urlPath, err := url.Parse(path + p)
		if err != nil {
			log.Fatal(err)
		}
		r.URL = urlPath

		middlewaresList := middlewares.Global
		if routeMiddlewares, ok := middlewares.Route[p]; ok {
			middlewaresList = append(middlewaresList, routeMiddlewares...)
		}

		var response *http.Response
		for _, middleware := range middlewaresList {
			if response, err = middleware(w, r, ps); err != nil {
				log.Fatal(err)
			}
		}

		res := response
		var data interface{}
		if nil == res {
			res, err = ProxyData(r, &data)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			if err := getBodyToInterface(&res.Body, &data); err != nil {
				log.Println(err)
			}
		}
		w.WriteHeader(res.StatusCode)
		responseBody, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprint(w, string(responseBody))
	}

	return http.HandlerFunc(fn)
}
