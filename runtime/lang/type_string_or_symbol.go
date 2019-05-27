package lang

var _ Value = (*StringOrSymbol)(nil) // ensure that StringOrSymbol implements Value

// StringOrSymbol is a wrapper around either a lang.String
// or a lang.Symbol.
// Upon creation with lang.NewStringOrSymbol, the passed
// value is checked to be a String or a Symbol.
// The Value methods Value() and Type() delegate to the
// wrapped String or Symbol.
type StringOrSymbol struct {
	underlying Value
}

// NewStringOrSymbol creates a new StringOrSymbol from a passed
// String or Symbol.
// This function will panic if the passed Value is not a String or a Symbol.
func NewStringOrSymbol(arg Value) StringOrSymbol {
	if arg.Type() != TypeString && arg.Type() != TypeSymbol {
		panic("Type of argument must be String or Symbol")
	}

	return StringOrSymbol{arg}
}

// Type returns lang.TypeString or lang.TypeSymbol, depending on the
// wrapped Type.
func (s StringOrSymbol) Type() Type { return s.underlying.Type() }

// Value will return the wrapped value's Value.
func (s StringOrSymbol) Value() interface{} { return s.underlying.Value() }

// String is a convenience method to convert the lang.StringOrSymbol
// to a lang.String.
func (s StringOrSymbol) String() String {
	if s.underlying.Type() == TypeSymbol {
		return s.underlying.(Symbol).String()
	}

	return s.underlying.(String)
}
