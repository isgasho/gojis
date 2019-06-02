package lang

import "unicode/utf16"

var _ Value = (*String)(nil) // ensure that String implements Value

// String is a language type as specified by the language spec.
// All codepoints are UTF-16 codepoints and will be handled as such.
type String []uint16

// NewString creates a new String from a given string.
// The given string will be encoded to UTF-16 and
// will be returned as a lang.String.
func NewString(str string) String {
	return String(utf16.Encode([]rune(str)))
}

// Value returns a string representing the lang.String.
// The value of a nil String is nil.
func (s String) Value() interface{} {
	if s == nil {
		return nil
	}

	return string(utf16.Decode(s))
}

// Type returns lang.TypeString.
func (String) Type() Type { return TypeString }

// StringsEqual can be used to determine the equality of
// two strings. This function compares two given strings
// to be equalcodepoint by codepoint.
func StringsEqual(s1, s2 String) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
