package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveFileUnderRoot(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "2026")
	if err := os.Mkdir(dir, 0700); err != nil {
		t.Fatal(err)
	}
	file := filepath.Join(dir, "call.wav")
	if err := os.WriteFile(file, []byte("test"), 0600); err != nil {
		t.Fatal(err)
	}
	got, err := resolveFileUnderRoot(root, "2026", "call.wav")
	if err != nil || got != file {
		t.Fatalf("got %q, %v", got, err)
	}
	for _, parts := range [][]string{{"..", "secret"}, {`..\\secret`}, {"2026", "missing.wav"}} {
		if got, err := resolveFileUnderRoot(root, parts...); err == nil {
			t.Fatalf("unsafe path accepted: %q", got)
		}
	}
}
