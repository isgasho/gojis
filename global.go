package vm

type global struct {
	Object,
	Function,
	Array,
	String,
	Boolean,
	Number,
	Math,
	Date,
	RegExp,
	Error,
	EvalError,
	TypeError,
	RangeError,
	ReferenceError,
	SyntaxError,
	URIError,
	JSON *object

	ObjectPrototype,
	FunctionPrototype,
	ArrayPrototype,
	StringPrototype,
	BooleanPrototype,
	NumberPrototype,
	DatePrototype,
	RegExpPrototype,
	ErrorPrototype,
	EvalErrorPrototype,
	TypeErrorPrototype,
	RangeErrorPrototype,
	ReferenceErrorPrototype,
	SyntaxErrorPrototype,
	URIErrorPrototype *object
}
