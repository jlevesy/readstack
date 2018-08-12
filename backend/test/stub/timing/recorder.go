package timing

import (
	"time"
)

// RecorderStub is a stub for the timing.Recorder interface
type RecorderStub struct {
	OnWrite func(string, time.Duration)
	OnRead  func(string) time.Duration
}

func (r *RecorderStub) Write(metric string, value time.Duration) {
	r.OnWrite(metric, value)
}

func (r *RecorderStub) Read(metric string) time.Duration {
	return r.OnRead(metric)
}
