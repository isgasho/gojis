package realm

import (
	"gitlab.com/gojis/vm/internal/runtime/binding"
	"gitlab.com/gojis/vm/internal/runtime/lang"
)

const (
	IntrinsicNameObjectPrototype   = "ObjectPrototype"
	IntrinsicNameFunctionPrototype = "FunctionPrototype"
	IntrinsicNameThrowTypeError    = "ThrowTypeError"
)

type Realm struct {
	intrinsics  *lang.Record
	globalObj   lang.Value                   // Object or Undefined
	globalEnv   lang.Value                   // Object or Undefined
	templateMap map[interface{}]*lang.Object // Parse Node -> Object
	HostDefined lang.Value
}

func CreateRealm() *Realm {
	r := new(Realm)
	CreateIntrinsics(r)
	r.globalObj = lang.Undefined
	r.globalEnv = lang.Undefined
	r.templateMap = make(map[interface{}]*lang.Object)
	return r
}

func CreateIntrinsics(r *Realm) {
	r.intrinsics = lang.NewRecord()
	objProto := lang.ObjectCreate(lang.Null)
	r.intrinsics.SetField(IntrinsicNameObjectPrototype, objProto)
	// FIXME: %ThrowTypeError% as in 8.2.2 and 9.2.9.1

	panic("TODO: 8.2.2")
}

func (r *Realm) GetIntrinsicObject(n string) lang.Value {
	val, ok := r.intrinsics.GetField(n)
	if !ok {
		return lang.Undefined
	}
	return val.(lang.Value)
}

func (r *Realm) SetRealmGlobalObject(globalObj, thisValue lang.Value) *Realm {
	if globalObj == lang.Undefined {
		panic("TODO: 8.2.3")
	}

	if thisValue == lang.Undefined {
		thisValue = globalObj
	}

	r.globalObj = globalObj

	globalEnv := binding.NewGlobalEnvironment(globalObj.(*lang.Object), thisValue.(*lang.Object))
	r.globalEnv = globalEnv

	return r
}

func (r *Realm) SetDefaultGlobalBindings() lang.Value {
	panic("TODO")
}
