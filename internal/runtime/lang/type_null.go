package lang

var (
	_ Value = (*valueNull)(nil) // ensure that valueNull implements Value
	_ Value = Null              // ensure that Null can actually be used as a value
)

const (
	// Null represents the Null value as specified by the language spec
	Null = valueNull(0)
)

type valueNull uint8

// Value returns nil.
func (valueNull) Value() interface{} { return nil }

// Type returns lang.TypeNull.
func (valueNull) Type() Type { return TypeNull }
