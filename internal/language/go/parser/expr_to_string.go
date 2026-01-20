package parser

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
)

// exprToString converts a Go AST expression into its textual representation.
// Deterministic, offline, no semantic analysis.
func exprToString(expr ast.Expr) string {
	if expr == nil {
		return ""
	}

	var buf bytes.Buffer
	_ = printer.Fprint(&buf, token.NewFileSet(), expr)

	return buf.String()
}
