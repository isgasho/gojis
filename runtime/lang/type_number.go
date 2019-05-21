package lang

import (
	"math"
	"math/big"
)

var _ Value = (*Number)(nil)

var (
	NaN         Value = NewNumber(math.NaN())
	PosInfinity       = NewNumber(math.Inf(+1))
	NegInfinity       = NewNumber(math.Inf(-1))
	Infinity          = PosInfinity
	PosZero           = NewNumber(+0)
	NegZero           = NewNumber(neg(0))
	Zero              = PosZero
)

func neg(x float64) float64 {
	return -x
}

type Number struct {
	value *big.Float
	isNaN bool
}

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

func (n Number) Value() interface{} {
	if n.isNaN {
		return math.NaN()
	}

	val, _ := n.value.Float64()
	return val
}

func (Number) Type() Type { return TypeNumber }
