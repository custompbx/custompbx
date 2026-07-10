package web

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestRegistryEventConstantsAreRegistered(t *testing.T) {
	eventConstants := registryEventConstants(t)
	registered := map[string]bool{}
	for _, name := range coreEvents.EventNames() {
		registered[name] = true
	}

	skipped := map[string]bool{
		eventLogin: true,
	}
	var missing []string
	for constName, eventName := range eventConstants {
		if skipped[eventName] {
			continue
		}
		if !registered[eventName] {
			missing = append(missing, constName+"="+eventName)
		}
	}
	sort.Strings(missing)
	if len(missing) != 0 {
		t.Fatalf("event constants are not registered: %s", strings.Join(missing, ", "))
	}
}

func registryEventConstants(t *testing.T) map[string]string {
	t.Helper()
	_, sourceFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("cannot resolve registry test path")
	}

	files, err := filepath.Glob(filepath.Join(filepath.Dir(sourceFile), "registry*.go"))
	if err != nil {
		t.Fatal(err)
	}
	if len(files) == 0 {
		t.Fatal("no registry files found")
	}

	fset := token.NewFileSet()
	events := map[string]string{}
	for _, file := range files {
		if strings.HasSuffix(file, "_test.go") {
			continue
		}
		parsed, err := parser.ParseFile(fset, file, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
		for _, decl := range parsed.Decls {
			gen, ok := decl.(*ast.GenDecl)
			if !ok || gen.Tok != token.CONST {
				continue
			}
			for _, spec := range gen.Specs {
				valueSpec, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}
				for i, name := range valueSpec.Names {
					if !strings.HasPrefix(name.Name, "event") || i >= len(valueSpec.Values) {
						continue
					}
					lit, ok := valueSpec.Values[i].(*ast.BasicLit)
					if !ok || lit.Kind != token.STRING {
						continue
					}
					value, err := strconv.Unquote(lit.Value)
					if err != nil {
						t.Fatal(err)
					}
					events[name.Name] = value
				}
			}
		}
	}
	return events
}
