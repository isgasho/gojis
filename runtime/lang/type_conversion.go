package lang

import (
	"fmt"

	"gitlab.com/gojis/vm/runtime/errors"
)

func ToPrimitive(input Value, preferredType interface{}) (Value, errors.Error) {
	if input.Type() == TypeObject {
		var hint string

		t, ok := preferredType.(Type)
		if preferredType != nil && !ok {
			panic("preferredType is not a Type")
		}

		if preferredType == nil {
			hint = "default"
		} else if t == TypeString {
			hint = "string"
		} else if t == TypeNumber {
			hint = "number"
		}

		o, ok := input.(*Object)
		if !ok {
			panic("input is TypeObject, but not *Object")
		}

		exoticToPrim, err := GetMethod(o, NewStringOrSymbol(SymbolToPrimitive))
		if err != nil {
			return nil, err
		}

		if exoticToPrim != Undefined {
			result, err := Call(exoticToPrim.(*Object), input, NewString(hint))
			if err != nil {
				return nil, err
			}

			if result.Type() != TypeObject {
				return result, nil
			}

			return nil, errors.NewTypeError("Call of internal primitive conversion returned non-primitive object")
		}

		if hint == "default" {
			hint = "number"
		}

		val, err := OrdinaryToPrimitive(o, hint)
		if err != nil {
			return nil, err
		}

		return val, nil
	}

	return input, nil
}

func ToBoolean(arg Value) Boolean {
	switch arg.Type() {
	case TypeUndefined,
		TypeNull:
		return False
	case TypeSymbol,
		TypeObject:
		return True
	case TypeBoolean:
		return arg.(Boolean)
	case TypeNumber:
		if val := arg.Value(); val == PosZero ||
			val == NegZero ||
			val == NaN {
			return False
		}
		return True
	case TypeString:
		if arg.Value() == "" {
			return False
		}
		return True
	}

	panic(fmt.Errorf("Unhandled type in conversion ToBoolean, type '%v'", arg.Type()))
}

func ToNumber(arg Value) Number {
	panic("TODO")
}

func ToInteger(arg Value) Number {
	panic("TODO")
}

func ToInt32(arg Value) Number {
	panic("TODO")
}

func ToUint32(arg Value) Number {
	panic("TODO")
}

func ToInt16(arg Value) Number {
	panic("TODO")
}

func ToUint16(arg Value) Number {
	panic("TODO")
}

func ToInt8(arg Value) Number {
	panic("TODO")
}

func ToUint8(arg Value) Number {
	panic("TODO")
}

func ToUint8Clamp(arg Value) Number {
	panic("TODO")
}

func ToString(arg Value) String {
	panic("TODO")
}

func NumberToString() String {
	panic("TODO")
}

func ToObject(arg Value) *Object {
	panic("TODO")
}

func ToPropertyKey(arg Value) StringOrSymbol {
	panic("TODO")
}

func ToLength(arg Value) Number {
	panic("TODO")
}

func CanonicalNumericIndexString(arg Value) Number {
	panic("TODO")
}

func ToIndex(arg Value) Number {
	panic("TODO")
}

func OrdinaryToPrimitive(o *Object, hint string) (Value, errors.Error) {
	methodNames := []string{"valueOf", "toString"}
	if hint == "string" {
		methodNames = []string{"toString", "valueOf"}
	}

	for _, name := range methodNames {
		method := Get(o, NewStringOrSymbol(NewString(name)))
		if IsCallable(method) {
			result, err := Call(method.(*Object), o)
			if err != nil {
				return nil, err
			}

			if result.Type() != TypeObject {
				return result, nil
			}
		}
	}

	return nil, errors.NewTypeError("Cannot convert ordinary object to primitive")
}
