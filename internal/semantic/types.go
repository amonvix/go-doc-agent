package semantic

// ---------------------------
// Type system
// ---------------------------

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
	Name      string
	Package   string
	IsPointer bool
	IsSlice   bool
	IsMap     bool
}

// ---------------------------
// Return values
// ---------------------------

type Return struct {
	Type TypeInfo
}
