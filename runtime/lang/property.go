package lang

const (
	FieldNameValue        = "Value"
	FieldNameWritable     = "Writable"
	FieldNameGet          = "Get"
	FieldNameSet          = "Set"
	FieldNameEnumerable   = "Enumerable"
	FieldNameConfigurable = "Configurable"
)

type Property struct {
	*Record
}

func NewPropertyBase(enumerable, configurable Boolean) *Property {
	p := new(Property)
	p.Record = NewRecord()
	p.SetField(FieldNameEnumerable, enumerable)
	p.SetField(FieldNameConfigurable, configurable)
	return p
}

func NewDataProperty(value Value, writable, enumerable, configurable Boolean) *Property {
	p := NewPropertyBase(enumerable, configurable)
	p.SetField(FieldNameValue, value)
	p.SetField(FieldNameWritable, writable)
	return p
}

func NewAccessorProperty(enumerable, configurable Boolean) *Property {
	p := NewPropertyBase(enumerable, configurable)
	return p
}

func (p *Property) Value() Value {
	val, ok := p.GetField(FieldNameValue)
	if !ok {
		return Undefined
	}
	return val.(Value)
}

func (p *Property) Writable() Boolean {
	val, ok := p.GetField(FieldNameWritable)
	if !ok {
		return False
	}
	return val.(Boolean)
}

func (p *Property) Get() func() Value {
	val, ok := p.GetField(FieldNameGet)
	if !ok {
		return nil
	}
	return val.(func() Value)
}

func (p *Property) Set() func(Value) Boolean {
	val, ok := p.GetField(FieldNameSet)
	if !ok {
		return nil
	}
	return val.(func(Value) Boolean)
}

func (p *Property) Enumerable() Boolean {
	val, ok := p.GetField(FieldNameEnumerable)
	if !ok {
		return False
	}
	return val.(Boolean)
}

func (p *Property) Configurable() Boolean {
	val, ok := p.GetField(FieldNameConfigurable)
	if !ok {
		return False
	}
	return val.(Boolean)
}
