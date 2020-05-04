package auth

type Token struct {
	Id int `json:"id"`
	ApiKey string `json:"api_key"`
}
