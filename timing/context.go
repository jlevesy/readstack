package timing

import (
	"context"
)

const (
	recorderKey recorderKeyType = "recorder"
)

type recorderKeyType string

// GetRecorder fetches recorder from given context.
func GetRecorder(context context.Context) Recorder {
	rec := context.Value(recorderKey)

	if rec == nil {
		return nil
	}

	return rec.(Recorder)
}

// WithRecorder add given recorder to the given context
func WithRecorder(ctx context.Context, r Recorder) context.Context {
	return context.WithValue(ctx, recorderKey, r)
}
