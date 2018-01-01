package create

type ValidatorFunc func(*Request) []error

func Validator(req *Request) []error {
	return []error{}
}
