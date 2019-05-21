package binding

import (
	"fmt"

	"gitlab.com/gojis/vm/runtime/errors"
	"gitlab.com/gojis/vm/runtime/lang"
)

type Reference struct {
	baseComponent  lang.Value
	referencedName lang.StringOrSymbol
	thisValue      lang.Value
	strict         bool
}

func NewReference(n lang.StringOrSymbol, base lang.Value, strict bool) *Reference {
	r := new(Reference)
	r.referencedName = n
	r.baseComponent = base
	r.strict = strict
	return r
}

func NewSuperReference(n lang.StringOrSymbol, base, this lang.Value, strict bool) *Reference {
	r := NewReference(n, base, strict)
	r.thisValue = this
	return r
}

func (r *Reference) GetBase() lang.Value {
	return r.baseComponent
}

func (r *Reference) GetReferencedName() lang.StringOrSymbol {
	return r.referencedName
}

func (r *Reference) IsStrictReference() bool {
	return r.strict
}

func (r *Reference) HasPrimitiveBase() bool {
	switch r.baseComponent.Type() {
	case lang.TypeBoolean, lang.TypeString, lang.TypeSymbol, lang.TypeNumber:
		return true
	default:
		return false
	}
}

func (r *Reference) IsPropertyReference() bool {
	return r.baseComponent.Type() == lang.TypeObject ||
		r.HasPrimitiveBase()
}

func (r *Reference) IsUnresolvableReference() bool {
	return r.baseComponent == lang.Undefined
}

func (r *Reference) IsSuperReference() bool {
	return r.thisValue != nil
}

func (r *Reference) GetValue() (lang.Value, errors.Error) {
	base := r.GetBase()

	if r.IsUnresolvableReference() {
		return nil, errors.NewReferenceError(fmt.Sprintf("Unresolvable reference: '%v'", r.referencedName))
	}

	if r.IsPropertyReference() {
		if r.HasPrimitiveBase() {
			base = lang.ToObject(base)
		}
		panic("TODO: properties")
	}

	return base.Value().(Environment).GetBindingValue(r.GetReferencedName().String(), r.IsStrictReference())
}

func (r *Reference) PutValue(lang.Value) errors.Error {
	panic("TODO: 7.3 Operations on Objects")
}

func (r *Reference) GetThisValue() lang.Value {
	if r.IsSuperReference() {
		return r.thisValue
	}

	return r.GetBase()
}

func (r *Reference) InitializeReferencedBinding(val lang.Value) errors.Error {
	return r.GetBase().(Environment).InitializeBinding(r.GetReferencedName().String(), val)
}
