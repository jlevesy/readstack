package middleware

import (
	"net/http"

	"github.com/jlevesy/readstack/logger"
	"github.com/jlevesy/readstack/timing"
)

// statusCodeCollector is a structure satisfying http.ResponseWriter
// It is injected by this middleware in order to capture the returned
// status code.
type statusCodeCollector struct {
	http.ResponseWriter
	StatusCode int
}

func (s *statusCodeCollector) WriteHeader(statusCode int) {
	s.ResponseWriter.WriteHeader(statusCode)
	s.StatusCode = statusCode
}

// RequestLogger is a middleware dedicated to log incoming and processed requests
// It currently logs status code, and timer thanks to the timing package.
// However it requires to be used under a middleware injecting a timing.Recorder into
// the request context.
func RequestLogger(logger logger.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logger.Info("%s %s", r.Method, r.URL.Path)

			s := statusCodeCollector{ResponseWriter: w}

			next.ServeHTTP(&s, r)

			logger.Info(
				"%s %s, Status %d, Duration %v",
				r.Method,
				r.URL.Path,
				s.StatusCode,
				timing.GetRecorder(r.Context()).Read(HandlerDuration),
			)
		},
	)
}
