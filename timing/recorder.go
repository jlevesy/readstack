package timing

import (
	"time"
)

// Recorder describes an object which can records duration metrics
type Recorder interface {
	Write(metric string, value time.Duration)
	Read(metric string) time.Duration
}

type inMemoryRecorder map[string]time.Duration

// NewInMemoryRecorder returns an instance of an inMemoryRecorder
func NewInMemoryRecorder() Recorder {
	return inMemoryRecorder{}
}

func (r inMemoryRecorder) Write(metric string, value time.Duration) {
	r[metric] = value
}

func (r inMemoryRecorder) Read(metric string) time.Duration {
	return r[metric]
}
