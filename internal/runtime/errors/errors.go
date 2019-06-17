package errors

import "fmt"

// ErrorKind is the type of error an error can have.
type ErrorKind uint8

// Available error types used in the specification.
const (
	ErrorKindTypeError ErrorKind = iota
	ErrorKindReferenceError
	ErrorKindRangeError
)

// Error is an error that can be thrown during runtime.
// When implementing the specification, maybe this layer will be
// reduced to the ThrowXXXError function objects.
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

// NewTypeError creates a new type error with the given error.
func NewTypeError(msg string) Error {
	return errorImpl{
		error: fmt.Errorf("TypeError: %v", msg),
		kind:  ErrorKindTypeError,
	}
}

// NewReferenceError creates a new reference error with the given error.
func NewReferenceError(msg string) Error {
	return errorImpl{
		error: fmt.Errorf("ReferenceError: %v", msg),
		kind:  ErrorKindReferenceError,
	}
}

// NewRangeError creates a new range error with the given error.
func NewRangeError(msg string) Error {
	return errorImpl{
		error: fmt.Errorf("RangeError: %v", msg),
		kind:  ErrorKindRangeError,
	}
}
