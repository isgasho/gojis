package stack

import (
	"testing"
)

func TestLinkedStack(t *testing.T) {
	testStack(t, NewLinkedStack())
}
