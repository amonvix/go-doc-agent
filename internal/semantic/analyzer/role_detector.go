/*
this file operates some actions, when:
receive a request + service call -> handler
receive the context +return another context -> mapper
calls databes -> repository
create structs -> factory
*/
package analyzer

import (
	"strings"

	"github.com/amonvix/go-doc-agent/internal/semantic"
	model "github.com/amonvix/go-doc-agent/internal/semantic"
)

func DetectFunctionRole(fn *semantic.Function) {
	name := strings.ToLower(fn.Name)

	switch {
	case strings.HasPrefix(name, "new"):
		fn.Role = model.RoleFactory

	case strings.Contains(name, "create"),
		strings.Contains(name, "update"),
		strings.Contains(name, "delete"),
		strings.Contains(name, "save"):
		fn.Role = model.RoleRepository

	case strings.Contains(name, "get"),
		strings.Contains(name, "find"),
		strings.Contains(name, "list"):
		fn.Role = model.RoleRepository

	case strings.Contains(name, "handle"),
		strings.Contains(name, "handler"):
		fn.Role = model.RoleHandler

	case strings.Contains(name, "service"):
		fn.Role = model.RoleService

	case strings.Contains(name, "map"),
		strings.Contains(name, "convert"),
		strings.Contains(name, "dto"):
		fn.Role = model.RoleMapper

	case strings.Contains(name, "validate"),
		strings.Contains(name, "check"):
		fn.Role = model.RoleValidator

	default:
		fn.Role = model.RoleUtility
	}
}
