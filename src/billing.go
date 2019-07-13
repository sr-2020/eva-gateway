package main

import (
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

type Transfer struct {
	Id                 int    `json:"id"`
	From               int    `json:"sin_from"`
	To                 int    `json:"sin_to"`
	Amount             int    `json:"amount"`
	RecurrentPaymentId int    `json:"recurrent_payment_id"`
	Comment            string `json:"comment"`
	CreatedAt          string `json:"created_at"`
}

func NewAccountInfoMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ps := context.Get(r, "params").(httprouter.Params)
		sin := r.Header.Get("X-User-Id")
		r.URL.Path = ps.ByName("path") + "/" + sin
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func NewTransferMiddleware(next http.Handler) http.Handler {
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

func AccountInfoMiddleware(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (*http.Response, error) {
	sin := r.Header.Get("X-User-Id")
	r.URL.Path = ps.ByName("path") + "/" + sin
	return nil, nil
}

func TransferMiddleware(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (*http.Response, error) {
	var transfer Transfer
	if err := getBodyToInterface(&r.Body, &transfer); err != nil {
		log.Println(err)
		return nil, err
	}

	sin, err := strconv.Atoi(r.Header.Get("X-User-Id"))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	transfer.From = sin

	if err := setInterfaceToBody(transfer, &r.Body); err != nil {
		log.Println(err)
		return nil, err
	}

	return nil, nil
}
