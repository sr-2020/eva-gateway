package app

import (
	"github.com/BurntSushi/toml"
	"github.com/julienschmidt/httprouter"
	"github.com/sr-2020/eva-gateway/app/adapter/client"
	"github.com/sr-2020/eva-gateway/app/adapter/routing"
	"github.com/sr-2020/eva-gateway/app/adapter/service"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port          int
	Auth          string
	Position      string
	Billing       string
	Push          string
	ModelEngine   string
	ModelsManager string
}

func InitConfig() Config {
	var cfg Config

	if _, err := toml.DecodeFile(".env", &cfg); err != nil {
		port, _ := strconv.Atoi(os.Getenv("GATEWAY_PORT"))
		cfg.Port = port

		cfg.Auth = os.Getenv("AUTH_HOST")
		cfg.Position = os.Getenv("POSITION_HOST")
		cfg.Billing = os.Getenv("BILLING_HOST")
		cfg.Push = os.Getenv("PUSH_HOST")
		cfg.ModelEngine = os.Getenv("MODEL_ENGINE_HOST")
		cfg.ModelsManager = os.Getenv("MODELS_MANAGER_HOST")
	}

	return cfg
}

func InitServices(cfg Config, client client.Client) map[string]service.Service {
	service.Services = map[string]service.Service{
		"auth": {
			Host:       cfg.Auth,
			Path:       "/api/v1",
			Client:     client,
		},
		"billing": {
			Host:       cfg.Billing,
			Path:       "",
			Client:     client,
		},
		"position": {
			Host:       cfg.Position,
			Path:       "/api/v1",
			Client:     client,
		},
		"push": {
			Host:       cfg.Push,
			Path:       "",
			Client:     client,
		},
		"model-engine": {
			Host:       cfg.ModelEngine,
			Path:       "",
			Client:     client,
		},
		"models-manager": {
			Host:       cfg.ModelsManager,
			Path:       "",
			Client:     client,
		},
	}

	return service.Services
}

func Start(cfg Config) error {
	httpClient := client.NewHttpClient(&http.Client{
		Timeout: time.Second * 10,
	})

	services := InitServices(cfg, httpClient)

	router := httprouter.New()
	routing.InitRoute("/api/v1", router, services)

	return http.ListenAndServe(":" + strconv.Itoa(cfg.Port), router)
}
