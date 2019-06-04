package binding

import (
	"gitlab.com/gojis/vm/internal/runtime/errors"
	"gitlab.com/gojis/vm/internal/runtime/lang"
)

var _ Environment = (*ObjectEnvironment)(nil)

type ObjectEnvironment struct {
	outer Environment

	bindingObject *lang.Object
}

func (e *ObjectEnvironment) Outer() Environment {
	return e.outer
}

func (e *ObjectEnvironment) IsGlobalEnvironment() bool {
	panic("TODO")
}

func (e *ObjectEnvironment) IsModuleEnvironment() bool {
	panic("TODO")
}

func (e *ObjectEnvironment) HasBinding(n lang.String) bool {
	panic("TODO")
}

func (e *ObjectEnvironment) CreateMutableBinding(n lang.String, deletable bool) errors.Error {
	panic("TODO")
}

func (e *ObjectEnvironment) CreateImmutableBinding(n lang.String, strict bool) errors.Error {
	panic("TODO")
}

func (e *ObjectEnvironment) InitializeBinding(n lang.String, val lang.Value) errors.Error {
	panic("TODO")
}

func (e *ObjectEnvironment) SetMutableBinding(n lang.String, val lang.Value, strict bool) errors.Error {
	panic("TODO")
}

func (e *ObjectEnvironment) GetBindingValue(n lang.String, strict bool) (lang.Value, errors.Error) {
	panic("TODO")
}

func (e *ObjectEnvironment) DeleteBinding(n lang.String) bool {
	panic("TODO")
}

func (e *ObjectEnvironment) HasThisBinding() bool {
	panic("TODO")
}

func (e *ObjectEnvironment) HasSuperBinding() bool {
	panic("TODO")
}

func (e *ObjectEnvironment) WithBaseObject() lang.Value {
	panic("TODO")
}

func (e *ObjectEnvironment) Type() lang.Type    { return lang.TypeInternal }
func (e *ObjectEnvironment) Value() interface{} { return e }
