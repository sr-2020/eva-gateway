package app

import (
	"github.com/BurntSushi/toml"
	"github.com/go-redis/redis/v7"
	"github.com/julienschmidt/httprouter"
	"github.com/sr-2020/eva-gateway/app/adapter/client"
	"github.com/sr-2020/eva-gateway/app/adapter/config"
	"github.com/sr-2020/eva-gateway/app/adapter/routing"
	"github.com/sr-2020/eva-gateway/app/adapter/service"
	"net/http"
	"os"
	"strconv"
	"time"
)

func InitConfig() config.Config {
	var cfg config.Config

	if _, err := toml.DecodeFile(".env", &cfg); err != nil {
		port, _ := strconv.Atoi(os.Getenv("GATEWAY_PORT"))
		cfg.Port = port

		cfg.ApiKey = os.Getenv("GATEWAY_API_KEY")
		cfg.JwtSecret = os.Getenv("JWT_SECRET")

		cfg.Gateway = os.Getenv("GATEWAY_HOST")
		cfg.Auth = os.Getenv("AUTH_HOST")
		cfg.Position = os.Getenv("POSITION_HOST")
		cfg.Billing = os.Getenv("BILLING_HOST")
		cfg.Push = os.Getenv("PUSH_HOST")
		cfg.ModelEngine = os.Getenv("MODEL_ENGINE_HOST")
		cfg.ModelsManager = os.Getenv("MODELS_MANAGER_HOST")
		cfg.Redis = os.Getenv("REDIS_HOST")
	}

	config.Cfg = cfg

	return cfg
}

func InitServices(cfg config.Config, client client.Client) map[string]service.Service {
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

func Start(cfg config.Config) error {
	httpClient := client.NewHttpClient(&http.Client{
		Timeout: time.Second * 10,
	})

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	services := InitServices(cfg, httpClient)

	router := httprouter.New()
	routing.InitRoute("/api/v1", router, redisClient, services)
	routing.EnableCors(router)

	return http.ListenAndServe(":" + strconv.Itoa(cfg.Port), router)
}
