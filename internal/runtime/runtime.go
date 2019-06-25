package runtime

import (
	"github.com/rs/zerolog"
	"github.com/gojisvm/gojis/internal/parser"
)

// Runtime is an object that will evaluate
// an AST.
type Runtime struct {
	log zerolog.Logger

	ast *parser.Ast
}

// New creates a new runtime using the given logger and
// evaluating the given AST.
func New(log zerolog.Logger, ast *parser.Ast) *Runtime {
	r := new(Runtime)
	r.log = log
	r.ast = ast
	return r
}
