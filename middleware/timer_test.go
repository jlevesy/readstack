package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	timingStub "github.com/jlevesy/readstack/test/stub/timing"
	"github.com/jlevesy/readstack/timing"
)

func TestItInjectsATimeRecorder(t *testing.T) {
	var (
		calledWriter  http.ResponseWriter
		calledRequest *http.Request
		recorder      timing.Recorder
	)

	callee := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			calledWriter = w
			calledRequest = r
			recorder = timing.GetRecorder(r.Context())
		},
	)

	responseWriter := httptest.NewRecorder()

	request := httptest.NewRequest(
		"GET", "/bar/buz", bytes.NewBuffer([]byte{}),
	)

	subject := WithInMemoryTimingRecorder(callee)

	subject.ServeHTTP(responseWriter, request)

	if calledWriter == nil {
		t.Fatal("Expected callee to be called")
	}

	if calledRequest == nil {
		t.Fatal("Expected callee to be called")
	}

	if recorder == nil {
		t.Fatal("Expected to get a recorder, got nothing")
	}
}

func TestItRecordsDurationOfAHandler(t *testing.T) {
	var (
		calledWriter  http.ResponseWriter
		calledRequest *http.Request
		wroteMetric   string
		wroteDuration time.Duration
	)

	wroteDuration = 999

	metricName := "Groot"

	callee := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			calledWriter = w
			calledRequest = r
		},
	)

	recorder := timingStub.RecorderStub{
		OnWrite: func(metric string, duration time.Duration) {
			wroteMetric = metric
			wroteDuration = duration
		},
	}

	responseWriter := httptest.NewRecorder()

	request := httptest.NewRequest(
		"GET", "/bar/buz", bytes.NewBuffer([]byte{}),
	)
	request = request.WithContext(
		timing.WithRecorder(request.Context(), &recorder),
	)

	subject := RecordDuration(metricName, callee)

	subject.ServeHTTP(responseWriter, request)

	if calledWriter == nil {
		t.Fatal("Expected callee to be called")
	}

	if calledRequest == nil {
		t.Fatal("Expected callee to be called")
	}

	if wroteMetric != metricName {
		t.Fatalf("Expected wrote metric to be %s, got %s", metricName, wroteMetric)
	}

	if wroteDuration == 999 {
		t.Fatal("Expected wrote duration to be greater than 0")
	}
}
