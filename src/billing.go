package main

import (
	"encoding/json"
	"fmt"
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

func AccountInfoMiddleware(w http.ResponseWriter, r *http.Request, ps httprouter.Params) string {
	Auth(r)
	return ps.ByName("path") + "/" + strconv.Itoa(authUser.Id)
}

func TransferMiddleware(w http.ResponseWriter, r *http.Request, ps httprouter.Params) string {
	Auth(r)

	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var transfer Transfer
	jsonErr := json.Unmarshal(body, &transfer)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	transfer.From = authUser.Id

	bodyRequest, err := json.Marshal(transfer)
	if err != nil {
		log.Fatal(err)
	}

	t := strings.NewReader(string(bodyRequest))
	rc := ioutil.NopCloser(t)
	r.Body = rc

	return ps.ByName("path")
}

func BillingService(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := ps.ByName("path")

	switch path {
	case "/account_info":
		path = AccountInfoMiddleware(w, r, ps)
	case "/transfer":
		path = TransferMiddleware(w, r, ps)
	default:
		//
	}

	var resp interface{}
	res, err := ProxyLite(r, cfg.Billing + path, &resp)
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
