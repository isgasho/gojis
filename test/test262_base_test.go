package test262

import (
	"bytes"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	test262repo = "https://github.com/tc39/test262"
)

func CloneTest262Repo(t *testing.T) {
	require := require.New(t)

	var stdout, stderr bytes.Buffer

	cmd := exec.Command("git", "clone", test262repo)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	require.NoErrorf(err, "Stdout: '%v'\nStderr: '%v'", stdout.String(), stderr.String())
}
