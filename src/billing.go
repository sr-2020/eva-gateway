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

type BillingErrorWrap struct {
	Error BillingError `json:"error"`
}

type BillingError struct {
	StatusCode int    `json:"statusCode"`
	Name       string `json:"name"`
	Message    string `json:"message"`
}

func PostTransfer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Auth(r, cfg.Auth+"/api/v1/profile", &authUser)

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

	var rTransfer interface{}
	var billingError BillingErrorWrap

	Proxy(r, cfg.Billing+"/transfer", &rTransfer, &billingError)

	if 400 == billingError.Error.StatusCode {
		response, err := json.Marshal(billingError)
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(400)
		fmt.Fprint(w, string(response))
		return
	}

	response, err := json.Marshal(rTransfer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(response))
}

func GetAccountInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Auth(r, cfg.Auth+"/api/v1/profile", &authUser)

	var resp interface{}

	Proxy(r, cfg.Billing+"/account_info/"+strconv.Itoa(authUser.Id), &resp, nil)

	response, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(response))
}
