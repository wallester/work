package errors

import "github.com/juju/errors"

var ErrMethodIsNotAllowedInProduction = errors.New("method is not allowed in production")

// Last discards all arguments except the last one and returns it as an error, if
// it is not nil. If the last argument is nil, Last returns nil.
// Example:
//
//	_, err := Foo()
//	return err
//
// becomes
//
//	return Last(Foo())
func Last(values ...any) error {
	if len(values) == 0 {
		return nil
	}

	if err := values[len(values)-1]; err != nil {
		return err.(error)
	}

	return nil
}
