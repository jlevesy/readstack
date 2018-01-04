package middleware

import (
	"context"
	"time"

	"net/http"
)

const (
	HandlerProbe = "handler"

	recorderKey = "recorder"
)

type Recorder interface {
	Write(metric string, value time.Duration)
	Read(metric string) time.Duration
}

type recorder map[string]time.Duration

func (r recorder) Write(metric string, value time.Duration) {
	r[metric] = value
}

func (r recorder) Read(metric string) time.Duration {
	return r[metric]
}

func GetTimeRecorder(context context.Context) Recorder {
	return context.Value(recorderKey).(Recorder)
}

func WithTimer(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			recorder := recorder{}

			next.ServeHTTP(
				w,
				r.WithContext(
					context.WithValue(
						r.Context(),
						recorderKey,
						recorder,
					),
				),
			)
		},
	)
}

func TimerProbe(probeName string, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			next.ServeHTTP(w, r)

			end := time.Now()

			GetTimeRecorder(r.Context()).Write(probeName, end.Sub(start))
		},
	)
}
