package test

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGeneratedHandlersCompile(t *testing.T) {
	handlers, err := filepath.Glob("handlers/*_resource.go")
	if err != nil {
		t.Fatalf("Error finding handler files: %v", err)

	}

	for _, file := range handlers {
		t.Run(file, func(t *testing.T) {
			src, err := os.ReadFile(file)
			if err != nil {
				t.Errorf("Error reading file %s: %v", file, err)
				return
			}

			fset := token.NewFileSet()
			_, err = parser.ParseFile(fset, file, src, parser.AllErrors)
			if err != nil {
				t.Errorf("Compilation-error in file %s: %v", file, err)
			}
		})
	}
}

func TestGeneratedTestCompile(t *testing.T) {
	tests, err := filepath.Glob("handlers/* _test.go")
	if err != nil {
		t.Fatalf("Error finding test files: %v", err)
		return
	}

	for _, file := range tests {
		if strings.Contains(file, "verify") {
			continue // Don't check this file yourself
		}
		t.Run(file, func(t *testing.T) {
			src, err := os.ReadFile(file)
			if err != nil {
				t.Errorf("Error reading file %s: %v", file, err)
				return
			}

			fset := token.NewFileSet()
			_, err = parser.ParseFile(fset, file, src, parser.AllErrors)
			if err != nil {
				t.Errorf("Compilation-error in file %s: %v", file, err)
			}
		})
	}
}
