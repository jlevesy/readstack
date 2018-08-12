package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	timingStub "github.com/jlevesy/readstack/test/stub/timing"
	"github.com/jlevesy/readstack/timing"

	"github.com/jlevesy/readstack/test/stub/logger"
)

func TestItLogsAnInfoOnRequest(t *testing.T) {
	var (
		calledWriter  http.ResponseWriter
		calledRequest *http.Request
		readMetric    string
	)

	logCalledCount := 0

	callee := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			calledWriter = w
			calledRequest = r
		},
	)

	responseWriter := httptest.NewRecorder()

	recorder := timingStub.RecorderStub{
		OnRead: func(metric string) time.Duration {
			readMetric = metric
			return 200 * time.Millisecond
		},
	}

	request := httptest.NewRequest(
		"GET", "/bar/buz", bytes.NewBuffer([]byte{}),
	)
	request = request.WithContext(
		timing.WithRecorder(request.Context(), &recorder),
	)

	logger := &logger.LoggerStub{
		OnInfo: func(string, ...interface{}) {
			logCalledCount++
		},
	}

	subject := RequestLogger(logger, callee)

	subject.ServeHTTP(responseWriter, request)

	if calledRequest == nil {
		t.Error("Expected middleware to forward the request")
	}

	if calledWriter == nil {
		t.Error("Expected writer to be called, got nothing")
	}

	if logCalledCount != 2 {
		t.Errorf("Expected log.Info to be call exactly 2 times, got %d", logCalledCount)
	}

	if readMetric != HandlerDuration {
		t.Errorf("Expected logger to read metric HandlerDuration, got %s", readMetric)
	}
}
