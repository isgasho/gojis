package binding

import "github.com/gojisvm/gojis/internal/runtime/lang"

func GetIdentifierReference(env lang.InternalValue, n lang.String, strict bool) *Reference {
	if env == lang.Null {
		return NewReference(lang.NewStringOrSymbol(n), lang.Undefined, strict)
	}

	lex := env.(Environment)
	if lex.HasBinding(n) {
		return NewReference(lang.NewStringOrSymbol(n), lex, strict)
	}

	return GetIdentifierReference(lex.Outer(), n, strict)
}

func NewDeclarativeEnvironment(outer Environment) *DeclarativeEnvironment {
	env := new(DeclarativeEnvironment)
	env.outer = outer
	env.bindings = make(map[string]*Binding)
	return env
}

func NewObjectEnvironment(outer Environment, obj *lang.Object) *ObjectEnvironment {
	env := new(ObjectEnvironment)
	env.outer = outer
	env.bindingObject = obj
	return env
}

func NewFunctionEnvironment(outer Environment, f interface{}) *FunctionEnvironment {
	panic("TODO: ECMAScript function objects")
}

func NewGlobalEnvironment(globalObj *lang.Object, thisValue *lang.Object) *GlobalEnvironment {
	env := new(GlobalEnvironment)
	objRec := NewObjectEnvironment(env, globalObj)
	dclRec := NewDeclarativeEnvironment(env)
	env.ObjectRecord = objRec
	env.GlobalThisValue = thisValue
	env.DeclarativeRecord = dclRec
	env.VarNames = make([]string, 0)
	return env
}

func NewModuleEnvironment(outer Environment) *ModuleEnvironment {
	env := new(ModuleEnvironment)
	env.outer = outer
	return env
}
