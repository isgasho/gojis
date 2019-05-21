package lang

import "unicode/utf16"

var _ Value = (*String)(nil)

type String []uint16

func NewString(str string) String {
	return String(utf16.Encode([]rune(str)))
}

func (s String) Value() interface{} {
	if s == nil {
		return nil
	}

	return string(utf16.Decode(s))
}

func (String) Type() Type { return TypeString }

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
