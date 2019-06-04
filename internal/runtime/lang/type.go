package lang

import "fmt"

// Type represents a language type as specified by the
// ECMAScript Language Types.
type Type uint8

// Language Type as specified in 6.1
const (
	TypeUndefined Type = iota
	TypeNull
	TypeBoolean
	TypeString
	TypeSymbol
	TypeNumber
	TypeObject

	TypeInternal
)

func (t Type) String() string {
	switch t {
	case TypeUndefined:
		return "Undefined"
	case TypeNull:
		return "Null"
	case TypeBoolean:
		return "Boolean"
	case TypeString:
		return "String"
	case TypeSymbol:
		return "Symbol"
	case TypeNumber:
		return "Number"
	case TypeObject:
		return "Object"
	default:
		return "Unknown"
	}
}

func EnsureTypeOneOf(arg Value, ts ...Type) {
	if !TypeIsOneOf(arg, ts...) {
		panic(fmt.Errorf("Value's type must be one of %v, but was %v", ts, arg.Type()))
	}
}

func TypeIsOneOf(arg Value, ts ...Type) bool {
	got := arg.Type()
	for _, t := range ts {
		if t == got {
			return true
		}
	}
	return false
}
