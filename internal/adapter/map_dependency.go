package adapter

import "github.com/amonvix/go-doc-agent/internal/semantic"

func mapDependency(name string) semantic.Dependency {
	return semantic.Dependency{
		Name: name,
		Type: semantic.DependencyUnknown,
	}
}
