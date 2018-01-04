package middleware

import (
	"log"
	"net/http"
)

type statusCodeCollector struct {
	http.ResponseWriter
	StatusCode int
}

func (s *statusCodeCollector) WriteHeader(statusCode int) {
	s.ResponseWriter.WriteHeader(statusCode)
	s.StatusCode = statusCode
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s", r.Method, r.URL.Path)

			s := statusCodeCollector{ResponseWriter: w}

			next.ServeHTTP(&s, r)

			log.Printf(
				"%s %s, Status %d, Duration %v",
				r.Method,
				r.URL.Path,
				s.StatusCode,
				GetTimeRecorder(r.Context()).Read(HandlerProbe),
			)
		},
	)
}
