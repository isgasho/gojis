package lang

type Value interface {
	Type() Type
	Value() interface{}
}
