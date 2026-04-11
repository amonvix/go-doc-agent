package fs

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestScanFile(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "sample.go")
	content := []byte("package main")
	if err := os.WriteFile(filePath, content, 0o644); err != nil {
		t.Fatalf("failed to write file: %v", err)
	}

	file, err := ScanFile(filePath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if file.Filename != "sample.go" {
		t.Fatalf("Filename = %q, want sample.go", file.Filename)
	}
	if string(file.Content) != string(content) {
		t.Fatalf("Content = %q, want %q", string(file.Content), string(content))
	}
}

func TestScanFile_Errors(t *testing.T) {
	tempDir := t.TempDir()

	_, err := ScanFile(filepath.Join(tempDir, "missing.go"))
	if err == nil || !strings.Contains(err.Error(), "input not found") {
		t.Fatalf("expected input not found error, got %v", err)
	}

	_, err = ScanFile(tempDir)
	if err == nil || !strings.Contains(err.Error(), "expected file, got directory") {
		t.Fatalf("expected directory error, got %v", err)
	}
}

func TestScanDirectory(t *testing.T) {
	tempDir := t.TempDir()
	if err := os.Mkdir(filepath.Join(tempDir, "pkg"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(tempDir, "vendor"), 0o755); err != nil {
		t.Fatal(err)
	}

	mustWrite := func(path, body string) {
		if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
			t.Fatalf("write %s: %v", path, err)
		}
	}

	mustWrite(filepath.Join(tempDir, "pkg", "included.go"), "package pkg")
	mustWrite(filepath.Join(tempDir, "pkg", "ignored.txt"), "not go")
	mustWrite(filepath.Join(tempDir, "vendor", "skipped.go"), "package vendor")

	files, err := ScanDirectory(tempDir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(files) != 1 {
		t.Fatalf("expected 1 go file, got %d", len(files))
	}
	if files[0].Filename != "included.go" {
		t.Fatalf("got file %q, want included.go", files[0].Filename)
	}
}

func TestScanDirectory_Errors(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "input.go")
	if err := os.WriteFile(filePath, []byte("package main"), 0o644); err != nil {
		t.Fatal(err)
	}

	_, err := ScanDirectory(filepath.Join(tempDir, "missing"))
	if err == nil || !strings.Contains(err.Error(), "input not found") {
		t.Fatalf("expected input not found error, got %v", err)
	}

	_, err = ScanDirectory(filePath)
	if err == nil || !strings.Contains(err.Error(), "expected directory, got file") {
		t.Fatalf("expected directory error, got %v", err)
	}
}
