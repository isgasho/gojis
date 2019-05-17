package vm

var (
	True  Value = Boolean(true)
	False Value = Boolean(false)
)

type Boolean bool

func (b Boolean) Value() interface{} { return bool(b) }
func (Boolean) Type() Type           { return TypeBoolean }
