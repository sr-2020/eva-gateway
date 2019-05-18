package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

var authUser AuthUser

func main() {
	InitConfig()

	router := httprouter.New()
	router.GET("/api/v1/users", GetUsers)
	router.POST("/api/v1/positions", PostPositions)

	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(cfg.Port) , router))
}
