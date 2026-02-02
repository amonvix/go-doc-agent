package analyzer

import "github.com/amonvix/go-doc-agent/internal/semantic"

func Analyze(project *semantic.Project) {
	for i := range project.Functions {
		fn := &project.Functions[i]

		DetectFunctionRole(fn)
		DetectFunctionLayer(fn)
		DetectDependencies(fn)
		DetectSideEffects(fn)
		DetectEntrypoint(fn)
	}
}
