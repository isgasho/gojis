package lang

var _ Value = (*valueUndefined)(nil)

var (
	Undefined Value = &valueUndefined{}
)

type valueUndefined struct{}

func (valueUndefined) Value() interface{} { return Undefined }
func (valueUndefined) Type() Type         { return TypeUndefined }
