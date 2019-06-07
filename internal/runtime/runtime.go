package runtime

import (
	"github.com/rs/zerolog"
	"github.com/TimSatke/gojis/internal/parser"
)

type Runtime struct {
	log zerolog.Logger

	ast *parser.Ast
}

func New(log zerolog.Logger, ast *parser.Ast) *Runtime {
	r := new(Runtime)
	r.log = log
	r.ast = ast
	return r
}
