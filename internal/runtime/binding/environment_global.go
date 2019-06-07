package binding

import (
	"fmt"

	"github.com/TimSatke/gojis/internal/runtime/errors"
	"github.com/TimSatke/gojis/internal/runtime/lang"
)

var _ Environment = (*GlobalEnvironment)(nil)

type GlobalEnvironment struct {
	ObjectRecord      *ObjectEnvironment
	GlobalThisValue   *lang.Object
	DeclarativeRecord *DeclarativeEnvironment
	VarNames          []string
}

func (e *GlobalEnvironment) Outer() Environment {
	return nil
}

func (e *GlobalEnvironment) GetThisBinding() (lang.Value, errors.Error) {
	return e.GlobalThisValue, nil
}

func (e *GlobalEnvironment) HasVarDeclaration(n lang.String) bool {
	nVal := n.Value().(string)
	for _, varName := range e.VarNames {
		if varName == nVal {
			return true
		}
	}
	return false
}

func (e *GlobalEnvironment) HasLexicalDeclaration(n lang.String) bool {
	return e.DeclarativeRecord.HasBinding(n)
}

func (e *GlobalEnvironment) HasRestrictedGlobalProperty(n lang.String) bool {
	existingProp := e.ObjectRecord.bindingObject.GetOwnProperty(lang.NewStringOrSymbol(n)).Value()
	if existingProp == lang.Undefined {
		return false
	}

	panic("TODO: properties")
}

func (e *GlobalEnvironment) CanDeclareGlobalVar(n lang.String) bool {
	globalObj := e.ObjectRecord.bindingObject

	if lang.HasOwnProperty(globalObj, lang.NewStringOrSymbol(n)) {
		return true
	}

	return lang.InternalIsExtensible(globalObj)
}

func (e *GlobalEnvironment) CanDeclareGlobalFunction(n lang.String) {
	panic("TODO: properties")
}

func (e *GlobalEnvironment) CreateGlobalVarBinding(n lang.String, deletable bool) {
	globalObj := e.ObjectRecord.bindingObject

	hasProperty := lang.HasOwnProperty(globalObj, lang.NewStringOrSymbol(n))
	extensible := lang.InternalIsExtensible(globalObj)
	if !hasProperty.Value().(bool) && extensible {
		e.ObjectRecord.CreateMutableBinding(n, deletable)
		e.ObjectRecord.InitializeBinding(n, lang.Undefined)
	}

	if !e.HasVarDeclaration(n) {
		e.VarNames = append(e.VarNames, n.Value().(string))
	}
}

func (e *GlobalEnvironment) CreateGlobalFunctionBinding(n lang.String, val lang.Value, deletable bool) {
	panic("TODO: properties")
}

/* -- implements Environment -- */

func (e *GlobalEnvironment) HasBinding(n lang.String) bool {
	return e.DeclarativeRecord.HasBinding(n) || e.ObjectRecord.HasBinding(n)
}

func (e *GlobalEnvironment) CreateMutableBinding(n lang.String, deletable bool) errors.Error {
	if e.DeclarativeRecord.HasBinding(n) {
		return errors.NewTypeError(fmt.Sprintf("Declarative environment record already has a binding for '%v'", n))
	}

	return e.DeclarativeRecord.CreateMutableBinding(n, deletable)
}

func (e *GlobalEnvironment) CreateImmutableBinding(n lang.String, strict bool) errors.Error {
	if e.DeclarativeRecord.HasBinding(n) {
		return errors.NewTypeError(fmt.Sprintf("Declarative environment record already has a binding for '%v'", n))
	}

	return e.DeclarativeRecord.CreateImmutableBinding(n, strict)
}

func (e *GlobalEnvironment) InitializeBinding(n lang.String, val lang.Value) errors.Error {
	if e.DeclarativeRecord.HasBinding(n) {
		return e.DeclarativeRecord.InitializeBinding(n, val)
	}

	return e.ObjectRecord.InitializeBinding(n, val)
}

func (e *GlobalEnvironment) SetMutableBinding(n lang.String, val lang.Value, strict bool) errors.Error {
	if e.DeclarativeRecord.HasBinding(n) {
		return e.DeclarativeRecord.SetMutableBinding(n, val, strict)
	}

	return e.ObjectRecord.SetMutableBinding(n, val, strict)
}

func (e *GlobalEnvironment) GetBindingValue(n lang.String, strict bool) (lang.Value, errors.Error) {
	if e.DeclarativeRecord.HasBinding(n) {
		return e.DeclarativeRecord.GetBindingValue(n, strict)
	}

	return e.ObjectRecord.GetBindingValue(n, strict)
}

func (e *GlobalEnvironment) DeleteBinding(n lang.String) bool {
	if e.DeclarativeRecord.HasBinding(n) {
		return e.DeclarativeRecord.DeleteBinding(n)
	}

	if lang.HasOwnProperty(e.ObjectRecord.bindingObject, lang.NewStringOrSymbol(n)) {
		status := e.ObjectRecord.DeleteBinding(n)
		if status {
			nVal := n.Value().(string)
			for i, varName := range e.VarNames {
				if varName == nVal {
					e.VarNames[i] = e.VarNames[len(e.VarNames)-1]
					e.VarNames = e.VarNames[:len(e.VarNames)-1]
					break
				}
			}
		}

		return status
	}

	return true
}

func (e *GlobalEnvironment) HasThisBinding() bool {
	return true
}

func (e *GlobalEnvironment) HasSuperBinding() bool {
	return false
}

func (e *GlobalEnvironment) WithBaseObject() lang.Value {
	return lang.Undefined
}

func (e *GlobalEnvironment) Type() lang.Type    { return lang.TypeInternal }
func (e *GlobalEnvironment) Value() interface{} { return e }
