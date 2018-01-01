package middleware

import (
	"log"
	"net/http"
	"time"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Start %s", r.URL.Path)

			next.ServeHTTP(w, r)

			log.Printf(
				"End %s, Status %s, Duration %v",
				r.URL.Path,
				"[TODO GET STATUS !?]",
				20*time.Millisecond,
			)
		},
	)
}
