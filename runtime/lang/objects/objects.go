package objects

import (
	"fmt"

	"gitlab.com/gojis/vm/runtime/errors"
	"gitlab.com/gojis/vm/runtime/lang"
	"gitlab.com/gojis/vm/runtime/lang/cmp"
)

func Get(o *lang.Object, p lang.StringOrSymbol) lang.Value {
	return o.Get(p, o)
}

func GetV(v lang.Value, p lang.StringOrSymbol) lang.Value {
	o := lang.ToObject(v)
	return o.Get(p, v)
}

func Set(o *lang.Object, p lang.StringOrSymbol, v lang.Value, throw bool) (lang.Boolean, errors.Error) {
	success := o.Set(p, v, o)
	if !success.Value().(bool) && throw {
		return lang.False, errors.NewTypeError(fmt.Sprintf("Cannot set '%v' of object", p.Value()))
	}

	return success, nil
}

func CreateDataProperty(o *lang.Object, p lang.StringOrSymbol, v lang.Value) lang.Boolean {
	desc := lang.NewDataProperty(v, lang.True, lang.True, lang.True)
	return o.DefineOwnProperty(p, desc)
}

func CreateMethodProperty(o *lang.Object, p lang.StringOrSymbol, v lang.Value) lang.Boolean {
	desc := lang.NewDataProperty(v, lang.True, lang.False, lang.True)
	return o.DefineOwnProperty(p, desc)
}

func CreateDataPropertyOrThrow(o *lang.Object, p lang.StringOrSymbol, v lang.Value) (lang.Boolean, errors.Error) {
	success := CreateDataProperty(o, p, v)
	if !success {
		return lang.False, errors.NewTypeError(fmt.Sprintf("Unable to create data property '%v'", p.Value()))
	}
	return success, nil
}

func DefinePropertyOrThrow(o *lang.Object, p lang.StringOrSymbol, desc *lang.Property) (lang.Boolean, errors.Error) {
	success := o.DefineOwnProperty(p, desc)
	if !success {
		return lang.False, errors.NewTypeError(fmt.Sprintf("Unable to define property '%v'", p.Value()))
	}

	return success, nil
}

func DeletePropertyOrThrow(o *lang.Object, p lang.StringOrSymbol) (lang.Boolean, errors.Error) {
	success := o.Delete(p)
	if !success {
		return lang.False, errors.NewTypeError(fmt.Sprintf("Unable to delete property '%v'", p.Value()))
	}

	return success, nil
}

func GetMethod(v lang.Value, p lang.StringOrSymbol) (lang.Value, errors.Error) {
	f := GetV(v, p)
	if f == lang.Undefined || f == lang.Null {
		return lang.Undefined, nil
	}

	if !cmp.InternalIsCallable(f) {
		return nil, errors.NewTypeError("Object is not callable")
	}

	return f, nil
}

func HasProperty(o *lang.Object, p lang.StringOrSymbol) lang.Boolean {
	return o.HasProperty(p)
}

func HasOwnProperty(o *lang.Object, p lang.StringOrSymbol) lang.Boolean {
	return o.HasOwnProperty(p)
}

func Call(f *lang.Object, thisValue lang.Value, args ...lang.Value) (lang.Value, errors.Error) {
	if args == nil {
		args = []lang.Value{}
	}

	if !cmp.InternalIsCallable(f) {
		return nil, errors.NewTypeError("Object is not callable")
	}

	return f.Call(thisValue, args...)
}

func Construct(f, newTarget *lang.Object, args ...lang.Value) (*lang.Object, errors.Error) {
	if newTarget == nil {
		newTarget = f
	}

	if args == nil {
		args = []lang.Value{}
	}

	return f.Construct(newTarget, args...)
}

const (
	IntegrityLevelSealed = "sealed"
	IntegrityLevelFrozen = "frozen"
)

func SetIntegrityLevel(o *lang.Object, level string) (lang.Boolean, errors.Error) {
	status := o.PreventExtensions()
	if !status {
		return lang.False, nil
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

func setIntegrityLevelSealed(o *lang.Object, keys []lang.StringOrSymbol) (lang.Boolean, errors.Error) {
	for _, k := range keys {
		p := lang.NewProperty()
		p.SetField(lang.FieldNameConfigurable, lang.False)
		_, err := DefinePropertyOrThrow(o, k, p)
		if err != nil {
			return lang.False, err
		}
	}
	return lang.True, nil
}

func setIntegrityLevelFrozen(o *lang.Object, keys []lang.StringOrSymbol) (lang.Boolean, errors.Error) {
	for _, k := range keys {
		p := o.GetOwnProperty(k)
		if p != nil {
			desc := lang.NewProperty()
			if p.IsAccessorDescriptor() {
				desc.SetField(lang.FieldNameConfigurable, lang.False)
			} else {
				desc.SetField(lang.FieldNameConfigurable, lang.False)
				desc.SetField(lang.FieldNameWritable, lang.False)
			}
			_, err := DefinePropertyOrThrow(o, k, desc)
			if err != nil {
				return lang.False, err
			}
		}
	}
	return lang.True, nil
}

func TestIntegrityLevel(o *lang.Object, level string) lang.Boolean {
	if o.IsExtensible() {
		return lang.False
	}

	keys := o.OwnPropertyKeys()
	for _, k := range keys {
		desc := o.GetOwnProperty(k)
		if desc != nil {
			if desc.Configurable() {
				return lang.False
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

func Invoke(v lang.Value, p lang.StringOrSymbol, args ...lang.Value) (lang.Value, errors.Error) {
	if args == nil {
		args = []lang.Value{}
	}

	f := GetV(v, p)
	return Call(f.(*lang.Object), v, args...)
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
