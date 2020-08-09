package support

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetBody(r *io.ReadCloser) ([]byte, error) {
	body, err := ioutil.ReadAll(*r)

	if err != nil {
		return []byte{}, err
	}

	*r = ioutil.NopCloser(strings.NewReader(string(body)))

	return body, nil
}

func GetBodyToInterface(r *io.ReadCloser, data interface{}) error {
	body, err := GetBody(r)
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

func SetInterfaceToBody(data interface{}, body *io.ReadCloser) error {
	bodyResp, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return err
	}
	*body = ioutil.NopCloser(strings.NewReader(string(bodyResp)))

	return nil
}

func SetCorsHeaders(w http.ResponseWriter, origin string) {
	if origin == "" {
		origin = "*"
	}

	w.Header().Add("Access-Control-Allow-Origin", origin)
	w.Header().Add("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE,HEAD,PATCH")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Origin,X-Requested-With,Content-Type,Accept")
}
