package model

type Call struct {
	Name    string
	Target  string
	Package string

	IsExternal bool

	Line int
}
