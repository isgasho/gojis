package objects

import (
	"gitlab.com/gojis/vm/runtime/lang"
)

func HasOwnProperty(o *lang.Object, k lang.StringOrSymbol) bool {
	return o.GetOwnProperty(k) != nil
}
