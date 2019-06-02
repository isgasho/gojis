package errors

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTypeError(t *testing.T) {
	msg := "this is the message"
	require := require.New(t)

	err := NewTypeError(msg)
	require.Equal(ErrorKindTypeError, err.Kind())
	require.Equal("TypeError: "+msg, err.Error())

	err = NewReferenceError(msg)
	require.Equal(ErrorKindReferenceError, err.Kind())
	require.Equal("ReferenceError: "+msg, err.Error())
}
