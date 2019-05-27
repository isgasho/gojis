package lang

import "testing"

func BenchmarkUndefinedToBoolean(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToBoolean(Undefined)
	}
}

func BenchmarkNullToBoolean(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToBoolean(Null)
	}
}

func BenchmarkSymbolToBoolean(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToBoolean(SymbolToPrimitive)
	}
}

func BenchmarkObjectToBoolean(b *testing.B) {
	o := &Object{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToBoolean(o)
	}
}

func BenchmarkBooleanToBoolean(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToBoolean(True)
	}
}

func BenchmarkNumberToBoolean(b *testing.B) {
	n := NewNumber(1.4)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToBoolean(n)
	}
}

func BenchmarkStringToBoolean(b *testing.B) {
	s := NewString("foobar")

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToBoolean(s)
	}
}
