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

	router.GET("/api/v1/billing/*path", BillingService)
	router.POST("/api/v1/billing/*path", BillingService)

	router.GET("/api/v1/position/*path", PositionService)
	router.POST("/api/v1/position/*path", PositionService)

	router.GET("/api/v1/auth/*path", AuthService)
	router.POST("/api/v1/auth/*path", AuthService)
	router.PUT("/api/v1/auth/*path", AuthService)

	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(cfg.Port), router))
}
