package context

type TypeKind string

const (
	TypePrimitive TypeKind = "primitive"
	TypeStruct    TypeKind = "struct"
	TypePointer   TypeKind = "pointer"
	TypeSlice     TypeKind = "slice"
	TypeMap       TypeKind = "map"
	TypeInterface TypeKind = "interface"
	TypeExternal  TypeKind = "external"
	TypeUnknown   TypeKind = "unknown"
)

type TypeInfo struct {
	Name string
	Kind TypeKind
}
