package realm

import (
	"github.com/TimSatke/gojis/internal/runtime/errors"
	"github.com/TimSatke/gojis/internal/runtime/lang"
)

// CreateBuiltinFunction creates a callable object, whose Call internal method will be the passed function fn.
// CreateBuiltinFunction is specified in 9.3.3.
func CreateBuiltinFunction(fn func(lang.Value, ...lang.Value) (lang.Value, errors.Error), realm *Realm, proto lang.Value, internalSlotsList ...lang.StringOrSymbol) *lang.Object {
	if realm == nil {
		panic("TODO: get current realm record")
	}
	if proto == nil {
		proto = realm.GetIntrinsicObject(IntrinsicNameFunctionPrototype)
	}
	fobj := lang.ObjectCreate(proto, internalSlotsList...)
	fobj.Call = fn
	panic("TODO: 9.3.3")
}
