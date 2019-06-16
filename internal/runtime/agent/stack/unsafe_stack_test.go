package stack

import "testing"

func TestUnsafeStack(t *testing.T) {
	testStack(t, NewUnsafeStack())
}
