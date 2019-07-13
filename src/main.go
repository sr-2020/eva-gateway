package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
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

func aHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("A[")
		next.ServeHTTP(w, r)
		fmt.Println("]")
	}

	return http.HandlerFunc(fn)
}

func bHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("B[")
		next.ServeHTTP(w, r)
		fmt.Println("]")
	}

	return http.HandlerFunc(fn)
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CHECK")
}

func indexHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Welcome2!")
	fmt.Println("C")

	chain := alice.New(aHandler, bHandler)
	handler := chain.ThenFunc(checkHandler)
	handler.ServeHTTP(w, r)
}

func main() {
	InitConfig()
	InitClient()
	InitService()

	//chain := New(aHandler, bHandler)
	//handler := chain.ThenFunc(indexHandler)
	//handler.ServeHTTP()

	router := httprouter.New()
	router.GET("/", indexHandler)

	router.GET("/api/v1/users", GetUsers)
	router.POST("/api/v1/positions", PostPositions)

	ServiceRouter(router, "/api/v1/auth/*path", "auth")
	ServiceRouter(router, "/api/v1/billing/*path", "billing")
	ServiceRouter(router, "/api/v1/position/*path", "position")
	ServiceRouter(router, "/api/v1/push/*path", "push")

	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(cfg.Port), router))
}
