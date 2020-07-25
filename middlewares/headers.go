package middlewares

import (
	"net/http"
)

// SetHeaders MW
func SetHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow_Origin", "*")
		res.Header().Set("content-type", "application/json")
		next.ServeHTTP(res, req)
	})
}
