package create

// Request is the create.Handler request type
// It carries all information needed in order to perform
// the create item action.
type Request struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// NewRequest returns a new request instance
func NewRequest(Name, URL string) *Request {
	return &Request{
		Name: Name,
		URL:  URL,
	}
}
