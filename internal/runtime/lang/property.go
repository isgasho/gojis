package lang

// Available field names for a property.
//
// A property is a data property descriptor if the fields
// 'Value', 'Writable' are set.
//
// A property is an accessor property descriptor if the fields
// 'Get', 'Set' are set.
//
// Every property must have the fields 'Enumerable', 'Configurable'
// set.
const (
	FieldNameValue        = "Value"
	FieldNameWritable     = "Writable"
	FieldNameGet          = "Get"
	FieldNameSet          = "Set"
	FieldNameEnumerable   = "Enumerable"
	FieldNameConfigurable = "Configurable"
)

// Property describes a property descriptor that holds
// 4 fields, ['Value', 'Writable', 'Enumerable', 'Configurable']
// or ['Get', 'Set', 'Enumerable', 'Configurable'].
type Property struct {
	*Record
}

// NewProperty creates a new, completely empty
// property record. The fields 'Value', 'Writable',
// 'Get', 'Set', 'Enumerable', 'Configurable' are NOT
// set yet.
func NewProperty() *Property {
	p := new(Property)
	p.Record = NewRecord()
	return p
}

// NewPropertyBase creates a new property base that is neither
// a data nor an accessor property. Only the fields 'Enumerable',
// 'Configurable' are set. A property created by this function
// is not bound to be a data or an accessor property. It can become
// either by adding the respective fields to the record.
func NewPropertyBase(enumerable, configurable Boolean) *Property {
	p := NewProperty()
	p.SetField(FieldNameEnumerable, enumerable)
	p.SetField(FieldNameConfigurable, configurable)
	return p
}

// NewDataProperty creates a new property that is a data property descriptor.
// The fields 'Value', 'Writable', 'Enumerable', 'Configurable' are set.
func NewDataProperty(value Value, writable, enumerable, configurable Boolean) *Property {
	p := NewPropertyBase(enumerable, configurable)
	p.SetField(FieldNameValue, value)
	p.SetField(FieldNameWritable, writable)
	return p
}

// NewAccessorProperty creates a new property that is an accessor property descriptor.
// The fields 'Get', 'Set', 'Enumerable', 'Configurable' are set.
func NewAccessorProperty(get, set *Object, enumerable, configurable Boolean) *Property {
	p := NewPropertyBase(enumerable, configurable)
	p.SetField(FieldNameGet, get)
	p.SetField(FieldNameSet, set)
	return p
}

// Value returns the value of this property, or Undefined
// if the field 'Value' is not set.
func (p *Property) Value() Value {
	val, ok := p.GetField(FieldNameValue)
	if !ok {
		return Undefined
	}
	return val.(Value)
}

// Writable returns the value of the field 'Writable', or False
// if the field is not set.
func (p *Property) Writable() Boolean {
	val, ok := p.GetField(FieldNameWritable)
	if !ok {
		return False
	}
	return val.(Boolean)
}

// Get returns the value of the field 'Get', or False
// if the field is not set.
func (p *Property) Get() Value {
	val, ok := p.GetField(FieldNameGet)
	if !ok {
		return Undefined
	}
	return val.(Value)
}

// Set returns the value of the field 'Set', or False
// if the field is not set.
func (p *Property) Set() Value {
	val, ok := p.GetField(FieldNameSet)
	if !ok {
		return Undefined
	}
	return val.(Value)
}

// Enumerable returns the value of the field 'Enumerable', or False
// if the field is not set.
func (p *Property) Enumerable() Boolean {
	val, ok := p.GetField(FieldNameEnumerable)
	if !ok {
		return False
	}
	return val.(Boolean)
}

// Configurable returns the value of the field 'Configurable', or False
// if the field is not set.
func (p *Property) Configurable() Boolean {
	val, ok := p.GetField(FieldNameConfigurable)
	if !ok {
		return False
	}
	return val.(Boolean)
}

// IsAccessorDescriptor determines whether the property is
// an accessor property descriptor.
// This is the case if the fields 'Get', 'Set' of the property
// are set.
func (p *Property) IsAccessorDescriptor() Boolean {
	_, hasGet := p.GetField(FieldNameGet)
	_, hasSet := p.GetField(FieldNameSet)
	if !hasGet && !hasSet {
		return False
	}

	return True
}

// IsDataDescriptor determines whether the property is
// an data property descriptor.
// This is the case if the fields 'Value', 'Writable' of the property
// are set.
func (p *Property) IsDataDescriptor() Boolean {
	_, hasGet := p.GetField(FieldNameValue)
	_, hasSet := p.GetField(FieldNameWritable)
	if !hasGet && !hasSet {
		return False
	}

	return True
}

// IsGenericDescriptor determines whether the property is
// a generic property descriptor.
// This is the case if the property is neither an accessor
// property descriptor nor a data property descriptor.
func (p *Property) IsGenericDescriptor() Boolean {
	if !p.IsAccessorDescriptor() && !p.IsDataDescriptor() {
		return True
	}

	return False
}

// FromPropertyDescriptor TODO:
func FromPropertyDescriptor(desc *Property) *Object {
	panic("TODO")
}

// ToPropertyDescriptor TODO:
func ToPropertyDescriptor(obj *Object) *Property {
	panic("TODO")
}

// CompletePropertyDescriptor TODO:
func CompletePropertyDescriptor(desc *Property) *Property {
	panic("TODO")
}
