package create

type Request struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func NewRequest(Name, URL string) *Request {
	return &Request{
		Name: Name,
		URL:  URL,
	}
}
