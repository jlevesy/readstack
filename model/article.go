package model

import (
	"net/url"
)

type Article struct {
	ID   uint
	Name string
	URL  url.URL
}

func NewArticle(name string, URL url.URL) *Article {
	return &Article{name, URL}
}
