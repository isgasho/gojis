package lang

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRecordGetField(t *testing.T) {
	require := require.New(t)

	r := NewRecord()
	require.Len(r.fields, 0)

	val, ok := r.GetField("foobar")
	require.Nil(val)
	require.False(ok)
}

func TestRecordSetField(t *testing.T) {
	require := require.New(t)

	r := NewRecord()
	require.Len(r.fields, 0)

	r.SetField("foobar", struct{}{})
	val, ok := r.GetField("foobar")
	require.Equal(struct{}{}, val)
	require.True(ok)
}
