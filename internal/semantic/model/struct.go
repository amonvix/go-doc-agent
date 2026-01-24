package model

type Struct struct {
	Name    string
	Fields  []Field
	Methods []string

	File     string
	Line     int
	Language string
}

type Field struct {
	Name string
	Type string
}
