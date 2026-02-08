package semantic

// ParamType represents the semantic meaning of a function parameter type.
// It is intentionally a string to allow language-agnostic mapping
// (go, python, js, etc).
type ParamType string

const (
	ParamTypeUnknown ParamType = "unknown"

	ParamTypeString ParamType = "string"
	ParamTypeInt    ParamType = "int"
	ParamTypeFloat  ParamType = "float"
	ParamTypeBool   ParamType = "bool"

	ParamTypeSlice   ParamType = "slice"
	ParamTypeMap     ParamType = "map"
	ParamTypeStruct  ParamType = "struct"
	ParamTypePointer ParamType = "pointer"
	ParamTypeAny     ParamType = "any"
)

// ---------------------------
// Parameters
// ---------------------------

type Param struct {
	Name        string
	Type        TypeInfo
	Description string
}

func (t TypeInfo) String() string {
	return t.Name
}
