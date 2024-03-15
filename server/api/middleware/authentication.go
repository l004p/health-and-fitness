package middleware

import (
	"net/http"
	"fmt"
)

func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//verify token
			fmt.Println("Auth Middleware!")
			next.ServeHTTP(w, r)
		})
	}
}