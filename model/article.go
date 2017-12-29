package model

import (
	"net/url"
)

type Article struct {
	Name string
	URL  url.URL
}

func NewArticle(name string, URL url.URL) *Article {
	return &Article{name, URL}
}
