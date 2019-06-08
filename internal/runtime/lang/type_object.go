package lang

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/TimSatke/gojis/internal/runtime/errors"
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

/* -- 9.1, ordinary object internal methods and internal slots -- */

func (o *Object) GetPrototypeOf() Value {
	return o.OrdinaryGetPrototypeOf()
}

func (o *Object) OrdinaryGetPrototypeOf() Value {
	return o.prototype
}

func (o *Object) SetPrototypeOf(v Value) Boolean {
	return o.OrdinarySetPrototypeOf(v)
}

func (o *Object) OrdinarySetPrototypeOf(v Value) Boolean {
	if t := v.Type(); t != TypeObject && t != TypeNull {
		panic(fmt.Sprintf("Type of prototype object must be Object or Null, cannot be %v", t))
	}

	extensible := o.extensible
	current := o.prototype

	if SameValue(v, current) {
		return True
	}

	if !extensible {
		return False
	}

	/*
	   This loop guarantees that there will be no circularities in any prototype chain that only
	   includes objects that use the ordinary object definitions for [[GetPrototypeOf]] and
	   [[SetPrototypeOf]].
	*/
	p := v
	done := false
	for !done {
		if p == Null {
			done = true
		} else if SameValue(p, o) {
			return False
		} else {
			// FIXME: if p.GetPrototypeOf is not the ordinary object internal method defined in 9.1.1,
			// set done to true.
			// else {
			p = p.(*Object).prototype // this type assertion cannot fail, since p is checked to be Object or Null, and Null is handled above
			// }
		}
	}

	o.prototype = v
	return True
}

func (o *Object) IsExtensible() Boolean {
	return o.OrdinaryIsExtensible()
}

func (o *Object) OrdinaryIsExtensible() Boolean {
	return Boolean(o.extensible)
}

func (o *Object) PreventExtensions() Boolean {
	return o.OrdinaryPreventExtensions()
}

func (o *Object) OrdinaryPreventExtensions() Boolean {
	o.extensible = false
	return True
}

func (o *Object) GetOwnProperty(p StringOrSymbol) *Property {
	return o.OrdinaryGetOwnProperty(p)
}

func (o *Object) OrdinaryGetOwnProperty(p StringOrSymbol) *Property {
	x, ok := o.fields[p]
	if !ok {
		return nil // actually Undefined
	}

	// TODO: maybe introduce a Propery#copy function?
	d := NewProperty()
	if x.IsDataDescriptor() {
		d.SetField(FieldNameValue, x.Value())
		d.SetField(FieldNameWritable, x.Writable())
	} else if x.IsAccessorDescriptor() {
		d.SetField(FieldNameGet, x.Get())
		d.SetField(FieldNameSet, x.Set())
	}
	d.SetField(FieldNameEnumerable, x.Enumerable())
	d.SetField(FieldNameConfigurable, x.Configurable())

	return d
}

func (o *Object) DefineOwnProperty(p StringOrSymbol, desc *Property) Boolean {
	return o.OrdinaryDefineOwnProperty(p, desc)
}

func (o *Object) OrdinaryDefineOwnProperty(p StringOrSymbol, desc *Property) Boolean {
	current := o.GetOwnProperty(p)
	extensible := o.extensible
	return o.ValidateAndApplyPropertyDescriptor(p, extensible, desc, current)
}

func (o *Object) IsCompatiblePropertyDescriptor(extensible bool, desc, current *Property) Boolean {
	// TODO: make code more beautiful
	return ((*Object)(nil)).ValidateAndApplyPropertyDescriptor(StringOrSymbol{}, extensible, desc, current)
}

func (o *Object) ValidateAndApplyPropertyDescriptor(p StringOrSymbol, extensible bool, desc, current *Property) Boolean {
	// if o != nil, p is not zero value
	if current == nil {
		if !extensible {
			return False
		}

		if desc.IsGenericDescriptor() || desc.IsDataDescriptor() {
			if o != nil {
				/*
					If O is not undefined, create an own accessor property named P of object O whose [[Get]], [[Set]],
					[[Enumerable]] and [[Configurable]] attribute values are described by Desc. If the value of an
					attribute field of Desc is absent, the attribute of the newly created property is set to its default
					value.
				*/
				panic("TODO")
			}
		} else {
			// desc.IsAccessorDescriptor() is true
			/*
			   If O is not undefined, create an own accessor property named P of object O whose [[Get]], [[Set]],
			   [[Enumerable]] and [[Configurable]] attribute values are described by Desc. If the value of an
			   attribute field of Desc is absent, the attribute of the newly created property is set to its default
			   value.
			*/
			panic("TODO")
		}
	}

	if len(desc.Record.fields) == 0 {
		return True
	}

	if !current.Configurable() {
		if desc.Configurable() {
			return False
		}

		if desc.Enumerable() && (current.Enumerable() == !desc.Enumerable()) {
			return False
		}
	}

	if desc.IsGenericDescriptor() {
		// no further valudation is required
	} else if current.IsDataDescriptor() != desc.IsDataDescriptor() {
		if !current.Configurable() {
			return False
		}

		if current.IsDataDescriptor() {
			/*
			   If O is not undefined, convert the property named P of object O from a data property to an
			   accessor property. Preserve the existing values of the converted property's [[Configurable]] and
			   [[Enumerable]] attributes and set the rest of the property's attributes to their default values.
			*/
			panic("TODO")
		} else {
			/*
			   If O is not undefined, convert the property named P of object O from an accessor property to a
			   data property. Preserve the existing values of the converted property's [[Configurable]] and
			   [[Enumerable]] attributes and set the rest of the property's attributes to their default values.
			*/
			panic("TODO")
		}
	} else if current.IsDataDescriptor() && desc.IsDataDescriptor() {
		if !current.Configurable() && !current.Writable() {
			if desc.Writable() {
				return False
			}

			if _, ok := desc.GetField(FieldNameValue); Boolean(ok) && !SameValue(desc.Value(), current.Value()) {
				return False
			}

			return True
		}
	} else if current.IsAccessorDescriptor() && desc.IsAccessorDescriptor() {
		if !current.Configurable() {
			if _, ok := desc.GetField(FieldNameSet); ok && reflect.ValueOf(desc.Set()).Pointer() != reflect.ValueOf(current.Set()).Pointer() {
				return False
			}

			return True
		}
	}

	if o != nil {
		prop := o.fields[p]
		for k, v := range desc.Record.fields {
			prop.fields[k] = v
		}
	}

	return True
}

func (o *Object) HasProperty(p StringOrSymbol) Boolean {
	return o.OrdinaryHasProperty(p)
}

func (o *Object) OrdinaryHasProperty(p StringOrSymbol) Boolean {
	if o.GetOwnProperty(p) != nil {
		return True
	}

	parent := o.GetPrototypeOf()
	if parent != Null {
		return parent.(*Object).HasProperty(p)
	}

	return False
}

func (o *Object) Get(p StringOrSymbol, receiver Value) (Value, errors.Error) {
	return o.OrdinaryGet(p, receiver)
}

func (o *Object) OrdinaryGet(p StringOrSymbol, receiver Value) (Value, errors.Error) {
	desc := o.GetOwnProperty(p)
	if desc == nil {
		parent := o.GetPrototypeOf()
		if parent == Null {
			return Undefined, nil
		}
		return parent.(*Object).Get(p, receiver)
	}

	if desc.IsDataDescriptor() {
		return desc.Value(), nil
	}

	if getter := desc.Get(); getter != Undefined {
		return Call(getter.(*Object), receiver)
	}
	return Undefined, nil
}

func (o *Object) Set(p StringOrSymbol, v, receiver Value) (Boolean, errors.Error) {
	return o.OrdinarySet(p, v, receiver)
}

func (o *Object) OrdinarySet(p StringOrSymbol, v, receiver Value) (Boolean, errors.Error) {
	return o.OrdinarySetWithOwnDescriptor(p, v, receiver, o.GetOwnProperty(p))
}

func (o *Object) OrdinarySetWithOwnDescriptor(p StringOrSymbol, v, receiver Value, ownDesc *Property) (Boolean, errors.Error) {
	if ownDesc == nil {
		parent := o.GetPrototypeOf()
		if parent != Null {
			return parent.(*Object).Set(p, v, receiver)
		} else {
			ownDesc = NewDataProperty(Undefined, True, True, True)
		}
	}

	if ownDesc.IsDataDescriptor() {
		if !ownDesc.Writable() {
			return False, nil
		}

		if receiver.Type() != TypeObject {
			return False, nil
		}

		receiverObj := receiver.(*Object)
		existingDescriptor := receiverObj.GetOwnProperty(p)
		if existingDescriptor != nil {
			if existingDescriptor.IsAccessorDescriptor() {
				return False, nil
			}

			if !existingDescriptor.Writable() {
				return False, nil
			}

			valueDesc := NewProperty()
			valueDesc.SetField(FieldNameValue, v)
			return receiverObj.DefineOwnProperty(p, valueDesc), nil
		} else {
			return CreateDataProperty(receiverObj, p, v), nil
		}
	}

	// assert: ownDesc.IsAccessorDescriptor is true

	if setter := ownDesc.Set(); setter != Undefined {
		_, err := Call(setter.(*Object), receiver, v)
		if err != nil {
			return False, err
		}
		return True, nil
	}
	return False, nil
}

func (o *Object) Delete(p StringOrSymbol) Boolean {
	return o.OrdinaryDelete(p)
}

func (o *Object) OrdinaryDelete(p StringOrSymbol) Boolean {
	desc := o.GetOwnProperty(p)
	if desc == nil {
		return True
	}

	if desc.Configurable() {
		delete(o.fields, p)
		return True
	}

	return False
}

func (o *Object) OwnPropertyKeys() []StringOrSymbol {
	return o.OrdinaryOwnPropertyKeys()
}

func (o *Object) OrdinaryOwnPropertyKeys() []StringOrSymbol {
	keys := []StringOrSymbol{}
	secs := make(map[int][]StringOrSymbol)
	const (
		integerIndex int = iota
		stringIndex
		symbolIndex
	)

	for k := range o.fields {
		if n, err := strconv.ParseInt(k.String().Value().(string), 10, 64); err == nil && n >= 0 {
			// k is an integer index
			secs[integerIndex] = append(secs[integerIndex], k)
			continue
		}

		if k.Type() == TypeString {
			secs[stringIndex] = append(secs[stringIndex], k)
		}

		if k.Type() == TypeSymbol {
			secs[symbolIndex] = append(secs[symbolIndex], k)
		}
	}

	keys = append(keys, secs[integerIndex]...)
	keys = append(keys, secs[stringIndex]...)
	keys = append(keys, secs[symbolIndex]...)
	return keys
}
