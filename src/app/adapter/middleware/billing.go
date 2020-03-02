package middleware

import (
	"fmt"
	"github.com/sr-2020/eva-gateway/app/adapter/support"
	"github.com/sr-2020/eva-gateway/app/entity"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

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

		r.Method = "GET"
		r.URL.Path = "/api/billing/transfer/maketransfersinsin"
		r.URL.RawQuery = fmt.Sprintf("character1=%d&character2=%d&amount=%d&comment=%s",
			transfer.From, transfer.To, transfer.Amount, url.QueryEscape(transfer.Comment))

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
