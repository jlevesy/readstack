package timing

import (
	"context"
)

const (
	recorderKey = "recorder"
)

func GetRecorder(context context.Context) Recorder {
	return context.Value(recorderKey).(Recorder)
}

func WithRecorder(ctx context.Context, r Recorder) context.Context {
	return context.WithValue(ctx, recorderKey, r)
}
