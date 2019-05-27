package runtime

import (
	"os"

	"github.com/rs/zerolog"
)

type Runtime struct {
	log zerolog.Logger
}

type RuntimeOption func(*Runtime)

func Trace() RuntimeOption {
	return func(r *Runtime) {
		r.log.UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Caller().Stack()
		})
	}
}

func New(opts ...RuntimeOption) *Runtime {
	r := new(Runtime)
	r.log = zerolog.New(os.Stdout).
		With().
		Timestamp().
		Str("component", "runtime").
		Logger()

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func (r *Runtime) SaySomething() {
	r.log.Info().Msg("Hello World!")
}
