package analyzer

import (
	"github.com/amonvix/go-doc-agent/internal/semantic"
)

// InferFunctionLayer determines the architectural layer of a function
// based on its dependencies and semantic role.
func DetectFunctionLayer(fn *semantic.Function) {

	for _, dep := range fn.Dependencies {
		switch dep.Type {
		case semantic.DependencyDatabase,
			semantic.DependencyNetwork,
			semantic.DependencyFile,
			semantic.DependencyRuntime:
			fn.Layer = semantic.LayerInfrastructure
			return
		}
	}

	switch fn.Role {
	case semantic.RoleHandler:
		fn.Layer = semantic.LayerInterface
	case semantic.RoleService, semantic.RoleValidator:
		fn.Layer = semantic.LayerApplication
	default:
		if fn.IsPure && len(fn.Dependencies) == 0 {
			fn.Layer = semantic.LayerDomain
			return
		}
	}

	fn.Layer = semantic.LayerUnknown
}
