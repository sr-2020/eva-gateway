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

	router := httprouter.New()
	router.GET("/api/v1/users", GetUsers)
	router.POST("/api/v1/positions", PostPositions)
	router.POST("/api/v1/transfer", PostTransfer)
	router.GET("/api/v1/account_info", GetAccountInfo)

	ServiceRouter(router, "/api/v1/auth/*path", Service(cfg.Auth + "/api/v1", MiddlewareMap.Auth))
	ServiceRouter(router, "/api/v1/billing/*path", Service(cfg.Billing, MiddlewareMap.Billing))
	ServiceRouter(router, "/api/v1/position/*path", Service(cfg.Position + "/api/v1", MiddlewareMap.Position))

	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(cfg.Port), router))
}
