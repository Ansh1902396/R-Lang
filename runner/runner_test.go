package runner

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRunEvaluatesSource(t *testing.T) {
	evaluated, parserErrors := Run("let double = fn(x) { x * 2; }; double(21);")
	if len(parserErrors) != 0 {
		t.Fatalf("unexpected parser errors: %v", parserErrors)
	}
	if evaluated == nil || evaluated.Inspect() != "42" {
		t.Fatalf("expected 42, got %v", evaluated)
	}
}

func TestRunFileRejectsWrongExtension(t *testing.T) {
	var stderr bytes.Buffer
	code := RunFile("program.txt", &stderr)

	if code == 0 {
		t.Fatal("expected non-zero exit code")
	}
	if stderr.String() != "unsupported file extension \".txt\"; expected \".rlang\"\n" {
		t.Fatalf("unexpected stderr: %q", stderr.String())
	}
}

func TestRunFileReportsParserErrors(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.rlang")
	if err := os.WriteFile(path, []byte("let = 10;"), 0644); err != nil {
		t.Fatal(err)
	}

	var stderr bytes.Buffer
	code := RunFile(path, &stderr)

	if code == 0 {
		t.Fatal("expected non-zero exit code")
	}
	if stderr.String() == "" {
		t.Fatal("expected parser errors on stderr")
	}
}
