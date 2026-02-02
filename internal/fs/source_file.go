package fs

import (
	"fmt"
	"io/fs"
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

var ignoredDirs = map[string]bool{
	".git":         true,
	"vendor":       true,
	"node_modules": true,
	".idea":        true,
	".vscode":      true,
}

func ScanDirectory(root string) ([]SourceFile, error) {

	info, err := os.Stat(root)
	if err != nil {
		return nil, fmt.Errorf("input not found: %w", err)
	}

	if !info.IsDir() {
		return nil, fmt.Errorf("expected directory, got file: %s", root)
	}

	var files []SourceFile

	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		// Ignore folders
		if d.IsDir() {
			if ignoredDirs[d.Name()] {
				return filepath.SkipDir
			}
			return nil
		}

		// Temporary filter (language detector later)
		if filepath.Ext(path) != ".go" {
			return nil
		}

		file, err := ScanFile(path)
		if err != nil {
			return err
		}

		files = append(files, *file)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
