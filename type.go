package vm

// Type represents a language type as specified by the
// ECMAScript Language Types.
type Type uint8

// Language Type as specified in 6.1
const (
	TypeUndefined = iota
	TypeNull
	TypeBoolean
	TypeString
	TypeSymbol
	TypeNumber
	TypeObject
)
