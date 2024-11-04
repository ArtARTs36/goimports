package goimports

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewImportGroupsFromAstImportSpecs(t *testing.T) {
	expected := [][]GoImport{
		[]GoImport{
			{Package: Package{Path: "os", LastName: "os"}},
			{Package: Package{Path: "strings", LastName: "strings"}},
		},
		[]GoImport{
			{Package: Package{Path: "github.com/vendor/package1", LastName: "package1"}},
			{Package: Package{Path: "github.com/vendor/package2", LastName: "package2"}},
		},
		[]GoImport{
			{Alias: "a", Package: Package{Path: "github.com/artarts36/goimports/a", LastName: "a"}},
			{Package: Package{Path: "github.com/artarts36/goimports/b", LastName: "b"}},
		},
	}

	g := NewImportGroupsFromAstImportSpecs([]*ast.ImportSpec{
		{Name: &ast.Ident{Name: "a"}, Path: &ast.BasicLit{Value: "github.com/artarts36/goimports/a"}},

		{Path: &ast.BasicLit{Value: "os"}},

		{Path: &ast.BasicLit{Value: "github.com/artarts36/goimports/b"}},

		{Path: &ast.BasicLit{Value: "github.com/vendor/package1"}},

		{Path: &ast.BasicLit{Value: "strings"}},

		{Path: &ast.BasicLit{Value: "github.com/vendor/package2"}},
	}, "github.com/artarts36/goimports")

	assert.Equal(t, expected, g.SortedImports())
}
