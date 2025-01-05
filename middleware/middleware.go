package middleware

import (
	"net/http"
)

func EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

		next.ServeHTTP(w, r)
	})
}