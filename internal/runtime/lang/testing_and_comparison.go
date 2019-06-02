package lang

import (
	"fmt"
	"strings"

	"gitlab.com/gojis/vm/internal/runtime/errors"
)

func RequireObjectCoercible(arg Value) (Value, errors.Error) {
	switch arg.Type() {
	case TypeNull,
		TypeUndefined:
		return nil, errors.NewTypeError("Object is not coercible")
	case TypeBoolean,
		TypeNumber,
		TypeString,
		TypeSymbol,
		TypeObject:
		return arg, nil
	}

	panic(fmt.Errorf("Unhandled argument type: %v", arg.Type()))
}

func IsArray(arg Value) Boolean { return Boolean(InternalIsArray(arg)) }

func InternalIsArray(arg Value) bool {
	if arg.Type() != TypeObject {
		return false
	}

	panic("TODO: Arrays")
}

func IsCallable(arg Value) Boolean { return Boolean(InternalIsCallable(arg)) }

func InternalIsCallable(arg Value) bool {
	if arg.Type() != TypeObject {
		return false
	}

	if fObj, ok := arg.(*Object); ok {
		return fObj.Call != nil
	}
	return false
}

func IsConstructor(arg Value) Boolean { return Boolean(InternalIsConstructor(arg)) }

func InternalIsConstructor(arg Value) bool {
	if arg.Type() != TypeObject {
		return false
	}

	if cObj, ok := arg.(*Object); ok {
		return cObj.Construct != nil
	}
	return false
}

func IsExtensible(o *Object) Boolean { return Boolean(InternalIsExtensible(o)) }

func InternalIsExtensible(o *Object) bool {
	return o.IsExtensible().Value().(bool)
}

func IsInteger(arg Value) Boolean { return Boolean(InternalIsInteger(arg)) }

func InternalIsInteger(arg Value) bool {
	if arg.Type() != TypeNumber {
		return false
	}

	if arg == NaN ||
		arg == PosInfinity ||
		arg == NegInfinity {
		return false
	}

	val := arg.Value().(float64)
	return val == float64(int(val))
}

func IsPropertyKey(arg Value) Boolean { return Boolean(InternalIsPropertyKey(arg)) }

func InternalIsPropertyKey(arg Value) bool {
	return arg.Type() == TypeString ||
		arg.Type() == TypeSymbol
}

func IsRegExp(arg Value) Boolean { return Boolean(InternalIsRegExp(arg)) }

func InternalIsRegExp(arg Value) bool {
	if arg.Type() != TypeObject {
		return false
	}

	panic("TODO")
}

func IsStringPrefix(p, q String) Boolean { return Boolean(InternalIsStringPrefix(p, q)) }

func InternalIsStringPrefix(p, q String) bool {
	pVal, qVal := p.Value().(string), q.Value().(string)
	return strings.HasPrefix(qVal, pVal)
}

func SameValue(x, y Value) Boolean { return Boolean(InternalSameValue(x, y)) }

func InternalSameValue(x, y Value) bool {
	if x.Type() != y.Type() {
		return false
	}

	if x.Type() == TypeNumber {
		if x == NaN && y == NaN {
			return true
		}

		if x == PosZero && y == NegZero {
			return false
		}

		if x == NegZero && y == PosZero {
			return false
		}

		return x.Value() == y.Value()
	}

	return InternalSameValueNonNumber(x, y)
}

func SameValueZero(x, y Value) Boolean { return Boolean(InternalSameValueZero(x, y)) }

func InternalSameValueZero(x, y Value) bool {
	if x.Type() != y.Type() {
		return false
	}
	if x.Type() != y.Type() {
		return false
	}

	if x.Type() == TypeNumber {
		if x == NaN && y == NaN {
			return true
		}

		if x == PosZero && y == NegZero {
			return true
		}

		if x == NegZero && y == PosZero {
			return true
		}

		return x.Value() == y.Value()
	}

	return InternalSameValueNonNumber(x, y)
}

func SameValueNonNumber(x, y Value) Boolean {
	return Boolean(InternalSameValueNonNumber(x, y))
}

func InternalSameValueNonNumber(x, y Value) bool {
	if x.Type() == TypeUndefined {
		return true
	}

	if x.Type() == TypeNull {
		return true
	}

	if x.Type() == TypeString {
		xVal, yVal := x.Value().(String), y.Value().(String)
		return StringsEqual(xVal, yVal)
	}

	if x.Type() == TypeBoolean {
		return x.Value() == y.Value()
	}

	if x.Type() == TypeSymbol {
		return x.Value() == y.Value()
	}

	return x.Value() == y.Value()
}
