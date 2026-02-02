package semantic

type SideEffectType string

const (
	SideEffectIO       SideEffectType = "io"
	SideEffectDatabase SideEffectType = "database"
	SideEffectNetwork  SideEffectType = "network"
	SideEffectFile     SideEffectType = "file"
	SideEffectUnknown  SideEffectType = "unknown"
)

type SideEffect struct {
	Type   SideEffectType
	Source string
}
