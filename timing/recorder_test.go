package timing

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestItCanStoreAndRetrievesMetrics(t *testing.T) {
	r := NewInMemoryRecorder()

	duration := 200 * time.Millisecond
	metricName := "metric"

	r.Write(metricName, duration)

	readValue := r.Read(metricName)
	if readValue != duration {
		t.Fatalf("Expected %v got %v", duration, readValue)
	}
}

func TestItIsSafeToUseWithANilContext(t *testing.T) {
	ctx := context.Background()

	rec := GetRecorder(ctx)

	if rec != nil {
		t.Fatal("We weren't expecting a value")
	}
}

func TestItCanRetrieveARecorder(t *testing.T) {
	ctx := context.Background()
	r := NewInMemoryRecorder()

	ctx2 := WithRecorder(ctx, r)

	rec := GetRecorder(ctx2)

	if !reflect.DeepEqual(rec, r) {
		t.Fatal("We were expecting to fetch the recorder")
	}
}
