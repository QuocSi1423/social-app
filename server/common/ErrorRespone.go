package common


type ErrorRespone struct {
	Error string
}

func NewErrorRespone(err string) ErrorRespone{
	return ErrorRespone{Error: err}
}