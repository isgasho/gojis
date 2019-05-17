package vm

var (
	Undefined Value = &valueUndefined{}
)

type valueUndefined struct{}

func (valueUndefined) Value() interface{} { return "undefined" }
func (valueUndefined) Type() Type         { return TypeUndefined }
