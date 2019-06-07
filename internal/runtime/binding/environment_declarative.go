package binding

import (
	"fmt"

	"github.com/TimSatke/gojis/internal/runtime/errors"
	"github.com/TimSatke/gojis/internal/runtime/lang"
)

var _ Environment = (*DeclarativeEnvironment)(nil)

type DeclarativeEnvironment struct {
	outer Environment

	bindings map[string]*Binding
}

func (e *DeclarativeEnvironment) Outer() Environment {
	return e.outer
}

func (e *DeclarativeEnvironment) setBinding(n lang.String, b *Binding) {
	e.bindings[n.Value().(string)] = b
}

func (e *DeclarativeEnvironment) getBinding(n lang.String) (*Binding, bool) {
	val, ok := e.bindings[n.Value().(string)]
	return val, ok
}

func (e *DeclarativeEnvironment) deleteBinding(n lang.String) {
	delete(e.bindings, n.Value().(string))
}

func (e *DeclarativeEnvironment) mustGetBinding(n lang.String) *Binding {
	val, ok := e.getBinding(n)
	if !ok {
		panic(fmt.Errorf("Missing binding for name '%v'", n.Value()))
	}

	return val
}

func (e *DeclarativeEnvironment) HasBinding(n lang.String) bool {
	_, ok := e.getBinding(n)
	return ok
}

func (e *DeclarativeEnvironment) CreateMutableBinding(n lang.String, deletable bool) errors.Error {
	b := NewBinding(n)
	b.deletable = deletable

	e.setBinding(n, b)

	return nil
}

func (e *DeclarativeEnvironment) CreateImmutableBinding(n lang.String, strict bool) errors.Error {
	b := NewBinding(n)
	b.immutable = true
	b.strict = strict

	e.setBinding(n, b)

	return nil
}

func (e *DeclarativeEnvironment) InitializeBinding(n lang.String, val lang.Value) errors.Error {
	e.mustGetBinding(n).value = val

	return nil
}

func (e *DeclarativeEnvironment) SetMutableBinding(n lang.String, val lang.Value, strict bool) errors.Error {
	if !e.HasBinding(n) {
		if strict {
			return errors.NewReferenceError("Cannot set mutable binding in strict mode if binding does not exist")
		}

		e.CreateMutableBinding(n, true)
		e.InitializeBinding(n, val)
		return nil
	}

	b := e.mustGetBinding(n)
	if b.IsStrict() {
		strict = true
	}

	if !b.IsInitialized() {
		return errors.NewReferenceError(fmt.Sprintf("The binding for '%v' has not been initialized yet", n))
	} else if !b.IsImmutable() {
		b.Set(val)
	} else {
		if strict {
			return errors.NewTypeError(fmt.Sprintf("Attempting to change value of an immutable binding '%v'", n))
		}
	}

	return nil
}

func (e *DeclarativeEnvironment) GetThisBinding() (lang.Value, errors.Error) {
	panic("TODO")
}

func (e *DeclarativeEnvironment) GetBindingValue(n lang.String, strict bool) (lang.Value, errors.Error) {
	b := e.mustGetBinding(n)
	if !b.IsInitialized() {
		return nil, errors.NewReferenceError(fmt.Sprintf("Binding '%v' has not been initialized yet", n))
	}

	return b.Value(), nil
}

func (e *DeclarativeEnvironment) DeleteBinding(n lang.String) bool {
	b, ok := e.getBinding(n)
	if !ok {
		return true
	}

	if !b.IsDeletable() {
		return false
	}

	e.deleteBinding(n)
	return true
}

func (e *DeclarativeEnvironment) HasThisBinding() bool {
	return false
}

func (e *DeclarativeEnvironment) HasSuperBinding() bool {
	return false
}

func (e *DeclarativeEnvironment) WithBaseObject() lang.Value {
	return lang.Undefined
}

func (e *DeclarativeEnvironment) Type() lang.Type    { return lang.TypeInternal }
func (e *DeclarativeEnvironment) Value() interface{} { return e }
