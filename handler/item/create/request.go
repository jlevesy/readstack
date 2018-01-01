package create

type Request struct {
	Name string
	URL  string
}

func NewRequest(Name, URL string) *Request {
	return &Request{Name, URL}
}
