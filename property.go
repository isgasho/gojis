package vm

type Property struct {
	Value Value

	Get Value // either a function object or Undefined
	Set Value // either a function object or Undefined

	Writable     Boolean
	Enumberable  Boolean
	Configurable Boolean
}

func NewProperty() *Property {
	p := new(Property)
	// default values as specified in 6.1.7.1, Table 4
	p.Value = Undefined
	p.Get = Undefined
	p.Set = Undefined
	p.Writable = False
	p.Enumberable = False
	p.Configurable = False
	return p
}
