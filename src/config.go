package main

import (
	"github.com/BurntSushi/toml"
	"os"
	"strconv"
)

type Config struct {
	Port     int
	Auth     string
	Position string
	Billing  string
	Push  string
}

var cfg Config

func InitConfig() {
	if _, err := toml.DecodeFile(".env", &cfg); err != nil {
		port, _ := strconv.Atoi(os.Getenv("GATEWAY_PORT"))
		cfg.Port = port

		cfg.Auth = os.Getenv("AUTH_HOST")
		cfg.Position = os.Getenv("POSITION_HOST")
		cfg.Billing = os.Getenv("BILLING_HOST")
		cfg.Push = os.Getenv("PUSH_HOST")
	}
}
