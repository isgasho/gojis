package binding

import (
	"github.com/gojisvm/gojis/internal/runtime/errors"
	"github.com/gojisvm/gojis/internal/runtime/lang"
)

type Environment interface {
	lang.Value

	Outer() Environment

	HasBinding(n lang.String) bool
	CreateMutableBinding(n lang.String, deletable bool) errors.Error
	CreateImmutableBinding(n lang.String, strict bool) errors.Error
	InitializeBinding(n lang.String, val lang.Value) errors.Error
	SetMutableBinding(n lang.String, val lang.Value, strict bool) errors.Error
	GetThisBinding() (lang.Value, errors.Error)
	GetBindingValue(n lang.String, strict bool) (lang.Value, errors.Error)
	DeleteBinding(n lang.String) bool
	HasThisBinding() bool
	HasSuperBinding() bool
	WithBaseObject() lang.Value
}
