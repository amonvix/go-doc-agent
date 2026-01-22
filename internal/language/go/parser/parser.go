package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

func ParseFile(path string) (*ast.File, *token.FileSet, error) {
	fset := token.NewFileSet()

	info, err := os.Stat(path)
	if err != nil {
		return nil, nil, err
	}

	// 🟢 Case 1: unique file comment
	if !info.IsDir() {
		astFile, err := parser.ParseFile(
			fset,
			path,
			nil,
			parser.ParseComments,
		)
		if err != nil {
			return nil, nil, err
		}

		return astFile, fset, nil
	}

	// 🟢 Case 2: directory (includ ".")
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, nil, err
	}

	pkgs, err := parser.ParseDir(
		fset,
		absPath,
		nil,
		parser.ParseComments,
	)
	if err != nil {
		return nil, nil, err
	}

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			return file, fset, nil
		}
	}

	return nil, nil, fmt.Errorf("no go files found in directory: %s", path)
}
