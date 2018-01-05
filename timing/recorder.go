package timing

import (
	"time"
)

type Recorder interface {
	Write(metric string, value time.Duration)
	Read(metric string) time.Duration
}

type inMemoryRecorder map[string]time.Duration

func NewInMemoryRecorder() Recorder {
	return inMemoryRecorder{}
}

func (r inMemoryRecorder) Write(metric string, value time.Duration) {
	r[metric] = value
}

func (r inMemoryRecorder) Read(metric string) time.Duration {
	return r[metric]
}
