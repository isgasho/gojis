package conformance

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
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
	if testing.Short() {
		// don't clone if short testing
		return
	}

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
			panic("TODO")
		})
	}
}
