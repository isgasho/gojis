package parser_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gojis/vm/internal/parser"
)

func TestParseFile(t *testing.T) {
	tests := []struct {
		path       string
		successful bool
	}{
		{"parseable/p001.js", true},
		{"unparseable/up001.js", false},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			require := require.New(t)

			basePath := "../../test/parser"

			p := parser.New()
			err := p.ParseFile(filepath.Join(basePath, tt.path))
			if tt.successful {
				require.NoError(err)
			} else {
				require.Error(err)
			}
		})
	}
}
