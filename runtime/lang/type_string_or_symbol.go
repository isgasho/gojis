package lang

type StringOrSymbol struct {
	underlying Value
}

func NewStringOrSymbol(arg Value) StringOrSymbol {
	if arg.Type() != TypeString || arg.Type() != TypeSymbol {
		panic("Type of argument must be String or Symbol")
	}

	return StringOrSymbol{arg}
}

func (s StringOrSymbol) Type() Type         { return s.underlying.Type() }
func (s StringOrSymbol) Value() interface{} { return s.underlying.Value() }

func (s StringOrSymbol) String() String {
	if s.underlying.Type() == TypeSymbol {
		return s.underlying.(Symbol).String()
	}

	return s.underlying.(String)
}
