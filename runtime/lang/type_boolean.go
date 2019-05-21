package lang

var _ Value = (*Boolean)(nil)

// Available Boolean values
var (
	True  = Boolean(true)
	False = Boolean(false)
)

type Boolean bool

func (b Boolean) Value() interface{} { return bool(b) }
func (Boolean) Type() Type           { return TypeBoolean }
