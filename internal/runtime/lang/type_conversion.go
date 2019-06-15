package lang

import (
	"fmt"
	"math"

	"github.com/TimSatke/gojis/internal/runtime/errors"
)

// ToPrimitive converts a Value to a primitive type according to 7.1.1 in
// the specification.
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

// ToBoolean converts a Value to a Boolean according to 7.1.2 in the specification.
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
		if val := arg.Value(); val == PosZero.Value() ||
			val == NegZero.Value() ||
			arg == NaN {
			return False
		}
		return True
	case TypeString:
		if arg.Value() == "" {
			return False
		}
		return True
	}

	panic(unhandledType(arg))
}

// ToNumber converts a Value to a Number according to 7.1.3 in the specification.
func ToNumber(arg Value) (Number, errors.Error) {
	switch arg.Type() {
	case TypeUndefined:
		return NaN, nil
	case TypeNull:
		return PosZero, nil
	case TypeBoolean:
		if arg.(Boolean) {
			return NewNumber(1), nil
		}
		return PosZero, nil
	case TypeNumber:
		return arg.(Number), nil
	case TypeString:
		panic("TODO: 7.1.3.1")
	case TypeSymbol:
		return Zero, errors.NewTypeError("Cannot convert from Symbol to Number")
	case TypeObject:
		primValue, err := ToPrimitive(arg, TypeNumber)
		if err != nil {
			return Zero, err
		}
		return ToNumber(primValue)
	}

	panic(unhandledType(arg))
}

// ToInteger converts a Value to a Number whose value is an integer, according to
// 7.1.4 in the specification.
func ToInteger(arg Value) (Number, errors.Error) {
	number, err := ToNumber(arg)
	if err != nil {
		return Zero, err
	}

	if number == NaN {
		return PosZero, nil
	}

	val := arg.Value()

	if val == PosZero.Value() ||
		val == NegZero.Value() ||
		arg == PosInfinity ||
		arg == NegInfinity {
		return arg.(Number), nil
	}

	return NewNumber(math.Floor(val.(float64))), nil
}

// ToInt32 converts a Value to a Number whose value is an int32
// according to 7.1.5 in the specification.
func ToInt32(arg Value) Number {
	panic("TODO")
}

// ToUint32 converts a Value to a Number whose value is an uint32
// according to 7.1.6 in the specification.
func ToUint32(arg Value) Number {
	panic("TODO")
}

// ToInt16 converts a Value to a Number whose value is an int16
// according to 7.1.7 in the specification.
func ToInt16(arg Value) Number {
	panic("TODO")
}

// ToUint16 converts a Value to a Number whose value is an uint16
// according to 7.1.8 in the specification.
func ToUint16(arg Value) Number {
	panic("TODO")
}

// ToInt8 converts a Value to a Number whose value is an int8
// according to 7.1.9 in the specification.
func ToInt8(arg Value) Number {
	panic("TODO")
}

// ToUint8 converts a Value to a Number whose value is an uint8
// according to 7.1.10 in the specification.
func ToUint8(arg Value) Number {
	panic("TODO")
}

// ToUint8Clamp converts a Value to a Number whose value is an uint8
// according to 7.1.11 in the specification.
func ToUint8Clamp(arg Value) Number {
	panic("TODO")
}

// ToString converts a Value to a String according to 7.1.12 in the
// specification.
func ToString(arg Value) String {
	panic("TODO")
}

// NumberToString converts a Number to a String according to 7.1.12.1
// in the specification.
func NumberToString(m Number) String {
	panic("TODO")
}

// ToObject converts a Value to an object according to 7.1.13 in the
// specification.
func ToObject(arg Value) *Object {
	panic("TODO")
}

// ToPropertyKey converts a Value to a property key (this implementation
// represents property keys with the type StringOrSymbol) according
// to 7.1.14 in the specification.
func ToPropertyKey(arg Value) StringOrSymbol {
	panic("TODO")
}

// ToLength converts a Value to a Number whose value is an integer
// according to 7.1.15 in the specification.
func ToLength(arg Value) Number {
	panic("TODO")
}

// CanonicalNumericIndexString converts a Value to a Number according
// to 7.1.16 in the specification.
func CanonicalNumericIndexString(arg Value) Number {
	panic("TODO")
}

// ToIndex converts a Value to a Number according to 7.1.17 in the specification.
func ToIndex(arg Value) Number {
	panic("TODO")
}

// OrdinaryToPrimitive converts an Object to a value, respecting the
// given hint, according to 7.1.1.1 in the specification.
func OrdinaryToPrimitive(o *Object, hint string) (Value, errors.Error) {
	methodNames := []string{"valueOf", "toString"}
	if hint == "string" {
		methodNames = []string{"toString", "valueOf"}
	}

	for _, name := range methodNames {
		method, err := Get(o, NewStringOrSymbol(NewString(name)))
		if err != nil {
			return nil, err
		}

		if InternalIsCallable(method) {
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

// unhandledType returns an error saying "Unhandled type in type conversion: '<arg.Type()>'".
func unhandledType(arg Value) error {
	return fmt.Errorf("Unhandled type in type conversion: '%v'", arg.Type())
}
