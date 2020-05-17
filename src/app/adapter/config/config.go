package config

type Config struct {
	Port          int
	ApiKey        string
	JwtSecret     string
	Gateway       string
	Auth          string
	Position      string
	Billing       string
	Push          string
	ModelEngine   string
	ModelsManager string
	Redis         string
}

var Cfg Config
