package adapter

import (
	"github.com/amonvix/go-doc-agent/internal/context"
	"github.com/amonvix/go-doc-agent/internal/semantic"
)

func mapSideEffectType(t context.SideEffectType) semantic.SideEffectType {
	switch t {
	case context.SideEffectDatabase:
		return semantic.SideEffectDatabase
	case context.SideEffectNetwork:
		return semantic.SideEffectNetwork
	case context.SideEffectFile:
		return semantic.SideEffectFile
	default:
		return semantic.SideEffectUnknown
	}
}
