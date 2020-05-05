package actors

import (
	"fmt"
	"net/http"
)

type Admin struct {
	ApiKey string
	Host string
}

func NewAdmin(apiKey, host string) *Admin {
	return &Admin{apiKey, host}
}

func (a *Admin) DeleteUser(id int) error {
	client := http.Client{}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/auth/users/%d", a.Host, id), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.ApiKey))
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	_ = resp
	return nil
}