package presenter

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Json struct {
}

func NewJson() Json {
	return Json{}
}

func (j Json) Write(w http.ResponseWriter, data interface{}, statusCode int) error {
	response, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(response))

	return nil
}

