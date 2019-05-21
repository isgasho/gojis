package errors

import "fmt"

type ErrorKind uint8

const (
	ErrorKindTypeError ErrorKind = iota
	ErrorKindReferenceError
)

type Error interface {
	error
	Kind() ErrorKind
}

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
