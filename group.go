package goimports

import (
	"github.com/artarts36/gds"
	"slices"
	"strings"
)

type ImportGroup struct {
	*gds.Map[string, GoImport]
}

func newImportGroup() *ImportGroup {
	return &ImportGroup{
		gds.NewMap[string, GoImport](),
	}
}

func (g *ImportGroup) Clone() *ImportGroup {
	return &ImportGroup{
		Map: g.Map.Clone(),
	}
}

func (g *ImportGroup) SortedList() []GoImport {
	list := g.List()
	slices.SortFunc(list, func(a, b GoImport) int {
		return strings.Compare(a.Package.Path, b.Package.Path)
	})
	return list
}

func (g *ImportGroup) Add(goImport GoImport) {
	g.Set(goImport.Package.Path, goImport)
}
