package middleware

import (
	"net/http"
	"time"

	"github.com/jlevesy/readstack/timing"
)

const (
	HandlerDuration = "handler"
)

func WithInMemoryTimingRecorder(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			rec := timing.NewInMemoryRecorder()

			next.ServeHTTP(
				w,
				r.WithContext(
					timing.WithRecorder(
						r.Context(),
						rec,
					),
				),
			)
		},
	)
}

func RecordDuration(metric string, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			next.ServeHTTP(w, r)

			end := time.Now()

			timing.GetRecorder(r.Context()).Write(metric, end.Sub(start))
		},
	)
}
