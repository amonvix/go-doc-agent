package golang

import (
	"errors"
	"go/ast"

	"github.com/amonvix/go-doc-agent/internal/language"

	"github.com/amonvix/go-doc-agent/internal/semantic"
)

type Adapter struct{}

var _ language.Adapter = (*Adapter)(nil)

func (a *Adapter) Language() string { return "go" }

func (a *Adapter) Build(input any) (*semantic.Project, error) {
	astFile, ok := input.(*ast.File)
	if !ok || astFile == nil {
		return nil, errors.New("goadapter: invalid input type, expected *ast.File")
	}

	return Build(astFile), nil
}
