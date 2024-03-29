package errors

import "fmt"

type ErrorKind uint8

const (
	ErrorKindTypeError ErrorKind = iota
	ErrorKindReferenceError
	ErrorKindRangeError
)

type Error interface {
	error
	Kind() ErrorKind
}

var _ Error = (*errorImpl)(nil) // ensure that errorImpl implements Error

type errorImpl struct {
	error
	kind ErrorKind
}

func (e errorImpl) Kind() ErrorKind {
	return e.kind
}

func NewTypeError(msg string) Error {
	return errorImpl{
		error: fmt.Errorf("TypeError: %v", msg),
		kind:  ErrorKindTypeError,
	}
}

func NewReferenceError(msg string) Error {
	return errorImpl{
		error: fmt.Errorf("ReferenceError: %v", msg),
		kind:  ErrorKindReferenceError,
	}
}

func NewRangeError(msg string) Error {
	return errorImpl{
		error: fmt.Errorf("RangeError: %v", msg),
		kind:  ErrorKindRangeError,
	}
}
