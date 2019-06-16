package stack

import "testing"

func BenchmarkSliceStack(b *testing.B) {
	benchmarkStack(b, NewSliceStack())
}
