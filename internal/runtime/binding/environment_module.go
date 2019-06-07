package binding

import "github.com/TimSatke/gojis/internal/runtime/lang"

type ModuleEnvironment struct {
	*DeclarativeEnvironment
}

func (e *ModuleEnvironment) CreateImportBinding(n lang.String, m interface{}, n2 lang.String) {
	panic("TODO: modules")
}

func (e *ModuleEnvironment) GetThisBinding() lang.Value {
	return lang.Undefined
}
