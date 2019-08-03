package main

import (
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

type Transfer struct {
	From               int    `json:"sin_from"`
	To                 int    `json:"sin_to"`
	Amount             int    `json:"amount"`
	RecurrentPaymentId int    `json:"recurrent_payment_id"`
	Comment            string `json:"comment"`
	CreatedAt          string `json:"created_at"`
}

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
		var transfer Transfer
		if err := getBodyToInterface(&r.Body, &transfer); err != nil {
			log.Println(err)
		}

		sin, err := strconv.Atoi(r.Header.Get("X-User-Id"))
		if err != nil {
			log.Println(err)
		}
		transfer.From = sin

		if err := setInterfaceToBody(transfer, &r.Body); err != nil {
			log.Println(err)
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
