package lang

var _ Value = (*valueNull)(nil) // ensure that valueNull implements Value

var (
	// Null represents the Null value as specified by the language spec
	Null Value = &valueNull{}
)

type valueNull struct{}

// Value returns nil.
func (valueNull) Value() interface{} { return nil }

// Type returns lang.TypeNull.
func (valueNull) Type() Type { return TypeNull }
