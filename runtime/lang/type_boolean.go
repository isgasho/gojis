package lang

var _ Value = (*Boolean)(nil) // ensure that Boolean implements Value

// Available Boolean values
var (
	True  = Boolean(true)
	False = Boolean(false)
)

// Boolean is a language type as specified by the language spec.
// Predefined and ready to use values are lang.True and lang.False.
type Boolean bool

// Value returns the Go value of this Boolean, either true or false.
func (b Boolean) Value() interface{} { return bool(b) }

// Type returns lang.TypeBoolean.
func (Boolean) Type() Type { return TypeBoolean }
