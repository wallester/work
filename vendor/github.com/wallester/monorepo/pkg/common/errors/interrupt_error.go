package errors

type InterruptedError struct{}

func (e *InterruptedError) Error() string {
	return "process interrupted"
}
