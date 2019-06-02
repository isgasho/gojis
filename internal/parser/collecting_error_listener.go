package parser

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type CollectingErrorListener struct {
	errs []error
}

func NewCollectingErrorListener() *CollectingErrorListener {
	l := new(CollectingErrorListener)
	return l
}

func (l *CollectingErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.Append(fmt.Errorf("Syntax error at line %v column %v: %v", line, column, msg))
}

func (l *CollectingErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (l *CollectingErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (l *CollectingErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
}

func (l *CollectingErrorListener) Append(errs ...error) {
	l.errs = append(l.errs, errs...)
}

func (l *CollectingErrorListener) Errors() ([]error, bool) {
	return l.errs, l.errs != nil && len(l.errs) > 0
}
