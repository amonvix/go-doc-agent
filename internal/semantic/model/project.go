package model

type Project struct {
	Functions    []Function
	Structs      []Struct
	Dependencies []Dependency
}
