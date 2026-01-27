package model

type Function struct {
	Metadata

	Name string

	Params  []string
	Returns []string

	Calls        []Call
	Dependencies []Dependency
	SideEffects  []SideEffect

	IsMethod bool
	IsPure   bool

	Role FunctionRole
	Layer Layer
	
	IsEntryPoint   bool
	IsTerminal     bool
	IsOrchestrator bool
}
