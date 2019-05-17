package vm

type Object struct {
	properties map[Value]*Property // Symbol/String -> *Property
	slots      map[Value]Value     // String (?) -> Value (Undefined by default)
}

func (o *Object) GetPrototypeOf() Value /* Object or Null */ {
	panic("TODO")
}

func (o *Object) SetPrototypeOf(v Value /* Object or Null */) Boolean {
	panic("TODO")
}

func (o *Object) IsExtensible() Boolean {
	panic("TODO")
}

func (o *Object) PreventExtensions() Boolean {
	panic("TODO")
}

func (o *Object) GetOwnProperty(v Value /* Symbol or String */) Value /* Undefined or PropertyDescriptor */ {
	panic("TODO")
}
