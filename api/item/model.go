package item

type Model struct {
	id int64

	Name string
	URL  string
}

func New(name, URL string) *Model {
	return &Model{
		Name: name,
		URL:  URL,
	}
}

func (m *Model) GetID() int64 {
	return m.id
}
