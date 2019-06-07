package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"net/url"
)

func ServiceRouter(router *httprouter.Router, path string, handle httprouter.Handle) {
	router.GET(path, handle)
	router.POST(path, handle)
	router.PUT(path, handle)
	router.DELETE(path, handle)
}

func Service(path string, middlewares ServiceMiddleware) httprouter.Handle {
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
