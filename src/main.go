package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"time"
)

func loggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

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
	ServiceRouter(router, "/api/v1/push/*path", "push")

	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(cfg.Port), router))
}
