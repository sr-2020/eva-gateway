package presenter

import "net/http"

type Interface interface {
	Write(http.ResponseWriter, interface{}, int) error
	WriteRaw(http.ResponseWriter, []byte, int) error
}
