package model

type Item struct {
	ID   uint
	Name string
	URL  string
}

func NewItem(name, URL string) *Item {
	return &Item{
		Name: name,
		URL:  URL,
	}
}
