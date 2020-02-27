package main

import (
	"fmt"
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

		r.Method = "GET"
		r.URL.Path = "/api/billing/transfer/maketransfersinsin"
		r.URL.RawQuery = fmt.Sprintf("character1=%d&character2=%d&amount=%d&comment=%s",
			transfer.From, transfer.To, transfer.Amount, transfer.Comment)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func AccountInfoMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		sin := r.Header.Get("X-User-Id")
		r.URL.Path = "/api/billing/info/getbalance"
		r.URL.RawQuery = "characterId=" + sin

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func GetTransfersMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		sin := r.Header.Get("X-User-Id")
		r.URL.Path = "/api/Billing/info/gettransfers"
		r.URL.RawQuery = "characterId=" + sin

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
