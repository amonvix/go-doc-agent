package analyzer

import (
	"strings"

	"github.com/amonvix/go-doc-agent/internal/semantic"
)

func DetectDependencies(fn *semantic.Function) {
	for _, call := range fn.Calls {

		pkg := strings.ToLower(call.Package)

		switch {
		case strings.Contains(pkg, "sql"),
			strings.Contains(pkg, "gorm"),
			strings.Contains(pkg, "mongo"):
			fn.Dependencies = append(fn.Dependencies, semantic.Dependency{
				Name:    pkg,
				Type:    semantic.DependencyDatabase,
				Package: call.Package,
			})

		case strings.Contains(pkg, "http"):
			fn.Dependencies = append(fn.Dependencies, semantic.Dependency{
				Name:    pkg,
				Type:    semantic.DependencyNetwork,
				Package: call.Package,
			})
		}
	}
}
