package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var transfer Transfer
	jsonErr := json.Unmarshal(body, &transfer)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	sin, err := strconv.Atoi(r.Header.Get("X-User-Id"))
	if err != nil {
		log.Fatal(err)
	}
	transfer.From = sin

	bodyRequest, err := json.Marshal(transfer)
	if err != nil {
		log.Fatal(err)
	}

	t := strings.NewReader(string(bodyRequest))
	rc := ioutil.NopCloser(t)
	r.Body = rc

	return nil, nil
}
