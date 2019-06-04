package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPushPop(t *testing.T) {
	require := require.New(t)

	var s Stack
	var elem interface{}
	var ok bool

	s = s.Push("foo")
	s = s.Push("bar")

	s, elem, ok = s.Pop()
	require.True(ok)
	require.Equal("bar", elem)
	require.NotNil(s)

	s, elem, ok = s.Pop()
	require.True(ok)
	require.Equal("foo", elem)
	require.NotNil(s)

	s, elem, ok = s.Pop()
	require.False(ok)
	require.Nil(elem)
	require.NotNil(s)
}

func TestPopUninitialized(t *testing.T) {
	require := require.New(t)

	var s Stack
	var elem interface{}
	var ok bool

	s, elem, ok = s.Pop()
	require.False(ok)
	require.Nil(elem)
	require.Nil(s)
}
