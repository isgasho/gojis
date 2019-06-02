package binding

import (
	"gitlab.com/gojis/vm/internal/runtime/errors"
	"gitlab.com/gojis/vm/internal/runtime/lang"
)

type BindingStatus string

const (
	BindingStatusLexical       BindingStatus = "lexical"
	BindingStatusInitialized   BindingStatus = "initialized"
	BindingStatusUninitialized BindingStatus = "uninitialized"
)

var _ Environment = (*FunctionEnvironment)(nil)

type FunctionEnvironment struct {
	*DeclarativeEnvironment

	ThisValue         lang.Value
	ThisBindingStatus BindingStatus
	FunctionObject    *lang.Object
	HomeObject        lang.Value // Object or Undefined
	NewTarget         lang.Value // Object or Undefined
}

func (e *FunctionEnvironment) BindThisValue(val lang.Value) (lang.Value, errors.Error) {
	if e.ThisBindingStatus == BindingStatusInitialized {
		return nil, errors.NewReferenceError("'This' has already been initialized")
	}

	e.ThisValue = val
	e.ThisBindingStatus = BindingStatusInitialized
	return val, nil
}

func (e *FunctionEnvironment) HasThisBinding() bool {
	return e.ThisBindingStatus != BindingStatusLexical
}

func (e *FunctionEnvironment) HasSuperBinding() bool {
	if e.ThisBindingStatus == BindingStatusLexical {
		return false
	}

	return e.HomeObject != lang.Undefined
}

func (e *FunctionEnvironment) GetThisBinding() (lang.Value, errors.Error) {
	if e.ThisBindingStatus == BindingStatusUninitialized {
		return nil, errors.NewReferenceError("'This' has not been initialized yet")
	}

	return e.ThisValue, nil
}

func (e *FunctionEnvironment) GetSuperBase() lang.Value {
	if e.HomeObject == lang.Undefined {
		return lang.Undefined
	}

	return e.HomeObject.(*lang.Object).GetPrototypeOf()
}
