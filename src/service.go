package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
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

		for _, middleware := range middlewares.Global {
			p = middleware(w, r, ps)
		}

		if routeMiddlewares, ok := middlewares.Route[p]; ok {
			for _, middleware := range routeMiddlewares {
				p = middleware(w, r, ps)
			}
		}

		url := path + p

		var resp interface{}
		res, err := Proxy(r, url, &resp)
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
