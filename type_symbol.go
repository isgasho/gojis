package vm

type SymbolDescription String

// well-known symbol descriptions as specified in 6.1.5.1
var (
	SymbolDescriptionAsyncIterator      = NewString("Symbol.asyncIterator")
	SymbolDescriptionHasInstance        = NewString("Symbol.hasInstance")
	SymbolDescriptionIsConcatSpreadable = NewString("Symbol.isConcatSpreadable")
	SymbolDescriptionIterator           = NewString("Symbol.iterator")
	SymbolDescriptionMatch              = NewString("Symbol.match")
	SymbolDescriptionReplace            = NewString("Symbol.replace")
	SymbolDescriptionSearch             = NewString("Symbol.search")
	SymbolDescriptionSpecies            = NewString("Symbol.species")
	SymbolDescriptionSplit              = NewString("Symbol.split")
	SymbolDescriptionToPrimitive        = NewString("Symbol.toPrimitive")
	SymbolDescriptionToStringTag        = NewString("Symbol.toStringTag")
	SymbolDescriptionUnscopables        = NewString("Symbol.unscopables")
)

type Symbol struct {
	Description Value // either Undefined or a String
}
