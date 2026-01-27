package model

// Layer represents the architectural layer inferred by the semantic analyzer.
type Layer string

const (
	LayerInterface      Layer = "interface"
	LayerApplication    Layer = "application"
	LayerDomain         Layer = "domain"
	LayerInfrastructure Layer = "infrastructure"
	LayerUnknown        Layer = "unknown"
)
