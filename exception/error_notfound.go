package exception

type NotFoundError struct {
	Error string
}

func NewNotFoundeError(error string) NotFoundError {
	return NotFoundError{Error: error}
}
