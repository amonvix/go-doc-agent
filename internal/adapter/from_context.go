package adapter

import (
	"github.com/amonvix/go-doc-agent/internal/context"
	"github.com/amonvix/go-doc-agent/internal/semantic"
)

func FromContext(ctx *context.Project) *semantic.Project {

	project := &semantic.Project{}

	for _, fn := range ctx.Functions {

		semanticFn := semantic.Function{
			Name:     fn.Name,
			FilePath: fn.FilePath,

			IsMethod:     fn.IsMethod,
			IsPure:       fn.IsPure,
			IsEntryPoint: fn.IsEntryPoint,

			Role:  semantic.RoleUnknown,
			Layer: semantic.LayerUnknown,
		}

		// -------------------------
		// Params
		// -------------------------
		for _, p := range fn.Params {
			semanticFn.Params = append(semanticFn.Params, semantic.Param{
				Name: p.Name,
				Type: semantic.TypeInfo{
					Name: p.Name,
				},
			})
		}

		// -------------------------
		// Returns
		// -------------------------
		for _, r := range fn.Returns {
			semanticFn.Returns = append(semanticFn.Returns, semantic.Return{
				Type: semantic.TypeInfo{
					Name: r.Type.Name,
				},
			})
		}

		// -------------------------
		// Calls
		// -------------------------
		for _, c := range fn.Calls {
			semanticFn.Calls = append(semanticFn.Calls, semantic.Call{
				Name:       c.Name,
				Target:     c.Target,
				Package:    c.Package,
				IsExternal: c.IsExternal,
				Line:       c.Line,
			})
		}

		// -------------------------
		// Dependencies
		// -------------------------
		for _, d := range semanticFn.Dependencies {
			semanticFn.Dependencies = append(
				semanticFn.Dependencies,
				semantic.Dependency{
					Name:    d.Name,
					Package: d.Package,
					Type:    d.Type,
				},
			)
		}

		// -------------------------
		// Side effects
		// -------------------------
		for _, s := range semanticFn.SideEffects {
			semanticFn.SideEffects = append(semanticFn.SideEffects, semantic.SideEffect{
				Type:   s.Type,
				Source: s.Source,
			})
		}

		project.Functions = append(project.Functions, semanticFn)
	}

	return project
}
