package adapter

import (
	"testing"

	"github.com/amonvix/go-doc-agent/internal/context"
	"github.com/amonvix/go-doc-agent/internal/language"
)

type fakeAdapter struct{ name string }

func (f fakeAdapter) Analyze(ctx *context.Project) error { return nil }
func (f fakeAdapter) Language() string                   { return f.name }

func TestSelect_ReturnsErrorWhenLanguageNotRegistered(t *testing.T) {
	registry = map[language.ID]Adapter{}

	_, err := Select(language.Go)
	if err == nil {
		t.Fatal("expected error for unregistered language")
	}
}

func TestRegisterAndSelect(t *testing.T) {
	registry = map[language.ID]Adapter{}
	want := fakeAdapter{name: "go-a"}
	Register(language.Go, want)

	got, err := Select(language.Go)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got.Language() != "go-a" {
		t.Fatalf("Language() = %q, want %q", got.Language(), "go-a")
	}
}

func TestRegister_OverridesExistingAdapter(t *testing.T) {
	registry = map[language.ID]Adapter{}
	Register(language.Go, fakeAdapter{name: "old"})
	Register(language.Go, fakeAdapter{name: "new"})

	got, err := Select(language.Go)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got.Language() != "new" {
		t.Fatalf("expected latest adapter to override previous one, got %q", got.Language())
	}
}
