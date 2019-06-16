package stack

import "testing"

func BenchmarkUnsafeStack(b *testing.B) {
	benchmarkStack(b, NewUnsafeStack())
}
