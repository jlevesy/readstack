package model

type Item struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func NewItem(name, URL string) *Item {
	return &Item{
		Name: name,
		URL:  URL,
	}
}
