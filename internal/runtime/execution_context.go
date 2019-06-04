package runtime

import (
	"gitlab.com/gojis/vm/internal/runtime/binding"
	"gitlab.com/gojis/vm/internal/runtime/lang"
	"gitlab.com/gojis/vm/internal/runtime/realm"
)

type ExecutionContext struct {
	Function       lang.Value // Object or Null
	Realm          *realm.Realm
	ScriptOrModule lang.InternalValue

	LexicalEnvironment  binding.Environment
	VariableEnvironment binding.Environment

	Generator interface{} // TODO: Table 23, GeneratorObject
}
