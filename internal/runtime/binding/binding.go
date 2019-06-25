package binding

import "github.com/gojisvm/gojis/internal/runtime/lang"

// Binding represents a concept described throughout the specification,
// especially in context with environment records.
type Binding struct {
	name      lang.String
	value     lang.Value
	immutable bool
	final     bool
	deletable bool
	strict    bool
}

// NewBinding creates a new named binding with the given name.
func NewBinding(name lang.String) *Binding {
	b := new(Binding)
	b.name = name
	return b
}

// Name returns the name of this binding.
func (b *Binding) Name() lang.String {
	return b.name
}

// Set assigns a value to this binding.
func (b *Binding) Set(val lang.Value) {
	b.value = val
}

// Value returns the value of this binding.
func (b *Binding) Value() lang.Value {
	return b.value
}

// IsStrict is used to determine whether this binding is a strict binding.
func (b *Binding) IsStrict() bool {
	return b.strict
}

// IsFinal is used to determine whether this binding is final.
func (b *Binding) IsFinal() bool {
	return b.final
}

// IsInitialized is used to determine whether this binding is initialized.
func (b *Binding) IsInitialized() bool {
	return b.value != nil
}

// IsDeletable is used to determine whether this binding is deletable.
func (b *Binding) IsDeletable() bool {
	return b.deletable
}

// IsImmutable is used to determine whether this binding is immutable.
func (b *Binding) IsImmutable() bool {
	return b.immutable
}
