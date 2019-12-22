package main

import (
	"github.com/sr-2020/eva-gateway/app"
	"log"
)

func main() {
	cfg := app.InitConfig()

	log.Fatal(app.Start(cfg))
}
