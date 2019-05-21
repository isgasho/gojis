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

func SetIntegrityLevel() {
	panic("TODO")
}

func TestIntegrityLevel() {
	panic("TODO")
}

func CreateArrayFromList() {
	panic("TODO")
}

func CreateListFromArrayLike() {
	panic("TODO")
}

func Invoke() {
	panic("TODO")
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
