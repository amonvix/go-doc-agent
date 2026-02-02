package analyzer

import (
	"github.com/amonvix/go-doc-agent/internal/semantic"
)

func DetectSideEffects(fn *semantic.Function) {

	if len(fn.Dependencies) == 0 {
		fn.IsPure = true
		return
	}

	fn.IsPure = false

	for _, dep := range fn.Dependencies {
		switch dep.Type {
		case semantic.DependencyDatabase:
			fn.SideEffects = append(fn.SideEffects, semantic.SideEffect{
				Type:   semantic.SideEffectDatabase,
				Source: dep.Name,
			})
		case semantic.DependencyNetwork:
			fn.SideEffects = append(fn.SideEffects, semantic.SideEffect{
				Type:   semantic.SideEffectNetwork,
				Source: dep.Name,
			})
		}
	}
}
