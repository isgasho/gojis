package conformance

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gojis/vm/internal/runtime"
)

const (
	test262repo = "https://github.com/tc39/test262"
)

func TestMain(m *testing.M) {
	setup()
	defer tearDown()

	os.Exit(m.Run())
}

func setup() {
	cloneTest262Repo()
}

func tearDown() {}

func cloneTest262Repo() {
	if _, err := os.Stat("test262"); os.IsNotExist(err) {
		log.Println("Conformance test directory does not exist, cloning it...")

		var stdout, stderr bytes.Buffer

		cmd := exec.Command("git", "clone", test262repo)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := cmd.Run()
		if err != nil {
			panic(fmt.Sprintf("Clone failed.\nStdout: '%v'\nStderr: '%v'", stdout.String(), stderr.String()))
		}
	}
}

type testCase struct {
	path                  string
	expectSuccessfulParse bool
	expectSuccessfulRun   bool
}

func runTable(t *testing.T, basePath string, testCases []testCase) {
	for _, tt := range testCases {
		t.Run(tt.path, func(t *testing.T) {
			require := require.New(t)
			var err error

			r := runtime.New()

			err = r.LoadFile(filepath.Join(basePath, tt.path))
			if tt.expectSuccessfulParse {
				require.NoError(err)
			} else {
				require.Error(err)
				return // abort if parse error occurred
			}

			err = r.Start()
			if tt.expectSuccessfulRun {
				require.NoError(err)
			} else {
				require.Error(err)
			}
		})
	}
}
