package parser_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseFile(t *testing.T) {
	tests := []struct {
		path       string
		successful bool
	}{
		{"test/parseable/p001.js", true},
		{"test/unparseable/up001.js", false},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			require := require.New(t)

			p := parser.New()
			err := p.ParseFile(tt.path)
			if tt.successful {
				require.NoError(err)
			} else {
				require.Error(err)
			}
		})
	}
}
