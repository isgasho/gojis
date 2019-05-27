package lang

import (
	"math"
	"math/big"
)

var _ Value = (*Number)(nil) // ensure that Number implements Value

var (
	// NaN as specified by the language spec.
	NaN = NewNumber(math.NaN())
	// PosInfinity as specified by the language spec.
	PosInfinity = NewNumber(math.Inf(+1))
	// NegInfinity as specified by the language spec.
	NegInfinity = NewNumber(math.Inf(-1))
	// Infinity as specified by the language spec.
	// This is an alias for PosInfinity.
	Infinity = PosInfinity
	// PosZero  as specified by the language spec.
	PosZero = NewNumber(+0)
	// NegZero as specified by the language spec.
	NegZero = NewNumber(neg(0))
	// Zero as specified by the language spec.
	// This is an alias for PosZero.
	Zero = PosZero
)

// neg is required to trick Go into making a zero actually negative.
// I don't really get why this is required. It works though.
func neg(x float64) float64 {
	return -x
}

// Number is a language type as specified by the language spec.
type Number struct {
	value *big.Float
	isNaN bool
}

// NewNumber generates a number language Value from a float64.
// If the given float64 value is not a number (NaN), the returned
// Number will represent the NaN value as specified by the language spec.
func NewNumber(x float64) Number {
	if math.IsNaN(x) {
		return Number{
			isNaN: true,
		}
	}

	return Number{
		value: big.NewFloat(x),
	}
}

// Value returns the float64 value of the Number.
// If the Number is NaN, math.NaN() will be returned.
func (n Number) Value() interface{} {
	if n.isNaN {
		return math.NaN()
	}

	val, _ := n.value.Float64()
	return val
}

// Type returns lang.TypeNumber.
func (Number) Type() Type { return TypeNumber }
