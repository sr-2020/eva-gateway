package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

func main() {
	InitConfig()
	InitClient()
	InitService()

	router := httprouter.New()
	router.GET("/api/v1/users", GetUsers)
	router.POST("/api/v1/positions", PostPositions)

	ServiceRouter(router, "/api/v1/auth/*path", "auth")
	ServiceRouter(router, "/api/v1/billing/*path", "billing")
	ServiceRouter(router, "/api/v1/position/*path", "position")

	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(cfg.Port), router))
}
