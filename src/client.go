package main

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Proxy(request *http.Request, url string, data interface{}) error {
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(request.Method, url, request.Body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header = request.Header
	request.Header.Set("X-User-Id", strconv.Itoa(authUser.Id))

	res, getErr := spaceClient.Do(req)
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

	Decode(&data, response)

	return nil
}

func Auth(request *http.Request, url string, data interface{}) error {
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", request.Header.Get("Authorization"))

	res, getErr := spaceClient.Do(req)
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

	Decode(&data, response)

	return nil
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
