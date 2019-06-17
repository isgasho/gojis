package stack

import "testing"

func BenchmarkLinkedStack(b *testing.B) {
	benchmarkStack(b, NewLinkedStack())
}
