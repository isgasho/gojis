package agent

import (
	"github.com/gojisvm/gojis/internal/runtime/binding"
	"github.com/gojisvm/gojis/internal/runtime/lang"
	"github.com/gojisvm/gojis/internal/runtime/realm"
)

type ExecutionContext struct {
	Function       lang.Value // Object or Null
	Realm          *realm.Realm
	ScriptOrModule lang.InternalValue

	LexicalEnvironment  binding.Environment
	VariableEnvironment binding.Environment

	Generator interface{} // TODO: Table 23, GeneratorObject
}
