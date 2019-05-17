package vm

type Value interface {
	Type() Type
	Value() interface{}
}
