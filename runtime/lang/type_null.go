package lang

var _ Value = (*valueNull)(nil)

var (
	Null Value = &valueNull{}
)

type valueNull struct{}

func (valueNull) Value() interface{} { return nil }
func (valueNull) Type() Type         { return TypeNull }
