package parser

import (
	"go/ast"

	"github.com/amonvix/go-doc-agent/internal/context"
)

func ExtractFunctions(file *ast.File) []context.Function
