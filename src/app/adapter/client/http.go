package client

import (
	"fmt"
	"github.com/sr-2020/eva-gateway/app/adapter/support"
	"log"
	"net/http"
)

type HttpClient struct {
	client *http.Client
}

func NewHttpClient(client *http.Client) *HttpClient {
	return &HttpClient{client}
}

func (c HttpClient) ProxyOld(request *http.Request, url string, data interface{}) error {
	req, err := http.NewRequest(request.Method, url, request.Body)
	if err != nil {
		return err
	}

	req.Header = request.Header

	res, getErr := c.client.Do(req)
	if getErr != nil {
		return getErr
	}

	if err := support.GetBodyToInterface(&res.Body, &data); err != nil {
		return err
	}

	return nil
}

func (c HttpClient) Proxy(request *http.Request) (*http.Response, error) {
	req, err := http.NewRequest(request.Method, request.URL.String(), request.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header = request.Header
	res, reqErr := c.client.Do(req)
	if reqErr != nil {
		log.Println(reqErr)
		return res, reqErr
	}

	return res, nil
}

func (c HttpClient) ProxyData(request *http.Request, data interface{}) (*http.Response, error) {
	res, err := c.Proxy(request)
	if err != nil {
		log.Println(err)
		return res, err
	}

	if err := support.GetBodyToInterface(&res.Body, data); err != nil {
		log.Println(err)
		return res, nil
	}

	return res, nil
}

func (c HttpClient) ErrorResponse(w http.ResponseWriter, code int, err error) {
	log.Printf("Status code: %d Error: %v", code, err)
	w.WriteHeader(code)
	fmt.Fprint(w, err.Error())
}
