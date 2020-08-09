package middleware

import (
	"github.com/sr-2020/eva-gateway/app/adapter/support"
	"net/http"
)

func CorsMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		support.SetCorsHeaders(w, r.Header.Get("origin"))

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
