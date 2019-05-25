package lang

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBoolean(t *testing.T) {
	require := require.New(t)

	_f := Boolean(false)
	_t := Boolean(true)

	require.Equal(True, _t)
	require.Equal(true, _t.Value())
	if !_t {
		require.FailNow("_t must be true")
	}

	require.Equal(False, _f)
	require.Equal(false, _f.Value())
	if _f {
		require.FailNow("_f must be false")
	}
}

func TestBooleanConstants(t *testing.T) {
	require := require.New(t)

	require.Equal(true, True.Value())
	if !True {
		require.FailNow("True must be true")
	}

	require.Equal(false, False.Value())
	if False {
		require.FailNow("False must be false")
	}
}
