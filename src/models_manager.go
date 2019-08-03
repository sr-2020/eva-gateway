package main

import (
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func CharacterModelMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ps := context.Get(r, "params").(httprouter.Params)
		userId := r.Header.Get("X-User-Id")
		r.URL.Path = ps.ByName("path") + "/" + userId
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
