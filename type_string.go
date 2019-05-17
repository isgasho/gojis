package vm

import "unicode/utf16"

// making sure that String implements Value interface
var (
	_ Value = NewString("")
	_ Value = String{1, 2}
)

type String []uint16

func NewString(str string) String {
	return String(utf16.Encode([]rune(str)))
}

func (s String) Value() interface{} {
	return string(utf16.Decode(s))
}

func (String) Type() Type { return TypeString }
