package context

type Path string

func NewPath(value string) Path {
	return Path(value)
}

func (p Path) String() string {
	return string(p)
}
