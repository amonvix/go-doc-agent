package model

// FunctionRole represents the semantic responsibility of a function.
// It is language-agnostic and produced by the semantic analyzer.
type FunctionRole string

const (
	// Handles external input (HTTP, CLI, events, etc.)
	RoleHandler FunctionRole = "handler"

	// Contains business logic and orchestration rules
	RoleService FunctionRole = "service"

	// Responsible for data persistence and retrieval
	RoleRepository FunctionRole = "repository"

	// Responsible for object creation and initialization
	RoleFactory FunctionRole = "factory"

	// Converts data between layers or representations
	RoleMapper FunctionRole = "mapper"

	// Validates input, state, or domain rules
	RoleValidator FunctionRole = "validator"

	// Stateless helper or pure function
	RoleUtility FunctionRole = "utility"

	// Role could not be inferred
	RoleUnknown FunctionRole = "unknown"
)
