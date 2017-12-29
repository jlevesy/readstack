package model

import (
	"net/url"
)

type Item struct {
	ID   uint
	Name string
	URL  url.URL
}

func NewItem(name string, URL url.URL) *Item {
	return &Item{
		Name: name,
		URL:  URL,
	}
}
