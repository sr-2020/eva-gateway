package middleware

import (
	"net/http"
)

func CorsMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
