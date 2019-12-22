package middleware

import (
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/sr-2020/eva-gateway/app/adapter/support"
	"github.com/sr-2020/eva-gateway/app/entity"
	"log"
	"net/http"
	"strconv"
)

func AccountInfoMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ps := context.Get(r, "params").(httprouter.Params)
		sin := r.Header.Get("X-User-Id")
		r.URL.Path = ps.ByName("path") + "/" + sin
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func TransferMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var transfer entity.Transfer
		if err := support.GetBodyToInterface(&r.Body, &transfer); err != nil {
			log.Println(err)
		}

		sin, err := strconv.Atoi(r.Header.Get("X-User-Id"))
		if err != nil {
			log.Println(err)
		}
		transfer.From = sin

		if err := support.SetInterfaceToBody(transfer, &r.Body); err != nil {
			log.Println(err)
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

