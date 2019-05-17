package vm

var (
	Null Value = &valueNull{}
)

type valueNull struct{}

func (valueNull) Value() interface{} { return nil }
func (valueNull) Type() Type         { return TypeNull }
