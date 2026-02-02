package language

import "github.com/amonvix/go-doc-agent/internal/semantic"

// Adapter translates language-specific syntax
// into language-agnostic semantic models.
type Adapter interface {

	// Language returns the language identifier
	// Example: "go", "python", "javascript"
	Language() string

	// Build converts parsed language structures
	// into semantic intermediate representation.
	//
	// Input must be the parser output of the language.
	// Output must be a semantic Project model.
	Build(input any) (*semantic.Project, error)
}
