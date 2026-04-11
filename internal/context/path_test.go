package context

import "testing"

func TestPath_StringAndConstructor(t *testing.T) {
	p := NewPath("internal/service/user.go")
	if got := p.String(); got != "internal/service/user.go" {
		t.Fatalf("String() = %q, want %q", got, "internal/service/user.go")
	}
}

func TestGroupFunctionsByFile(t *testing.T) {
	funcs := []Function{
		{Name: "A", FilePath: "a.go"},
		{Name: "B", FilePath: "a.go"},
		{Name: "C", FilePath: "b.go"},
	}

	grouped := GroupFunctionsByFile(funcs)

	if len(grouped["a.go"]) != 2 {
		t.Fatalf("expected 2 functions in a.go, got %d", len(grouped["a.go"]))
	}
	if len(grouped["b.go"]) != 1 {
		t.Fatalf("expected 1 function in b.go, got %d", len(grouped["b.go"]))
	}
}

func TestGroupFunctionsByFile_EmptyInput(t *testing.T) {
	grouped := GroupFunctionsByFile(nil)
	if len(grouped) != 0 {
		t.Fatalf("expected empty map, got %d entries", len(grouped))
	}
}
