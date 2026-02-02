package generator

import (
	"strings"

	"github.com/amonvix/go-doc-agent/internal/semantic"
)

// buildFunctionSummary returns a short, human-readable summary for a function.
// Keep it dumb for now; we can get smarter later.
func buildFunctionSummary(fn semantic.Function) string {

	var b strings.Builder

	b.WriteString("**Function:** ")
	b.WriteString(fn.Name)
	b.WriteString("\n\n")

	if fn.Role != semantic.RoleUnknown {
		b.WriteString("- Role: ")
		b.WriteString(string(fn.Role))
		b.WriteString("\n")
	}

	if fn.Layer != semantic.LayerUnknown {
		b.WriteString("- Layer: ")
		b.WriteString(string(fn.Layer))
		b.WriteString("\n")
	}

	if fn.IsEntryPoint {
		b.WriteString("- Entry point\n")
	}

	if len(fn.Dependencies) > 0 {
		b.WriteString("\n**Dependencies:**\n")
		for _, d := range fn.Dependencies {
			b.WriteString("- ")
			b.WriteString(d.Name)
			b.WriteString(" (")
			b.WriteString(string(d.Type))
			b.WriteString(")\n")
		}
	}

	if fn.IsPure {
		b.WriteString("\nThis function appears to be pure.\n")
	}

	return b.String()
}
