package middleware

import (
	"amp-templates/server/services/log"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.RemoteAddr + " " + r.Method + " " + r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
