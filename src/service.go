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

func ServiceRouter(router *httprouter.Router, path string, serviceName string) {
	if service, ok := Services[serviceName]; ok {
		handle := wrapHandler(ServiceRegister(service.Host + service.Path, service.Middleware))

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

		handler := alice.New(middlewaresList...).ThenFunc(proxyHandler)
		handler.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		context.Set(r, "params", ps)
		h.ServeHTTP(w, r)
	}
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	var data interface{}
	res, err := ProxyData(r, &data)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(res.StatusCode)
	responseBody, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(responseBody))
}
