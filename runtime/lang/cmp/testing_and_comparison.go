package cmp

import (
	"fmt"
	"strings"

	"gitlab.com/gojis/vm/runtime/errors"
	"gitlab.com/gojis/vm/runtime/lang"
)

func RequireObjectCoercible(arg lang.Value) (lang.Value, errors.Error) {
	switch arg.Type() {
	case lang.TypeNull,
		lang.TypeUndefined:
		return nil, errors.NewTypeError("Object is not coercible")
	case lang.TypeBoolean,
		lang.TypeNumber,
		lang.TypeString,
		lang.TypeSymbol,
		lang.TypeObject:
		return arg, nil
	}

	panic(fmt.Errorf("Unhandled argument type: %v", arg.Type()))
}

func IsArray(arg lang.Value) lang.Boolean { return lang.Boolean(InternalIsArray(arg)) }

func InternalIsArray(arg lang.Value) bool {
	if arg.Type() != lang.TypeObject {
		return false
	}

	panic("TODO: Arrays")
}

func IsCallable(arg lang.Value) lang.Boolean { return lang.Boolean(InternalIsCallable(arg)) }

func InternalIsCallable(arg lang.Value) bool {
	if arg.Type() != lang.TypeObject {
		return false
	}

	if fObj, ok := arg.(*lang.FunctionObject); ok {
		return fObj.Call != nil
	}
	return false
}

func IsConstructor(arg lang.Value) lang.Boolean { return lang.Boolean(InternalIsConstructor(arg)) }

func InternalIsConstructor(arg lang.Value) bool {
	if arg.Type() != lang.TypeObject {
		return false
	}

	if cObj, ok := arg.(*lang.ConstructorFunctionObject); ok {
		return cObj.Construct != nil
	}
	return false
}

func IsExtensible(o *lang.Object) lang.Boolean { return lang.Boolean(InternalIsExtensible(o)) }

func InternalIsExtensible(o *lang.Object) bool {
	return o.IsExtensible().Value().(bool)
}

func IsInteger(arg lang.Value) lang.Boolean { return lang.Boolean(InternalIsInteger(arg)) }

func InternalIsInteger(arg lang.Value) bool {
	if arg.Type() != lang.TypeNumber {
		return false
	}

	if arg == lang.NaN ||
		arg == lang.PosInfinity ||
		arg == lang.NegInfinity {
		return false
	}

	val := arg.Value().(float64)
	return val == float64(int(val))
}

func IsPropertyKey(arg lang.Value) lang.Boolean { return lang.Boolean(InternalIsPropertyKey(arg)) }

func InternalIsPropertyKey(arg lang.Value) bool {
	return arg.Type() == lang.TypeString ||
		arg.Type() == lang.TypeSymbol
}

func IsRegExp(arg lang.Value) lang.Boolean { return lang.Boolean(InternalIsRegExp(arg)) }

func InternalIsRegExp(arg lang.Value) bool {
	if arg.Type() != lang.TypeObject {
		return false
	}

	panic("TODO")
}

func IsStringPrefix(p, q lang.String) lang.Boolean { return lang.Boolean(InternalIsStringPrefix(p, q)) }

func InternalIsStringPrefix(p, q lang.String) bool {
	pVal, qVal := p.Value().(string), q.Value().(string)
	return strings.HasPrefix(qVal, pVal)
}

func SameValue(x, y lang.Value) lang.Boolean { return lang.Boolean(InternalSameValue(x, y)) }

func InternalSameValue(x, y lang.Value) bool {
	if x.Type() != y.Type() {
		return false
	}

	if x.Type() == lang.TypeNumber {
		if x == lang.NaN && y == lang.NaN {
			return true
		}

		if x == lang.PosZero && y == lang.NegZero {
			return false
		}

		if x == lang.NegZero && y == lang.PosZero {
			return false
		}

		return x.Value() == y.Value()
	}

	return InternalSameValueNonNumber(x, y)
}

func SameValueZero(x, y lang.Value) lang.Boolean { return lang.Boolean(InternalSameValueZero(x, y)) }

func InternalSameValueZero(x, y lang.Value) bool {
	if x.Type() != y.Type() {
		return false
	}
	if x.Type() != y.Type() {
		return false
	}

	if x.Type() == lang.TypeNumber {
		if x == lang.NaN && y == lang.NaN {
			return true
		}

		if x == lang.PosZero && y == lang.NegZero {
			return true
		}

		if x == lang.NegZero && y == lang.PosZero {
			return true
		}

		return x.Value() == y.Value()
	}

	return InternalSameValueNonNumber(x, y)
}

func SameValueNonNumber(x, y lang.Value) lang.Boolean {
	return lang.Boolean(InternalSameValueNonNumber(x, y))
}

func InternalSameValueNonNumber(x, y lang.Value) bool {
	if x.Type() == lang.TypeUndefined {
		return true
	}

	if x.Type() == lang.TypeNull {
		return true
	}

	if x.Type() == lang.TypeString {
		xVal, yVal := x.Value().(lang.String), y.Value().(lang.String)
		return lang.StringsEqual(xVal, yVal)
	}

	if x.Type() == lang.TypeBoolean {
		return x.Value() == y.Value()
	}

	if x.Type() == lang.TypeSymbol {
		return x.Value() == y.Value()
	}

	return x.Value() == y.Value()
}
