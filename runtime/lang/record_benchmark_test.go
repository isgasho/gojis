package lang

import "testing"

func BenchmarkNewRecord(b *testing.B) {
	var r *Record
	for i := 0; i < b.N; i++ {
		r = NewRecord()
	}
	_ = r
}

func BenchmarkRecordGetField(b *testing.B) {
	b.Run("Positive", func(b *testing.B) {
		val := struct {
			id  string
			val string
		}{
			id:  "thisisunique",
			val: "snafu",
		}

		r := NewRecord()
		r.SetField("foobar", val)

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = r.GetField("foobar")
		}
	})

	b.Run("Negative", func(b *testing.B) {
		r := NewRecord()

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = r.GetField("foobar")
		}
	})
}

func BenchmarkRecordSetField(b *testing.B) {
	r := NewRecord()

	for i := 0; i < b.N; i++ {
		r.SetField("foobar", struct{}{})
	}
}
