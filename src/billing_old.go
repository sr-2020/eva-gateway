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

type BillingErrorWrap struct {
	Error BillingError `json:"error"`
}

type BillingError struct {
	StatusCode int    `json:"statusCode"`
	Name       string `json:"name"`
	Message    string `json:"message"`
}

func PostTransfer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	var rTransfer interface{}
	var billingError BillingErrorWrap

	ProxyOld(r, cfg.Billing+"/transfer", &rTransfer, &billingError)

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
	Auth(r)

	var resp interface{}

	ProxyOld(r, cfg.Billing+"/account_info/" + r.Header.Get("X-User-Id"), &resp, nil)

	response, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(response))
}
