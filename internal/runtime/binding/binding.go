package binding

import "github.com/TimSatke/gojis/internal/runtime/lang"

type Binding struct {
	name      lang.String
	value     lang.Value
	immutable bool
	final     bool
	deletable bool
	strict    bool
}

func NewBinding(name lang.String) *Binding {
	b := new(Binding)
	b.name = name
	return b
}

func (b *Binding) Name() lang.String {
	return b.name
}

func (b *Binding) Set(val lang.Value) {
	b.value = val
}

func (b *Binding) Value() lang.Value {
	return b.value
}

func (b *Binding) IsStrict() bool {
	return b.strict
}

func (b *Binding) IsFinal() bool {
	return b.final
}

func (b *Binding) IsInitialized() bool {
	return b.value != nil
}

func (b *Binding) IsDeletable() bool {
	return b.deletable
}

func (b *Binding) IsImmutable() bool {
	return b.immutable
}
