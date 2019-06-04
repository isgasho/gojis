package stack

import "testing"

var r interface{}

func BenchmarkPush(b *testing.B) {
	var s Stack

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s = s.Push("foo")
	}

	r = s
}

func BenchmarkPushPop(b *testing.B) {
	s := Stack(make([]interface{}, b.N))
	var elem interface{}
	var ok bool
	val := "foo"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s = s.Push(val)
		s, elem, ok = s.Pop()
	}

	_, r, _ = s, elem, ok
}

func BenchmarkPeek(b *testing.B) {
	var s Stack
	var elem interface{}
	var ok bool

	s = s.Push("foo")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		elem, ok = s.Peek()
	}

	r, _ = elem, ok
}
