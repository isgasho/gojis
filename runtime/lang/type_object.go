package lang

import (
	"fmt"

	"gitlab.com/gojis/vm/runtime/errors"
)

var _ Value = (*Object)(nil)

type Object struct {
	fields map[StringOrSymbol]*Property
	slots  map[StringOrSymbol]Value

	prototype  *Object
	extensible bool

	// Function Object

	Call func(Value, ...Value) (Value, errors.Error)

	// Constructor Function Object

	Construct func(*Object, ...Value) (*Object, errors.Error)
}

func (o *Object) Type() Type { return TypeObject }

func (o *Object) Value() interface{} { return o }

/* -- 6.1.7.2, Table 5, essential internal methods -- */

func (o *Object) GetPrototypeOf() Value {
	if o.prototype == nil {
		return Null
	}

	return o.prototype
}

func (o *Object) SetPrototypeOf(val Value) Boolean {
	if val == Null {
		o.prototype = nil
		return True
	}

	if t := val.Type(); t != TypeObject {
		panic(fmt.Sprintf("Invalid type %v for field prototype, need Object or Null", t))
	}

	o.prototype = val.(*Object)
	return True
}

func (o *Object) IsExtensible() Boolean {
	return Boolean(o.extensible)
}

func (o *Object) PreventExtensions() Boolean {
	o.extensible = false
	return True
}

func (o *Object) GetOwnProperty(key StringOrSymbol) *Property {
	return o.fields[key]
}

func (o *Object) DefineOwnProperty(key StringOrSymbol, val *Property) Boolean {
	o.fields[key] = val
	return True
}

func (o *Object) HasProperty(key StringOrSymbol) Boolean {
	_, ok := o.fields[key]
	return Boolean(ok)
}

func (o *Object) HasOwnProperty(key StringOrSymbol) Boolean {
	panic("TODO")
}

func (o *Object) Get(key StringOrSymbol, receiver Value) Value {
	if p, ok := o.fields[key]; ok {
		return p.Value()
	}
	return Undefined
}

func (o *Object) Set(key StringOrSymbol, val, receiver Value) Boolean {
	if p, ok := o.fields[key]; ok {
		return p.Set()(val)
	}
	return False
}

func (o *Object) Delete(key StringOrSymbol) Boolean {
	// FIXME: deletable?
	delete(o.fields, key)
	return True
}

func (o *Object) OwnPropertyKeys() (res []StringOrSymbol) {
	res = make([]StringOrSymbol, len(o.fields))
	i := 0
	for k := range o.fields {
		res[i] = k
	}
	return
}
