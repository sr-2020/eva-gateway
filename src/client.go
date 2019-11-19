package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var httpClient *http.Client

func InitClient() {
	httpClient = &http.Client{
		Timeout: time.Second * 10,
	}
}

func ProxyOld(request *http.Request, url string, data interface{}) error {
	req, err := http.NewRequest(request.Method, url, request.Body)
	if err != nil {
		return err
	}

	req.Header = request.Header

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		return getErr
	}

	if err := getBodyToInterface(&res.Body, &data); err != nil {
		return err
	}

	return nil
}

func Proxy(request *http.Request) (*http.Response, error) {
	req, err := http.NewRequest(request.Method, request.URL.String(), request.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header = request.Header
	res, reqErr := httpClient.Do(req)
	if reqErr != nil {
		log.Println(reqErr)
		return res, reqErr
	}

	return res, nil
}

func ProxyData(request *http.Request, data interface{}) (*http.Response, error) {
	res, err := Proxy(request)
	if err != nil {
		log.Println(err)
		return res, err
	}

	if err := getBodyToInterface(&res.Body, data); err != nil {
		log.Println(err)
		return res, nil
	}

	return res, nil
}

func getBody(r *io.ReadCloser) ([]byte, error) {
	body, err := ioutil.ReadAll(*r)

	if err != nil {
		return []byte{}, err
	}

	*r = ioutil.NopCloser(strings.NewReader(string(body)))

	return body, nil
}

func getBodyToInterface(r *io.ReadCloser, data interface{}) error {
	body, err := getBody(r)
	if err != nil {
		log.Println(err)
		return err
	}

	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		log.Println(jsonErr)
		return jsonErr
	}

	return nil
}

func setInterfaceToBody(data interface{}, body *io.ReadCloser) error {
	bodyResp, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return err
	}
	*body = ioutil.NopCloser(strings.NewReader(string(bodyResp)))

	return nil
}

func ErrorResponse(w http.ResponseWriter, code int, err error) {
	log.Printf("Status code: %d Error: %v", code, err)
	w.WriteHeader(code)
	fmt.Fprint(w, err.Error())
}
