package lang

import (
	"fmt"

	"gitlab.com/gojis/vm/internal/runtime/errors"
)

func Get(o *Object, p StringOrSymbol) Value {
	return o.Get(p, o)
}

func GetV(v Value, p StringOrSymbol) Value {
	o := ToObject(v)
	return o.Get(p, v)
}

func Set(o *Object, p StringOrSymbol, v Value, throw bool) (Boolean, errors.Error) {
	success := o.Set(p, v, o)
	if !success.Value().(bool) && throw {
		return False, errors.NewTypeError(fmt.Sprintf("Cannot set '%v' of object", p.Value()))
	}

	return success, nil
}

func CreateDataProperty(o *Object, p StringOrSymbol, v Value) Boolean {
	desc := NewDataProperty(v, True, True, True)
	return o.DefineOwnProperty(p, desc)
}

func CreateMethodProperty(o *Object, p StringOrSymbol, v Value) Boolean {
	desc := NewDataProperty(v, True, False, True)
	return o.DefineOwnProperty(p, desc)
}

func CreateDataPropertyOrThrow(o *Object, p StringOrSymbol, v Value) (Boolean, errors.Error) {
	success := CreateDataProperty(o, p, v)
	if !success {
		return False, errors.NewTypeError(fmt.Sprintf("Unable to create data property '%v'", p.Value()))
	}
	return success, nil
}

func DefinePropertyOrThrow(o *Object, p StringOrSymbol, desc *Property) (Boolean, errors.Error) {
	success := o.DefineOwnProperty(p, desc)
	if !success {
		return False, errors.NewTypeError(fmt.Sprintf("Unable to define property '%v'", p.Value()))
	}

	return success, nil
}

func DeletePropertyOrThrow(o *Object, p StringOrSymbol) (Boolean, errors.Error) {
	success := o.Delete(p)
	if !success {
		return False, errors.NewTypeError(fmt.Sprintf("Unable to delete property '%v'", p.Value()))
	}

	return success, nil
}

func GetMethod(v Value, p StringOrSymbol) (Value, errors.Error) {
	f := GetV(v, p)
	if f == Undefined || f == Null {
		return Undefined, nil
	}

	if !InternalIsCallable(f) {
		return nil, errors.NewTypeError("Object is not callable")
	}

	return f, nil
}

func HasProperty(o *Object, p StringOrSymbol) Boolean {
	return o.HasProperty(p)
}

func HasOwnProperty(o *Object, p StringOrSymbol) Boolean {
	return o.HasOwnProperty(p)
}

func Call(f *Object, thisValue Value, args ...Value) (Value, errors.Error) {
	if args == nil {
		args = []Value{}
	}

	if !InternalIsCallable(f) {
		return nil, errors.NewTypeError("Object is not callable")
	}

	return f.Call(thisValue, args...)
}

func Construct(f, newTarget *Object, args ...Value) (*Object, errors.Error) {
	if newTarget == nil {
		newTarget = f
	}

	if args == nil {
		args = []Value{}
	}

	return f.Construct(newTarget, args...)
}

const (
	IntegrityLevelSealed = "sealed"
	IntegrityLevelFrozen = "frozen"
)

func SetIntegrityLevel(o *Object, level string) (Boolean, errors.Error) {
	status := o.PreventExtensions()
	if !status {
		return False, nil
	}

	keys := o.OwnPropertyKeys()
	if level == IntegrityLevelSealed {
		return setIntegrityLevelSealed(o, keys)
	}
	if level == IntegrityLevelFrozen {
		return setIntegrityLevelFrozen(o, keys)
	}

	panic(fmt.Errorf("Unknown integrity level '%v'", level))
}

func setIntegrityLevelSealed(o *Object, keys []StringOrSymbol) (Boolean, errors.Error) {
	for _, k := range keys {
		p := NewProperty()
		p.SetField(FieldNameConfigurable, False)
		_, err := DefinePropertyOrThrow(o, k, p)
		if err != nil {
			return False, err
		}
	}
	return True, nil
}

func setIntegrityLevelFrozen(o *Object, keys []StringOrSymbol) (Boolean, errors.Error) {
	for _, k := range keys {
		p := o.GetOwnProperty(k)
		if p != nil {
			desc := NewProperty()
			if p.IsAccessorDescriptor() {
				desc.SetField(FieldNameConfigurable, False)
			} else {
				desc.SetField(FieldNameConfigurable, False)
				desc.SetField(FieldNameWritable, False)
			}
			_, err := DefinePropertyOrThrow(o, k, desc)
			if err != nil {
				return False, err
			}
		}
	}
	return True, nil
}

func TestIntegrityLevel(o *Object, level string) Boolean {
	if o.IsExtensible() {
		return False
	}

	keys := o.OwnPropertyKeys()
	for _, k := range keys {
		desc := o.GetOwnProperty(k)
		if desc != nil {
			if desc.Configurable() {
				return False
			}
			if level == IntegrityLevelFrozen && desc.IsDataDescriptor() {
				if desc.Writable() {
					return false
				}
			}
		}
	}
	return true
}

func CreateArrayFromList() {
	panic("TODO: Arrays")
}

func CreateListFromArrayLike() {
	panic("TODO: Arrays")
}

func Invoke(v Value, p StringOrSymbol, args ...Value) (Value, errors.Error) {
	if args == nil {
		args = []Value{}
	}

	f := GetV(v, p)
	return Call(f.(*Object), v, args...)
}

func OrdinaryHasInstance() {
	panic("TODO")
}

func SpeciesConstructor() {
	panic("TODO")
}

func EnumerableOwnPropertyNames() {
	panic("TODO")
}

func GetFunctionRealm() {
	panic("TODO")
}

func CopyDataProperties() {
	panic("TODO")
}
