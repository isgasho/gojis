package lang

var _ Value = (*Symbol)(nil)

// Well-known symbol descriptions as specified in 6.1.5.1.
var (
	SymbolAsyncIterator      = Symbol{NewString("Symbol.asyncIterator")}
	SymbolHasInstance        = Symbol{NewString("Symbol.hasInstance")}
	SymbolIsConcatSpreadable = Symbol{NewString("Symbol.isConcatSpreadable")}
	SymbolIterator           = Symbol{NewString("Symbol.iterator")}
	SymbolMatch              = Symbol{NewString("Symbol.match")}
	SymbolReplace            = Symbol{NewString("Symbol.replace")}
	SymbolSearch             = Symbol{NewString("Symbol.search")}
	SymbolSpecies            = Symbol{NewString("Symbol.species")}
	SymbolSplit              = Symbol{NewString("Symbol.split")}
	SymbolToPrimitive        = Symbol{NewString("Symbol.toPrimitive")}
	SymbolToStringTag        = Symbol{NewString("Symbol.toStringTag")}
	SymbolUnscopables        = Symbol{NewString("Symbol.unscopables")}
)

// Symbol is a language type as specified by the language spec.
type Symbol struct {
	Description Value // either Undefined or a String
}

// Value returns the Value of this symbol's description.
func (s Symbol) Value() interface{} {
	return s.Description.Value()
}

// Type returns lang.TypeSymbol.
func (Symbol) Type() Type { return TypeSymbol }

// String returns a lang.String, containing the Symbols description.
// If the symbol's description is not a string, this will panic.
func (s Symbol) String() String {
	return s.Description.(String)
}
