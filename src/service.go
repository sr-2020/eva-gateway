package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
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
			if _, err := middleware(w, r, ps); err != nil {
				log.Fatal(err)
			}
		}

		var response *http.Response
		if routeMiddlewares, ok := middlewares.Route[p]; ok {
			for _, middleware := range routeMiddlewares {
				if response, err = middleware(w, r, ps); err != nil {
					log.Fatal(err)
				}
			}
		}

		res := response
		var resp interface{}
		if nil == res {
			res, err = Proxy(r, &resp)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			body, readErr := ioutil.ReadAll(res.Body)
			if readErr != nil {
				log.Fatal(readErr)
			}
			var data interface{}

			json.Unmarshal(body, &data)
			Decode(&resp, data)
		}
		w.WriteHeader(res.StatusCode)
		respons, err := json.Marshal(resp)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprint(w, string(respons))
	}
}
