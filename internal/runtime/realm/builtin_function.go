package realm

import (
	"gitlab.com/gojis/vm/internal/runtime/errors"
	"gitlab.com/gojis/vm/internal/runtime/lang"
)

func CreateBuiltinFunction(fn func(lang.Value, ...lang.Value) (lang.Value, errors.Error), realm *Realm, proto lang.Value, internalSlotsList ...lang.StringOrSymbol) {
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
