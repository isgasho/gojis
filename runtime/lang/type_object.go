package lang

import (
	"fmt"

	"gitlab.com/gojis/vm/runtime/errors"
)

var _ Value = (*Object)(nil) // ensure that Object implements Value

// Object is a language type as specified by the language spec.
type Object struct {
	fields map[StringOrSymbol]*Property
	slots  map[StringOrSymbol]Value

	prototype  Value // *Object or Null
	extensible bool

	// Function Object

	// Call is the Call function of an object as (kind of)
	// specified by the language spec.
	// This is only not nil for function and constructor function
	// objects.
	Call func(Value, ...Value) (Value, errors.Error)

	// Constructor Function Object

	// Construct is the Construct function of an obejct as (kind of)
	// specified by the language spec.
	// This is only not nil for constructor function objects.
	Construct func(*Object, ...Value) (*Object, errors.Error)
}

func ObjectCreate(proto Value, internalSlotsList ...StringOrSymbol) *Object {
	if internalSlotsList == nil {
		internalSlotsList = []StringOrSymbol{}
	}

	obj := new(Object)
	obj.fields = make(map[StringOrSymbol]*Property)
	obj.slots = make(map[StringOrSymbol]Value)
	for _, slot := range internalSlotsList {
		obj.slots[slot] = Undefined
	}
	EnsureTypeOneOf(proto, TypeObject, TypeNull) // panic if proto is not TypeObject or TypeNull
	obj.prototype = proto
	obj.extensible = true

	return obj
}

// Value returns the object itself.
//
//	var o *Object
//	...
//	o == o.Value() // true
func (o *Object) Value() interface{} { return o }

// Type returns lang.TypeObject.
func (o *Object) Type() Type { return TypeObject }

/* -- 6.1.7.2, Table 5, essential internal methods -- */

// GetPrototypeOf as specified by the language spec.
func (o *Object) GetPrototypeOf() Value {
	if o.prototype == nil {
		// should not happen
		return Null
	}

	return o.prototype
}

// SetPrototypeOf as specified by the language spec.
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

// IsExtensible as specified by the language spec.
func (o *Object) IsExtensible() Boolean {
	return Boolean(o.extensible)
}

// PreventExtensions as specified by the language spec.
func (o *Object) PreventExtensions() Boolean {
	o.extensible = false
	return True
}

// GetOwnProperty as specified by the language spec.
func (o *Object) GetOwnProperty(key StringOrSymbol) *Property {
	return o.fields[key]
}

// DefineOwnProperty as specified by the language spec.
func (o *Object) DefineOwnProperty(key StringOrSymbol, val *Property) Boolean {
	o.fields[key] = val
	return True
}

// HasProperty as specified by the language spec.
func (o *Object) HasProperty(key StringOrSymbol) Boolean {
	_, ok := o.fields[key]
	return Boolean(ok)
}

// HasOwnProperty as specified by the language spec.
func (o *Object) HasOwnProperty(key StringOrSymbol) Boolean {
	panic("TODO")
}

// Get as specified by the language spec.
func (o *Object) Get(key StringOrSymbol, receiver Value) Value {
	if p, ok := o.fields[key]; ok {
		return p.Value()
	}
	return Undefined
}

// Set as specified by the language spec.
func (o *Object) Set(key StringOrSymbol, val, receiver Value) Boolean {
	if p, ok := o.fields[key]; ok {
		return p.Set()(val)
	}
	return False
}

// Delete as specified by the language spec.
func (o *Object) Delete(key StringOrSymbol) Boolean {
	// FIXME: deletable?
	delete(o.fields, key)
	return True
}

// OwnPropertyKeys as specified by the language spec.
func (o *Object) OwnPropertyKeys() (res []StringOrSymbol) {
	res = make([]StringOrSymbol, len(o.fields))
	i := 0
	for k := range o.fields {
		res[i] = k
	}
	return
}
