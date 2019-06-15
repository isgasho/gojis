package lang

var (
	_ Value = (*valueUndefined)(nil) // ensure valueUndefined implements Value
	_ Value = Undefined              // ensure that Undefined can actually be used as a Value
)

const (
	// Undefined represents the Undefined value as specified by the language spec.
	Undefined = valueUndefined(0)
)

type valueUndefined uint8

// Value returns Undefined.
func (valueUndefined) Value() interface{} { return Undefined }

// Type returns lang.TypeUndefined.
func (valueUndefined) Type() Type { return TypeUndefined }
