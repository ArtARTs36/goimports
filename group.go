package goimports

import (
	"slices"
	"strings"

	"github.com/artarts36/gods"
)

type ImportGroup struct {
	*gods.Set[GoImport]
}

func (g *ImportGroup) SortedList() []GoImport {
	list := g.List()
	slices.SortFunc(list, func(a, b GoImport) int {
		return strings.Compare(a.Package.Path, b.Package.Path)
	})
	return list
}
