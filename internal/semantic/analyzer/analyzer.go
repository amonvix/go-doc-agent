package analyzer

import "github.com/amonvix/go-doc-agent/internal/semantic/model"

// Analyze walks through the semantic IR and enriches it with meaning.
// This layer does not understand syntax or language-specific constructs.
// It operates only on language-agnostic semantic models.
func Analyze(project *model.Project) {

	for i := range project.Functions {
		project.Functions[i].Role =
			DetectFunctionRole(project.Functions[i])
	}
}
