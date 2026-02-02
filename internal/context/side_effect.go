package context

type SideEffectType string

const (
	SideEffectIO       SideEffectType = "io"
	SideEffectDatabase SideEffectType = "database"
	SideEffectNetwork  SideEffectType = "network"
	SideEffectFile     SideEffectType = "file"
)

type SideEffect struct {
	Type   SideEffectType
	Source string
}
