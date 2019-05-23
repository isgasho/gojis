package lang

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestNewProperty ensures that if a property is
// created with NewProperty, it does not have any
// fields set, especially not default property fields.
func TestNewProperty(t *testing.T) {
	require := require.New(t)

	p := NewProperty()
	require.Empty(p.Record.fields)
}

// TestNewPropertyBase ensures that a property
// that has been created with NewPropertyBase
// only has fields set that belong to both a
// data and an accessor property descriptor.
// These are the fields 'Enumerable' and 'Configurable'.
func TestNewPropertyBase(t *testing.T) {
	require := require.New(t)

	p := NewPropertyBase(False, True)
	require.Len(p.Record.fields, 2)

	// Enumerable field must be set although it is False
	field, ok := p.GetField(FieldNameEnumerable)
	require.True(ok)
	require.Equal(False, field)

	// Configurable field must be set and be True
	field, ok = p.GetField(FieldNameConfigurable)
	require.True(ok)
	require.Equal(True, field)
}

// TestNewDataProperty tests the creation of a new data property.
func TestNewDataProperty(t *testing.T) {
	require := require.New(t)

	p := NewDataProperty(NewString("foobar"), False, True, True)
	require.Len(p.Record.fields, 4)

	field, ok := p.GetField(FieldNameValue)
	require.True(ok)
	require.Equal(NewString("foobar"), field)

	field, ok = p.GetField(FieldNameWritable)
	require.True(ok)
	require.Equal(False, field)

	field, ok = p.GetField(FieldNameEnumerable)
	require.True(ok)
	require.Equal(True, field)

	field, ok = p.GetField(FieldNameConfigurable)
	require.True(ok)
	require.Equal(True, field)
}

// TestNewAccessorProperty tests the creation of a new accessor property.
func TestNewAccessorProperty(t *testing.T) {
	require := require.New(t)

	p := NewAccessorProperty(func() Value { return nil }, nil, True, False)
	require.Len(p.Record.fields, 4)

	field, ok := p.GetField(FieldNameGet)
	require.True(ok)
	require.NotNil(field)

	field, ok = p.GetField(FieldNameSet)
	require.True(ok)
	require.Nil(field)

	field, ok = p.GetField(FieldNameEnumerable)
	require.True(ok)
	require.Equal(True, field)

	field, ok = p.GetField(FieldNameConfigurable)
	require.True(ok)
	require.Equal(False, field)
}

// TestPropertyValue ensures that a property's method
// Value returns Undefined if the Value field is not set.
// It also ensures that Null is returned if the Value is Null,
// and that it returns a specific Value if it is set.
func TestPropertyValue(t *testing.T) {
	t.Run("Value Undefined", func(t *testing.T) {
		require := require.New(t)

		p := NewProperty()
		require.Equal(Undefined, p.Value())
	})

	t.Run("Accessor Value Undefined", func(t *testing.T) {
		require := require.New(t)

		p := NewAccessorProperty(func() Value { return nil }, nil, True, False)
		require.Equal(Undefined, p.Value())
	})

	t.Run("Value Null", func(t *testing.T) {
		require := require.New(t)

		p := NewDataProperty(Null, False, False, False)
		require.Equal(Null, p.Value())
	})

	t.Run("Value", func(t *testing.T) {
		require := require.New(t)

		p := NewDataProperty(NewString("foobar"), False, False, False)
		require.Equal(NewString("foobar"), p.Value())
	})
}
func TestPropertyEnumerable(t *testing.T) {
	t.Run("Enumerable Undefined", func(t *testing.T) {
		require := require.New(t)

		p := NewProperty()
		require.Equal(False, p.Enumerable())
	})

	t.Run("Enumerable False", func(t *testing.T) {
		require := require.New(t)

		p := NewDataProperty(NewString("foobar"), False, False, False)
		require.Equal(False, p.Enumerable())
	})

	t.Run("Enumerable True", func(t *testing.T) {
		require := require.New(t)

		p := NewDataProperty(NewString("foobar"), False, True, False)
		require.Equal(True, p.Enumerable())
	})
}

func TestPropertyConfigurable(t *testing.T) {
	t.Run("Configurable Undefined", func(t *testing.T) {
		require := require.New(t)

		p := NewProperty()
		require.Equal(False, p.Configurable())
	})

	t.Run("Configurable False", func(t *testing.T) {
		require := require.New(t)

		p := NewDataProperty(NewString("foobar"), False, False, False)
		require.Equal(False, p.Configurable())
	})

	t.Run("Configurable True", func(t *testing.T) {
		require := require.New(t)

		p := NewDataProperty(NewString("foobar"), False, False, True)
		require.Equal(True, p.Configurable())
	})
}

func TestPropertyWritable(t *testing.T) {
	t.Run("Writable Undefined", func(t *testing.T) {
		require := require.New(t)

		p := NewProperty()
		require.Equal(False, p.Writable())
	})

	t.Run("Writable False", func(t *testing.T) {
		require := require.New(t)

		p := NewDataProperty(NewString("foobar"), False, False, False)
		require.Equal(False, p.Writable())
	})

	t.Run("Writable True", func(t *testing.T) {
		require := require.New(t)

		p := NewDataProperty(NewString("foobar"), True, False, False)
		require.Equal(True, p.Writable())
	})
}

func TestPropertyGet(t *testing.T) {
	t.Run("Get Property", func(t *testing.T) {
		require := require.New(t)

		p := NewProperty()
		require.Nil(p.Get())
	})

	t.Run("Get DataProperty", func(t *testing.T) {
		require := require.New(t)

		p := NewDataProperty(NewString("foobar"), False, False, False)
		require.Nil(p.Get())
	})

	t.Run("Get AccessorProperty nil", func(t *testing.T) {
		require := require.New(t)

		p := NewAccessorProperty(nil, nil, False, False)
		require.Nil(p.Get())
	})

	t.Run("Get AccessorProperty", func(t *testing.T) {
		require := require.New(t)

		p := NewAccessorProperty(func() Value { return nil }, nil, False, False)
		require.NotNil(p.Get())
	})
}

func TestPropertySet(t *testing.T) {
	t.Run("Set Property", func(t *testing.T) {
		require := require.New(t)

		p := NewProperty()
		require.Nil(p.Set())
	})

	t.Run("Set DataProperty", func(t *testing.T) {
		require := require.New(t)

		p := NewDataProperty(NewString("foobar"), False, False, False)
		require.Nil(p.Set())
	})

	t.Run("Set AccessorProperty nil", func(t *testing.T) {
		require := require.New(t)

		p := NewAccessorProperty(nil, nil, False, False)
		require.Nil(p.Set())
	})

	t.Run("Set AccessorProperty", func(t *testing.T) {
		require := require.New(t)

		p := NewAccessorProperty(nil, func(Value) Boolean { return False }, False, False)
		require.NotNil(p.Set())
	})
}

func TestPropertyIsXXXDescriptor(t *testing.T) {
	prop := NewProperty()
	propBase := NewPropertyBase(False, False)
	propData := NewDataProperty(Null, False, False, False)
	propAccess := NewAccessorProperty(func() Value { return nil }, func(Value) Boolean { return False }, False, False)

	t.Run("IsAccessorDescriptor", func(t *testing.T) {
		require := require.New(t)

		require.Equal(False, prop.IsAccessorDescriptor())
		require.Equal(False, propBase.IsAccessorDescriptor())
		require.Equal(False, propData.IsAccessorDescriptor())
		require.Equal(True, propAccess.IsAccessorDescriptor())
	})

	t.Run("IsDataDescriptor", func(t *testing.T) {
		require := require.New(t)

		require.Equal(False, prop.IsDataDescriptor())
		require.Equal(False, propBase.IsDataDescriptor())
		require.Equal(True, propData.IsDataDescriptor())
		require.Equal(False, propAccess.IsDataDescriptor())
	})

	t.Run("IsGenericDescriptor", func(t *testing.T) {
		require := require.New(t)

		require.Equal(True, prop.IsGenericDescriptor())
		require.Equal(True, propBase.IsGenericDescriptor())
		require.Equal(False, propData.IsGenericDescriptor())
		require.Equal(False, propAccess.IsGenericDescriptor())
	})
}
