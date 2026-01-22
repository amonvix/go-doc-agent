package context

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
	Name string
	Kind TypeKind
}

// ---------------------------
// Function model
// ---------------------------

type Function struct {
	Name        string
	Description string

	FilePath string

	Params  []Param
	Returns []Return
}

// ---------------------------
// Parameters
// ---------------------------

type Param struct {
	Name        string
	Type        TypeInfo
	Description string
}

// ---------------------------
// Return values
// ---------------------------

type Return struct {
	Type        TypeInfo
	Description string
}

// ---------------------------
// Project model
// ---------------------------

type Project struct {
	Path      string
	Functions []Function
}
