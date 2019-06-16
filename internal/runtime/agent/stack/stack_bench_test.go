package stack

import "testing"

var r interface{}

func benchmarkStack(b *testing.B, s Stack) {
	b.Run("Push", benchmarkPush(s))
	b.Run("Push Pop", benchmarkPushPop(s))
	b.Run("Peek", benchmarkPeek(s))
}

func benchmarkPush(s Stack) func(b *testing.B) {
	return func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			s.Push("foo")
		}
	}
}

func benchmarkPushPop(s Stack) func(b *testing.B) {
	return func(b *testing.B) {
		var elem interface{}
		val := "foo"

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			s.Push(val)
			elem = s.Pop()
		}

		r = elem
	}
}

func benchmarkPeek(s Stack) func(b *testing.B) {
	return func(b *testing.B) {
		var elem interface{}

		s.Push("foo")

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			elem = s.Peek()
		}

		r = elem
	}
}
