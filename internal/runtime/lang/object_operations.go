package lang

import (
	"fmt"

	"github.com/TimSatke/gojis/internal/runtime/errors"
)

// Get implements the function Get as specified in 7.3.1 in the specification.
func Get(o *Object, p StringOrSymbol) (Value, errors.Error) {
	return o.Get(p, o)
}

// GetV implements the function GetV as specified in 7.3.2 in the specification.
func GetV(v Value, p StringOrSymbol) (Value, errors.Error) {
	o := ToObject(v)
	return o.Get(p, v)
}

// Set implements the function Set as specified in 7.3.3 in the specification.
func Set(o *Object, p StringOrSymbol, v Value, throw bool) (Boolean, errors.Error) {
	success, err := o.Set(p, v, o)
	if err != nil {
		return False, err
	}

	if !success.Value().(bool) && throw {
		return False, errors.NewTypeError(fmt.Sprintf("Cannot set '%v' of object", p.Value()))
	}

	return success, nil
}

// CreateDataProperty implements the function CreateDataProperty as specified in 7.3.4 in the specification.
func CreateDataProperty(o *Object, p StringOrSymbol, v Value) Boolean {
	desc := NewDataProperty(v, True, True, True)
	return o.DefineOwnProperty(p, desc)
}

// CreateMethodProperty implements the function CreateMethodProperty as specified in 7.3.5 in the specification.
func CreateMethodProperty(o *Object, p StringOrSymbol, v Value) Boolean {
	desc := NewDataProperty(v, True, False, True)
	return o.DefineOwnProperty(p, desc)
}

// CreateDataPropertyOrThrow implements the function CreateDataPropertyOrThrow as specified in 7.3.6 in the specification.
func CreateDataPropertyOrThrow(o *Object, p StringOrSymbol, v Value) (Boolean, errors.Error) {
	success := CreateDataProperty(o, p, v)
	if !success {
		return False, errors.NewTypeError(fmt.Sprintf("Unable to create data property '%v'", p.Value()))
	}
	return success, nil
}

// DefinePropertyOrThrow implements the function DefinePropertyOrThrow as specified in 7.3.7 in the specification.
func DefinePropertyOrThrow(o *Object, p StringOrSymbol, desc *Property) (Boolean, errors.Error) {
	success := o.DefineOwnProperty(p, desc)
	if !success {
		return False, errors.NewTypeError(fmt.Sprintf("Unable to define property '%v'", p.Value()))
	}

	return success, nil
}

// DeletePropertyOrThrow implements the function DeletePropertyOrThrow as specified in 7.3.8 in the specification.
func DeletePropertyOrThrow(o *Object, p StringOrSymbol) (Boolean, errors.Error) {
	success := o.Delete(p)
	if !success {
		return False, errors.NewTypeError(fmt.Sprintf("Unable to delete property '%v'", p.Value()))
	}

	return success, nil
}

// GetMethod implements the function GetMethod as specified in 7.3.9 in the specification.
func GetMethod(v Value, p StringOrSymbol) (Value, errors.Error) {
	f, err := GetV(v, p)
	if err != nil {
		return nil, err
	}

	if f == Undefined || f == Null {
		return Undefined, nil
	}

	if !InternalIsCallable(f) {
		return nil, errors.NewTypeError("Object is not callable")
	}

	return f, nil
}

// HasProperty implements the function HasProperty as specified in 7.3.10 in the specification.
func HasProperty(o *Object, p StringOrSymbol) Boolean {
	return o.HasProperty(p)
}

// HasOwnProperty implements the function HasOwnProperty as specified in 7.3.11 in the specification.
func HasOwnProperty(o *Object, p StringOrSymbol) Boolean {
	return o.GetOwnProperty(p) != nil
}

// Call implements the function Call as specified in 7.3.12 in the specification.
func Call(f *Object, thisValue Value, args ...Value) (Value, errors.Error) {
	if args == nil {
		args = []Value{}
	}

	if !InternalIsCallable(f) {
		return nil, errors.NewTypeError("Object is not callable")
	}

	return f.Call(thisValue, args...)
}

// Construct implements the function Construct as specified in 7.3.13 in the specification.
func Construct(f, newTarget *Object, args ...Value) (*Object, errors.Error) {
	if newTarget == nil {
		newTarget = f
	}

	if args == nil {
		args = []Value{}
	}

	return f.Construct(newTarget, args...)
}

// Integrity levels as specified in 7.3.14 in the specification.
const (
	IntegrityLevelSealed = "sealed"
	IntegrityLevelFrozen = "frozen"
)

// SetIntegrityLevel implements the function SetIntegrityLevel as specified in 7.3.14 in the specification.
func SetIntegrityLevel(o *Object, level string) (Boolean, errors.Error) {
	status := o.PreventExtensions()
	if !status {
		return False, nil
	}

	keys := o.OwnPropertyKeys()
	switch level {
	case IntegrityLevelSealed:
		return setIntegrityLevelSealed(o, keys)
	case IntegrityLevelFrozen:
		return setIntegrityLevelFrozen(o, keys)
	default:
		panic(fmt.Errorf("Unknown integrity level '%v'", level))
	}
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

// TestIntegrityLevel implements the function TestIntegrityLevel as specified in 7.3.15 in the specification.
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

// CreateArrayFromList implements the function CreateArrayFromList as specified in 7.3.16 in the specification.
func CreateArrayFromList() {
	panic("TODO: Arrays")
}

// CreateListFromArrayLike implements the function CreateListFromArrayLike as specified in 7.3.17 in the specification.
func CreateListFromArrayLike() {
	panic("TODO: Arrays")
}

// Invoke implements the function Invoke as specified in 7.3.18 in the specification.
func Invoke(v Value, p StringOrSymbol, args ...Value) (Value, errors.Error) {
	if args == nil {
		args = []Value{}
	}

	f, err := GetV(v, p)
	if err != nil {
		return nil, err
	}

	return Call(f.(*Object), v, args...)
}

// OrdinaryHasInstance implements the function OrdinaryHasInstance as specified in 7.3.19 in the specification.
func OrdinaryHasInstance() {
	panic("TODO")
}

// SpeciesConstructor implements the function SpeciesConstructor as specified in 7.3.20 in the specification.
func SpeciesConstructor() {
	panic("TODO")
}

// EnumerableOwnPropertyNames implements the function EnumerableOwnPropertyNames as specified in 7.3.21 in the specification.
func EnumerableOwnPropertyNames() {
	panic("TODO")
}

// CopyDataProperties implements the function CopyDataProperties as specified in 7.3.23 in the specification.
func CopyDataProperties() {
	panic("TODO")
}
