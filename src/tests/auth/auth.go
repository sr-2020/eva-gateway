package auth

import (
	"encoding/json"
	"fmt"
	"github.com/sr-2020/eva-gateway/app/entity"
	"io/ioutil"
	"net/http"
	"strings"
)

type Auth struct {
	host string
}

func NewAuth(host string) *Auth {
	return &Auth{host}
}

func (a *Auth) Check() bool {
	resp, _ := http.Get(fmt.Sprintf("%s/auth/version", a.host))
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body) == "2016"
}

func (a *Auth) Auth(data map[string]string) (Token, int, error) {
	requestBody, _ := json.Marshal(data)

	dt := strings.NewReader(string(requestBody))
	resp, _ := http.Post(fmt.Sprintf("%s/auth/login", a.host), "application/json", dt)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	token := Token{}
	if err := json.Unmarshal(body, &token); err != nil {
		return token, resp.StatusCode, err
	}

	return token, resp.StatusCode, nil
}

func (a *Auth) Register(data map[string]string) (Token, int, error) {
	requestBody, _ := json.Marshal(data)

	dt := strings.NewReader(string(requestBody))
	resp, _ := http.Post(fmt.Sprintf("%s/auth/register", a.host), "application/json", dt)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	token := Token{}
	if err := json.Unmarshal(body, &token); err != nil {
		return token, resp.StatusCode, err
	}

	return token, resp.StatusCode, nil
}

func (a *Auth) ReadProfile(apiKey string) (entity.ProfileUser, int, error) {
	var user entity.ProfileUser

	client := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/auth/profile", a.host), nil)
	if err != nil {
		return user, 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &user); err != nil {
		return user, resp.StatusCode, err
	}

	return user, resp.StatusCode, nil
}

func (a *Auth) EditProfile(apiKey string, data map[string]interface{}) (entity.ProfileUser, int, error) {
	var user entity.ProfileUser

	requestBody, _ := json.Marshal(data)
	dt := strings.NewReader(string(requestBody))

	client := http.Client{}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/auth/profile", a.host), dt)
	if err != nil {
		return user, 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &user); err != nil {
		return user, resp.StatusCode, err
	}

	return user, resp.StatusCode, nil
}
