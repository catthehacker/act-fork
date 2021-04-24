package runner

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/nektos/act/pkg/common"
)

func TestStepContextExecutor(t *testing.T) {
	platforms := map[string]string{
		"ubuntu-latest": "node:12.20.1-buster-slim",
	}
	tables := []TestJobFileInfo{
		{"testdata", "uses-and-run-in-one-step", "push", "Invalid run/uses syntax for job:test step:Test", platforms, "linux/amd64"},
		{"testdata", "uses-github-empty", "push", "Expected format {org}/{repo}[/path]@ref", platforms, "linux/amd64"},
		{"testdata", "uses-github-noref", "push", "Expected format {org}/{repo}[/path]@ref", platforms, "linux/amd64"},
		{"testdata", "uses-github-root", "push", "", platforms, "linux/amd64"},
		{"testdata", "uses-github-path", "push", "", platforms, "linux/amd64"},
	}
	// These tests are sufficient to only check syntax.
	ctx := common.WithDryrun(context.Background(), true)
	secrets, _ := godotenv.Read(filepath.Join("..", ".secrets"))
	for _, table := range tables {
		runTestJobFile(ctx, t, table, secrets)
	}
}
