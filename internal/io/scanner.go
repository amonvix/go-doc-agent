package io

import (
	"fmt"
	"os"
	"path/filepath"
)

type SourceFile struct {
	Path     string
	Filename string
	Content  []byte
}

func ScanFile(path string) (*SourceFile, error) {

	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("input not found: %w", err)
	}

	if info.IsDir() {
		return nil, fmt.Errorf("expected file, got directory: %s", path)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return &SourceFile{
		Path:     path,
		Filename: filepath.Base(path),
		Content:  content,
	}, nil
}
