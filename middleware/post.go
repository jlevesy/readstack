package middleware

import (
	"net/http"
)

func Post(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			if req.Method != http.MethodPost {
				http.NotFound(w, req)
				return
			}

			next.ServeHTTP(w, req)
		},
	)
}
