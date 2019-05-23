package lang

import "testing"

func BenchmarkPropertyValue(b *testing.B) {
	p := NewDataProperty(NewString("foobar"), False, False, False)

	for i := 0; i < b.N; i++ {
		_ = p.Value()
	}
}

func BenchmarkPropertyEnumerable(b *testing.B) {
	p := NewDataProperty(NewString("foobar"), False, False, False)

	for i := 0; i < b.N; i++ {
		_ = p.Enumerable()
	}
}

func BenchmarkPropertyConfigurable(b *testing.B) {
	p := NewDataProperty(NewString("foobar"), False, False, False)

	for i := 0; i < b.N; i++ {
		_ = p.Configurable()
	}
}

func BenchmarkNewProperty(b *testing.B) {
	var p *Property
	for i := 0; i < b.N; i++ {
		p = NewProperty()
	}
	_ = p
}

func BenchmarkNewDataProperty(b *testing.B) {
	var p *Property
	for i := 0; i < b.N; i++ {
		p = NewDataProperty(Null, False, False, False)
	}
	_ = p
}

func BenchmarkIsAccessorDescriptor(b *testing.B) {
	p := NewDataProperty(Null, False, False, False)
	for i := 0; i < b.N; i++ {
		_ = p.IsAccessorDescriptor()
	}
}

func BenchmarkIsDataDescriptor(b *testing.B) {
	p := NewDataProperty(Null, False, False, False)
	for i := 0; i < b.N; i++ {
		_ = p.IsDataDescriptor()
	}
}

func BenchmarkIsGenericDescriptor(b *testing.B) {
	p := NewDataProperty(Null, False, False, False)
	for i := 0; i < b.N; i++ {
		_ = p.IsGenericDescriptor()
	}
}
