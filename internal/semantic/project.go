package semantic

type Project struct {
	Name string

	Functions    []Function
	Structs      []Struct
	Dependencies []Dependency
	Path         string
}
