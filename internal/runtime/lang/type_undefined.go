package lang

var _ Value = (*valueUndefined)(nil) // ensure valueUndefined implements Value

var (
	// Undefined represents the Undefined value as specified by the language spec.
	Undefined Value = &valueUndefined{}
)

type valueUndefined struct{}

// Value returns Undefined.
func (valueUndefined) Value() interface{} { return Undefined }

// Type returns lang.TypeUndefined.
func (valueUndefined) Type() Type { return TypeUndefined }
