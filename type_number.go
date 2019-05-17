package vm

import (
	"math"
	"math/big"
)

var (
	NaN         Value = NewNumber(math.NaN())
	PosInfinity       = NewNumber(math.Inf(+1))
	NegInfinity       = NewNumber(math.Inf(-1))
	Infinity          = PosInfinity
	PosZero           = NewNumber(+0)
	NegZero           = NewNumber(-0)
	Zero              = PosZero
)

type Number struct {
	value *big.Float
}

func NewNumber(x float64) Number {
	return Number{
		value: big.NewFloat(x),
	}
}

func (n Number) Value() interface{} {
	val, _ := n.value.Float64()
	return val
}

func (Number) Type() Type { return TypeNumber }
