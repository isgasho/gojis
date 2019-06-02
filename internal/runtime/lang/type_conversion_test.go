package lang

import (
	"math"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
	"gitlab.com/gojis/vm/internal/runtime/errors"
)

type unknownType struct{}

func (unknownType) Type() Type         { return Type(math.MaxUint8) }
func (unknownType) Value() interface{} { return nil }

func requirePanic(t *testing.T) {
	var panicked bool

	if err := recover(); err != nil {
		panicked = true
	}

	if !panicked {
		require.FailNow(t, "Test should have panicked")
	}
}

func TestCheckForPanic(t *testing.T) {
	defer requirePanic(t)

	panic("foo")
}

func TestToBooleanUnhandledType(t *testing.T) {
	defer requirePanic(t)
	_ = ToBoolean(&unknownType{})
}

func TestToIntegerUnhandledType(t *testing.T) {
	defer requirePanic(t)
	_, _ = ToInteger(&unknownType{})
}

func TestToInt32UnhandledType(t *testing.T) {
	defer requirePanic(t)
	_ = ToInt32(&unknownType{})
}

func TestToUint32UnhandledType(t *testing.T) {
	defer requirePanic(t)
	_ = ToUint32(&unknownType{})
}

func TestToInt16UnhandledType(t *testing.T) {
	defer requirePanic(t)
	_ = ToInt16(&unknownType{})
}

func TestToUint16UnhandledType(t *testing.T) {
	defer requirePanic(t)
	_ = ToUint16(&unknownType{})
}

func TestToInt8UnhandledType(t *testing.T) {
	defer requirePanic(t)
	_ = ToInt8(&unknownType{})
}

func TestToUint8UnhandledType(t *testing.T) {
	defer requirePanic(t)
	_ = ToUint8(&unknownType{})
}

func TestToUint8ClampUnhandledType(t *testing.T) {
	defer requirePanic(t)
	_ = ToUint8Clamp(&unknownType{})
}

func TestToBoolean(t *testing.T) {
	tests := []struct {
		name string
		arg  Value
		want Boolean
	}{
		{"True", True, True},
		{"False", False, False},
		{"Undefined", Undefined, False},
		{"Null", Null, False},
		{"Symbol", Symbol{NewString("")}, True},
		{"Symbol", Symbol{NewString("not empty string")}, True},
		{"Symbol", Symbol{NewString("@#$%RTYHUYVGFCDE$%R^TUY")}, True},
		{"Object", &Object{}, True},
		{"Boolean", Boolean(true), True},
		{"Boolean", Boolean(false), False},
		{"Number PosZero", PosZero, False},
		{"Number NegZero", NegZero, False},
		{"Number Zero", Zero, False},
		{"Number 0.0", NewNumber(0.0), False},
		{"Number NaN", NaN, False},
		{"Number math.NaN()", NewNumber(math.NaN()), False},
		{"Number PosInfinity", PosInfinity, True},
		{"Number NegInfinity", NegInfinity, True},
		{"Number Infinity", Infinity, True},
		// Number to Boolean conversion is tested in TestToBooleanFuzzy
		{"Empty String", NewString(""), False},
		{"String", NewString("OI*&I^FUTY N&*^IHF"), True},
		{"String", NewString(" "), True},
		{"Empty StringOrSymbol String", NewStringOrSymbol(NewString("")), False},
		{"Empty StringOrSymbol Symbol", NewStringOrSymbol(Symbol{NewString("")}), True},
		{"StringOrSymbol", NewStringOrSymbol(SymbolToPrimitive), True},
		{"StringOrSymbol", NewStringOrSymbol(NewString("foobar")), True},
		{"StringOrSymbol", NewStringOrSymbol(NewString(" ")), True},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, ToBoolean(tt.arg))
		})
	}
}

func TestToNumber(t *testing.T) {
	tests := []struct {
		name    string
		arg     Value
		want    Number
		wantErr errors.Error
	}{
		{"Undefined", Undefined, NaN, nil},
		{"Null", Null, PosZero, nil},
		{"True", True, NewNumber(1), nil},
		{"False", False, PosZero, nil},
		{"Number PosZero", PosZero, PosZero, nil},
		{"Number NegZero", NegZero, NegZero, nil},
		{"Number Zero", Zero, Zero, nil},
		{"Number 0.0", NewNumber(0.0), NewNumber(0.0), nil},
		{"Number NaN", NaN, NaN, nil},
		{"Number math.NaN()", NewNumber(math.NaN()), NewNumber(math.NaN()), nil},
		{"Number PosInfinity", PosInfinity, PosInfinity, nil},
		{"Number NegInfinity", NegInfinity, NegInfinity, nil},
		{"Number Infinity", Infinity, Infinity, nil},
		// Number to Number conversion is tested in TestToNumberFuzzy
		// String to Number is to be implemented and will panic
		{"Symbol", SymbolToPrimitive, Zero, errors.NewTypeError("Cannot convert from Symbol to Number")},
		// TODO: Object to Number conversion
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)

			got, err := ToNumber(tt.arg)

			require.Equal(tt.want, got)
			require.Equal(tt.wantErr, err)
		})
	}
}

func TestToBooleanFuzzy(t *testing.T) {
	require := require.New(t)

	conv := func(x float64) Boolean {
		n := NewNumber(x)
		return ToBoolean(n)
	}
	constTrue := func(float64) Boolean { return True }

	require.NoError(quick.CheckEqual(conv, constTrue, nil))
}

func TestToNumberFuzzy(t *testing.T) {
	require := require.New(t)

	conv := func(x float64) Number {
		n := NewNumber(x)
		val, err := ToNumber(n)
		require.NoError(err)
		return val
	}
	require.NoError(quick.CheckEqual(conv, NewNumber, nil))
}

func TestToIntegerFuzzy(t *testing.T) {
	require := require.New(t)

	conv := func(x float64) Number {
		n := NewNumber(x)
		val, err := ToInteger(n)
		require.NoError(err)
		return val
	}
	expected := func(x float64) Number {
		return NewNumber(math.Floor(x))
	}
	require.NoError(quick.CheckEqual(conv, expected, nil))
}
