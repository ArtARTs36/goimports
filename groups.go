package goimports

import (
	"github.com/artarts36/gods"
	"strings"
)

type ImportGroups struct {
	standard *ImportGroup
	vendor   *ImportGroup
	current  *ImportGroup
	unused   *ImportGroup

	goModule string
}

func NewImportGroups(goModule string) *ImportGroups {
	return &ImportGroups{
		standard: &ImportGroup{Set: gods.NewSet[GoImport]()},
		vendor:   &ImportGroup{Set: gods.NewSet[GoImport]()},
		current:  &ImportGroup{Set: gods.NewSet[GoImport]()},
		unused:   &ImportGroup{Set: gods.NewSet[GoImport]()},
		goModule: goModule,
	}
}

func (g *ImportGroups) AddStandard(pkgPath string) {
	g.AddStandardWithAlias("", pkgPath)
}

func (g *ImportGroups) AddStandardWithAlias(alias, pkgPath string) {
	g.standard.Add(newGoImport(alias, pkgPath))
}

func (g *ImportGroups) AddVendor(pkgPath string) {
	g.AddVendorWithAlias("", pkgPath)
}

func (g *ImportGroups) AddVendorWithAlias(alias, pkgPath string) {
	g.vendor.Add(newGoImport(alias, pkgPath))
}

func (g *ImportGroups) AddCurrent(pkgPath string) {
	g.AddCurrentWithAlias("", pkgPath)
}

func (g *ImportGroups) AddCurrentWithAlias(alias, pkgPath string) {
	g.current.Add(newGoImport(alias, pkgPath))
}

func (g *ImportGroups) AddUnused(pkgPath string) {
	g.unused.Add(newGoImport("_", pkgPath))
}

func (g *ImportGroups) Add(pkgPaths ...string) {
	for _, path := range pkgPaths {
		g.AddWithAlias("", path)
	}
}

func (g *ImportGroups) AddWithAlias(alias, pkgPath string) {
	if alias == "_" { //nolint:gocritic//not need
		g.AddUnused(pkgPath)
	} else if stdLibsSet.Has(pkgPath) {
		g.AddStandardWithAlias(alias, pkgPath)
	} else if strings.HasPrefix(pkgPath, g.goModule) {
		g.AddCurrentWithAlias(alias, pkgPath)
	} else {
		g.AddVendorWithAlias(alias, pkgPath)
	}
}

func (g *ImportGroups) SortedImports() [][]GoImport {
	const groupCapacity = 4

	groups := make([][]GoImport, 0, groupCapacity)

	if g.standard.Valid() {
		groups = append(groups, g.standard.SortedList())
	}

	if g.vendor.Valid() {
		groups = append(groups, g.vendor.SortedList())
	}

	if g.current.Valid() {
		groups = append(groups, g.current.SortedList())
	}

	if g.unused.Valid() {
		groups = append(groups, g.current.SortedList())
	}

	return groups
}

func (g *ImportGroups) Valid() bool {
	return g.standard.Valid() || g.vendor.Valid() || g.current.Valid() || g.unused.Valid()
}
