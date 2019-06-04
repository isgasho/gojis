package runtime

import (
	"github.com/rs/zerolog"
	"gitlab.com/gojis/vm/internal/parser"
)

type Runtime struct {
	log zerolog.Logger

	parser *parser.Parser
}

func New(log zerolog.Logger) *Runtime {
	r := new(Runtime)
	r.log = log
	r.parser = parser.New()
	return r
}
