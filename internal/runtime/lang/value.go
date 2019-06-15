package lang

// Value represents a language value as specified by the language spec.
type Value interface {
	// Type returns the language type of this value as specified by the language spec.
	Type() Type
	// Value returns the Go representation of this value.
	// For example, the Value() of a lang.Boolean should be a bool.
	Value() interface{}
}

// InternalValue represents any value used in the specification.
// This interface was introduced because the specification
// also uses values like Null and Undefined for non language values
// like Environments or Properties.
// Objects that implement this interface must implement a Type method
// that returns TypeInternal.
type InternalValue interface {
	Value
}
