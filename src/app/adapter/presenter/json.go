package presenter

import (
	"encoding/json"
	"net/http"
)

type Json struct {
}

func NewJson() Json {
	return Json{}
}

func (j Json) Write(w http.ResponseWriter, data interface{}, statusCode int) error {
	response, ok := data.([]byte)
	if ok {
		w.WriteHeader(statusCode)
		_, _ = w.Write(response)
		return nil
	}

	response, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_, _ = w.Write(response)

	return nil
}

func (j Json) WriteRaw(w http.ResponseWriter, data []byte, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_, _ = w.Write(data)

	return nil
}
