package analyzer

import "github.com/amonvix/go-doc-agent/internal/semantic/model"

// InferFunctionLayer determines the architectural layer of a function
// based on its dependencies and semantic role.
func InferFunctionLayer(fn model.Function) model.Layer {

	// 1. Infrastructure layer — touches external systems
	for _, dep := range fn.Dependencies {
		switch dep.Type {
		case model.DependencyDatabase,
			model.DependencyNetwork,
			model.DependencyFile,
			model.DependencyRuntime:
			return model.LayerInfrastructure
		}
	}

	// 2. Interface layer — receives external input
	if fn.Role == model.RoleHandler {
		return model.LayerInterface
	}

	// 3. Application layer — orchestrates use cases
	if fn.Role == model.RoleService {
		return model.LayerApplication
	}

	// 4. Domain layer — pure business logic
	if fn.IsPure && len(fn.Dependencies) == 0 {
		return model.LayerDomain
	}

	return model.LayerUnknown
}
