package main

import (
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
