package context

import "github.com/amonvix/go-doc-agent/internal/fs"

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
	KeyType   *TypeInfo
	ValueType *TypeInfo
}

// ---------------------------
// Function model
// ---------------------------

type Function struct {
	Name        string
	Description string
	FilePath    string

	Params  []Param
	Returns []Return

	Calls []Call

	IsMethod bool
	IsPure   bool

	IsEntryPoint bool
}

// ---------------------------
// Parameters
// ---------------------------

type Param struct {
	Name        string
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
	Files     []fs.SourceFile
	Functions []Function
}
