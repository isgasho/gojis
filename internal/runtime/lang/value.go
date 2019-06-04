package lang

// Value represents a language value as specified by the language spec.
type Value interface {
	// Type returns the language type of this value as specified by the language spec.
	Type() Type
	// Value returns the Go representation of this value.
	// For example, the Value() of a lang.Boolean should be a bool.
	Value() interface{}
}

type InternalValue interface {
	Value
}
