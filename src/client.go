package main

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var httpClient *http.Client

func InitClient() {
	httpClient = &http.Client{
		Timeout: time.Second * 5,
	}
}

func ProxyOld(request *http.Request, url string, data interface{}, error interface{}) error {
	req, err := http.NewRequest(request.Method, url, request.Body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header = request.Header

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var response interface{}

	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	if 400 == res.StatusCode {
		Decode(&error, response)
		return nil
	}

	Decode(&data, response)

	return nil
}

func Proxy(request *http.Request, url string, data interface{}) (*http.Response, error) {
	req, err := http.NewRequest(request.Method, url, request.Body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header = request.Header
	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var response interface{}
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		Decode(&data, jsonErr.Error())
		return res, nil
	}

	Decode(&data, response)

	return res, nil
}

func Decode(out interface{}, in interface{}) {
	decConfig := mapstructure.DecoderConfig{
		TagName: "json",
		Result: &out,
		WeaklyTypedInput:true,
	}
	dec1, _ := mapstructure.NewDecoder(&decConfig)
	dec1.Decode(in)
}
