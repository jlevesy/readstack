package index

import (
	"github.com/jlevesy/readstack/model"
)

// Response is an index.Handler result
// It carries all model.Item found by the handler
type Response struct {
	Items []*model.Item `json:"items"`
}
