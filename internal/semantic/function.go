package semantic

type Function struct {
	Metadata

	Name     string
	FilePath string

	Params  []Param
	Returns []Return

	Calls        []Call
	Dependencies []Dependency
	SideEffects  []SideEffect

	IsMethod bool
	IsPure   bool

	Role  FunctionRole
	Layer Layer

	IsEntryPoint   bool
	IsTerminal     bool
	IsOrchestrator bool
}
