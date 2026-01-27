package goadapter

import (
	"errors"
	"go/ast"

	"github.com/amonvix/go-doc-agent/internal/semantic/adapter"
	"github.com/amonvix/go-doc-agent/internal/semantic/model"
)

type Adapter struct{}

var _ adapter.Adapter = (*Adapter)(nil)

func (a *Adapter) Language() string { return "go" }

func (a *Adapter) Build(input any) (*model.Project, error) {
	astFile, ok := input.(*ast.File)
	if !ok || astFile == nil {
		return nil, errors.New("goadapter: invalid input type, expected *ast.File")
	}

	return Build(astFile), nil
}
