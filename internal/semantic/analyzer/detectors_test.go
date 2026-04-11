package analyzer

import (
	"testing"

	"github.com/amonvix/go-doc-agent/internal/semantic"
)

func TestDetectDependencies(t *testing.T) {
	fn := &semantic.Function{Calls: []semantic.Call{
		{Package: "database/sql"},
		{Package: "net/http"},
		{Package: "fmt"},
	}}

	DetectDependencies(fn)

	if len(fn.Dependencies) != 2 {
		t.Fatalf("expected 2 dependencies, got %d", len(fn.Dependencies))
	}
	if fn.Dependencies[0].Type != semantic.DependencyDatabase {
		t.Fatalf("first dependency type = %s, want database", fn.Dependencies[0].Type)
	}
	if fn.Dependencies[1].Type != semantic.DependencyNetwork {
		t.Fatalf("second dependency type = %s, want network", fn.Dependencies[1].Type)
	}
}

func TestDetectEntrypoint(t *testing.T) {
	tests := []struct {
		name string
		fn   semantic.Function
		want bool
	}{
		{name: "exported function", fn: semantic.Function{Name: "HandleRequest"}, want: true},
		{name: "unexported function", fn: semantic.Function{Name: "handleRequest"}, want: false},
		{name: "method", fn: semantic.Function{Name: "Handle", IsMethod: true}, want: false},
		{name: "empty name", fn: semantic.Function{Name: ""}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := tt.fn
			DetectEntrypoint(&fn)
			if fn.IsEntryPoint != tt.want {
				t.Fatalf("IsEntryPoint = %v, want %v", fn.IsEntryPoint, tt.want)
			}
		})
	}
}

func TestDetectFunctionRole(t *testing.T) {
	tests := []struct {
		name string
		fn   semantic.Function
		want semantic.FunctionRole
	}{
		{"factory", semantic.Function{Name: "NewService"}, semantic.RoleFactory},
		{"repository create", semantic.Function{Name: "CreateUser"}, semantic.RoleRepository},
		{"repository read", semantic.Function{Name: "FindUser"}, semantic.RoleRepository},
		{"handler", semantic.Function{Name: "HandleHTTP"}, semantic.RoleHandler},
		{"service", semantic.Function{Name: "UserService"}, semantic.RoleService},
		{"mapper", semantic.Function{Name: "MapDTO"}, semantic.RoleMapper},
		{"validator", semantic.Function{Name: "ValidateInput"}, semantic.RoleValidator},
		{"utility default", semantic.Function{Name: "ComputeHash"}, semantic.RoleUtility},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := tt.fn
			DetectFunctionRole(&fn)
			if fn.Role != tt.want {
				t.Fatalf("Role = %s, want %s", fn.Role, tt.want)
			}
		})
	}
}

func TestDetectSideEffects(t *testing.T) {
	t.Run("pure when no dependencies", func(t *testing.T) {
		fn := &semantic.Function{}
		DetectSideEffects(fn)
		if !fn.IsPure {
			t.Fatal("expected function to be pure")
		}
		if len(fn.SideEffects) != 0 {
			t.Fatalf("expected no side effects, got %d", len(fn.SideEffects))
		}
	})

	t.Run("database and network side effects", func(t *testing.T) {
		fn := &semantic.Function{Dependencies: []semantic.Dependency{
			{Name: "database/sql", Type: semantic.DependencyDatabase},
			{Name: "net/http", Type: semantic.DependencyNetwork},
			{Name: "os", Type: semantic.DependencyFile},
		}}

		DetectSideEffects(fn)

		if fn.IsPure {
			t.Fatal("expected function to be impure")
		}
		if len(fn.SideEffects) != 2 {
			t.Fatalf("expected 2 side effects, got %d", len(fn.SideEffects))
		}
	})
}

func TestDetectFunctionLayer(t *testing.T) {
	tests := []struct {
		name string
		fn   semantic.Function
		want semantic.Layer
	}{
		{
			name: "infrastructure by dependency",
			fn: semantic.Function{Dependencies: []semantic.Dependency{{Type: semantic.DependencyNetwork}}},
			want: semantic.LayerInfrastructure,
		},
		{name: "interface by handler role", fn: semantic.Function{Role: semantic.RoleHandler}, want: semantic.LayerInterface},
		{name: "application by service role", fn: semantic.Function{Role: semantic.RoleService}, want: semantic.LayerApplication},
		{name: "application by validator role", fn: semantic.Function{Role: semantic.RoleValidator}, want: semantic.LayerApplication},
		{name: "domain by purity", fn: semantic.Function{IsPure: true}, want: semantic.LayerDomain},
		{name: "unknown fallback", fn: semantic.Function{Role: semantic.RoleUtility, IsPure: false}, want: semantic.LayerUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := tt.fn
			DetectFunctionLayer(&fn)
			if fn.Layer != tt.want {
				t.Fatalf("Layer = %s, want %s", fn.Layer, tt.want)
			}
		})
	}
}
