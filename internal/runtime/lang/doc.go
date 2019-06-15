// Package lang contains all ECMAScript-specified language values,
// i.e. String, Undefined, Null, Symbol etc., as well as Object and
// some types needed for internal implementation of their behaviours.
// These additional types are Record, StringAndSymbol (not exhaustive).
//
// The interface Value defined by this package describes any ECMAScript
// language value (it is referenced as such by the specification) and describes
// a value that can be accessed and manipulated by ECMAScript code.
// Since the specification also uses Undefined in some cases for other than
// ECMAScript language values, this package also declares the interface
// InternalValue, which represents any value referenced in the specification.
//
// Sometimes, something different than an InternalValue (e.g. a pointer) is used.
// This often happens, when the language specification states that the object
// is either provided or is Null. In this case, nil represents the case that
// the value is Null, and a pointer to an object or property (etc.) represents
// the case that the value is present.
//
// When using Boolean values, you must not use Boolean(true) to create a Boolean
// that represents the value True (although this currently would work).
// Instead, use the provided constants True and False.
//
// You cannot use the type StringOrSymbol as a replacement for a String or Symbol
// object. Only use a StringOrSymbol object when necessary.
// The long time goal is, to eliminate the type StringOrSymbol, as it is not a nice
// solution, but at this point necessary to follow the specification as exactly as possible
// in many other points.
//
// As for symbols, you should not define new Symbols if not absolutely necessary.
// This package already defines all symbols necessary to follow the specification.
//
// If you require a bool value, but a function of this package returns a Boolean object,
// there most likely exists an identical function with an Internal prefix (e.g.
// SameValue(a, b) Boolean and InternalSameValue(a, b) bool). Use the Internal function
// instead, since most of the time, the non-internal function uses the return value of
// the internal function and converts the type to a Boolean object.
package lang
