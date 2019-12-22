package app

import (
	"github.com/BurntSushi/toml"
	"github.com/julienschmidt/httprouter"
	"github.com/sr-2020/eva-gateway/app/adapter/client"
	"github.com/sr-2020/eva-gateway/app/adapter/middleware"
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

var cfg Config
var Services map[string]service.Service

func InitConfig() Config {
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

func InitServices(cfg Config, client client.Client) {
	Services = map[string]service.Service{
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

}

func Start(cfg Config) error {
	httpClient := client.NewHttpClient(&http.Client{
		Timeout: time.Second * 10,
	})

	InitServices(cfg, httpClient)
	middleware.InitMiddleware(Services)

	router := httprouter.New()
	routing.InitRoute(router, Services)

	return http.ListenAndServe(":" + strconv.Itoa(cfg.Port), router)
}
