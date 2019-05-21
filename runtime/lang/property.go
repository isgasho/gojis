package lang

type Property struct {
	*Record
}

func NewPropertyBase() *Property {
	p := new(Property)
	p.Record = NewRecord()
	return p
}

func NewDataProperty() *Property {
	p := NewPropertyBase()
	return p
}

func NewAccessorProperty() *Property {
	p := NewPropertyBase()
	return p
}

func (p *Property) Value() Value {
	val, ok := p.GetField("Value")
	if !ok {
		return Undefined
	}
	return val.(Value)
}

func (p *Property) Writable() Boolean {
	val, ok := p.GetField("Writable")
	if !ok {
		return False
	}
	return val.(Boolean)
}

func (p *Property) Get() func() Value {
	val, ok := p.GetField("Get")
	if !ok {
		return nil
	}
	return val.(func() Value)
}

func (p *Property) Set() func(Value) Boolean {
	val, ok := p.GetField("Set")
	if !ok {
		return nil
	}
	return val.(func(Value) Boolean)
}

func (p *Property) Enumerable() Boolean {
	val, ok := p.GetField("Enumerable")
	if !ok {
		return False
	}
	return val.(Boolean)
}

func (p *Property) Configurable() Boolean {
	val, ok := p.GetField("Configurable")
	if !ok {
		return False
	}
	return val.(Boolean)
}
