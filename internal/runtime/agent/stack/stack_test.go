package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func testStack(t *testing.T, s Stack) {
	t.Run("Push Pop", testPushPop(s))
	t.Run("Peek", testPeek(s))
}

func testPushPop(s Stack) func(t *testing.T) {
	return func(t *testing.T) {
		require := require.New(t)

		require.Nil(s.Pop())

		s.Push("foo")
		s.Push("bar")

		require.Equal("bar", s.Pop())
		require.Equal("foo", s.Pop())
		require.Nil(s.Pop())
		require.Nil(s.Pop()) // ensure that multiple calls to Pop don't fail
	}
}

func testPeek(s Stack) func(t *testing.T) {
	return func(t *testing.T) {
		require := require.New(t)

		require.Nil(s.Peek())
		require.Nil(s.Peek()) // ensure multiple calls to Peek don't fail

		s.Push("foo")
		s.Push("bar")
		s.Push("snafu")

		require.Equal("snafu", s.Peek())
		require.Equal("snafu", s.Peek()) // ensure multiple calls to Peek don't Pop

		_ = s.Pop()

		require.Equal("bar", s.Peek()) // ensure a Pop changes Peek

		_ = s.Pop()
		_ = s.Pop()

		require.Nil(s.Peek())
	}
}
