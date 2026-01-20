package parser

import (
	"go/ast"

	"github.com/amonvix/go-doc-agent/internal/context"
)

func parseType(expr ast.Expr) context.TypeInfo {
	switch t := expr.(type) {

	case *ast.Ident:
		return context.TypeInfo{
			Name: t.Name,
			Kind: context.TypePrimitive,
		}

	case *ast.StarExpr:
		inner := parseType(t.X)
		return context.TypeInfo{
			Name: "*" + inner.Name,
			Kind: context.TypePointer,
		}

	case *ast.ArrayType:
		inner := parseType(t.Elt)
		return context.TypeInfo{
			Name: "[]" + inner.Name,
			Kind: context.TypeSlice,
		}

	case *ast.SelectorExpr:
		return context.TypeInfo{
			Name: exprToString(expr),
			Kind: context.TypeExternal,
		}

	case *ast.MapType:
		return context.TypeInfo{
			Name: "map[" + exprToString(t.Key) + "]" + exprToString(t.Value),
			Kind: context.TypeMap,
		}

	case *ast.InterfaceType:
		return context.TypeInfo{
			Name: "interface{}",
			Kind: context.TypeInterface,
		}

	default:
		return context.TypeInfo{
			Name: "unknown",
			Kind: context.TypeUnknown,
		}
	}
}
