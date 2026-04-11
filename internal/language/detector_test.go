package language

import "testing"

func TestDetect(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
		ok   bool
	}{
		{name: "go file", path: "main.go", want: "go", ok: true},
		{name: "python file", path: "script.py", want: "python", ok: true},
		{name: "javascript file", path: "app.js", want: "javascript", ok: true},
		{name: "typescript file", path: "app.ts", want: "javascript", ok: true},
		{name: "unsupported extension", path: "README.md", want: "", ok: false},
		{name: "without extension", path: "Dockerfile", want: "", ok: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := Detect(tt.path)
			if got != tt.want || ok != tt.ok {
				t.Fatalf("Detect(%q) = (%q, %v), want (%q, %v)", tt.path, got, ok, tt.want, tt.ok)
			}
		})
	}
}
