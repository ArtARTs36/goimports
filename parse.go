package goimports

import (
	"go/ast"
	"strings"
)

func NewImportGroupsFromAstImportSpecs(specs []*ast.ImportSpec, goModule string) *ImportGroups {
	g := NewImportGroups(goModule)

	for _, spec := range specs {
		pkgPath := strings.Trim(spec.Path.Value, `"`)

		alias := ""
		if spec.Name != nil {
			alias = spec.Name.Name
		}

		g.Add(alias, pkgPath)
	}

	return g
}
