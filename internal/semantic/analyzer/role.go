/*
this file operates some actions, when:
receive a request + service call -> handler
receive the model +return another model -> mapper
calls databes -> repository
create structs -> factory
*/
package analyzer

import (
	"strings"

	"github.com/amonvix/go-doc-agent/internal/semantic/model"
)

func DetectFunctionRole(fn model.Function) model.FunctionRole {

	name := strings.ToLower(fn.Name)

	switch {

	case strings.HasPrefix(name, "new"):
		return model.RoleFactory

	case strings.Contains(name, "create"),
		strings.Contains(name, "update"),
		strings.Contains(name, "delete"),
		strings.Contains(name, "save"):
		return model.RoleRepository

	case strings.Contains(name, "get"),
		strings.Contains(name, "find"),
		strings.Contains(name, "list"):
		return model.RoleRepository

	case strings.Contains(name, "handle"),
		strings.Contains(name, "handler"):
		return model.RoleHandler

	case strings.Contains(name, "service"):
		return model.RoleService

	case strings.Contains(name, "map"),
		strings.Contains(name, "convert"),
		strings.Contains(name, "dto"):
		return model.RoleMapper

	case strings.Contains(name, "validate"),
		strings.Contains(name, "check"):
		return model.RoleValidator
	}

	return model.RoleUnknown
}
