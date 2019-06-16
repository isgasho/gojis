package stack

import (
	"testing"
)

func TestSliceStack(t *testing.T) {
	testStack(t, NewSliceStack())
}
