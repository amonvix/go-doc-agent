package model

type Dependency struct {
	Name    string
	Package string
	Type    DependencyType
}

type DependencyType string

const (
	DependencyDatabase DependencyType = "database"
	DependencyNetwork  DependencyType = "network"
	DependencyFile     DependencyType = "file"
	DependencyRuntime  DependencyType = "runtime"
	DependencyUnknown  DependencyType = "unknown"
)
