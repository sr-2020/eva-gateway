package client

import "net/http"

type Client interface {
	ProxyOld(*http.Request, string, interface{}) error
	ProxyData(*http.Request, interface{}) (*http.Response, error)
}
