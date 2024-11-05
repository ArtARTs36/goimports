package goimports

import (
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
		standard: newImportGroup(),
		vendor:   newImportGroup(),
		current:  newImportGroup(),
		unused:   newImportGroup(),
		goModule: goModule,
	}
}

func (g *ImportGroups) Clone() *ImportGroups {
	return &ImportGroups{
		standard: g.standard.Clone(),
		vendor:   g.vendor.Clone(),
		current:  g.current.Clone(),
		unused:   g.unused.Clone(),
		goModule: g.goModule,
	}
}

func (g *ImportGroups) AddStandard(alias, pkgPath string) {
	g.standard.Add(newGoImport(alias, pkgPath))
}

func (g *ImportGroups) AddVendor(alias, pkgPath string) {
	g.vendor.Add(newGoImport(alias, pkgPath))
}

func (g *ImportGroups) AddCurrent(alias, pkgPath string) {
	g.current.Add(newGoImport(alias, pkgPath))
}

func (g *ImportGroups) AddUnused(pkgPath string) {
	g.unused.Add(newGoImport("_", pkgPath))
}

func (g *ImportGroups) AddPkgPaths(pkgPaths ...string) {
	for _, path := range pkgPaths {
		g.Add("", path)
	}
}

func (g *ImportGroups) KeepPkgPaths(pkgPaths []string) {
	if !g.standard.IsEmpty() {
		g.standard.Keep(pkgPaths...)
	}

	if !g.vendor.IsEmpty() {
		g.vendor.Keep(pkgPaths...)
	}

	if !g.current.IsEmpty() {
		g.current.Keep(pkgPaths...)
	}

	if !g.unused.IsEmpty() {
		g.unused.Keep(pkgPaths...)
	}
}

func (g *ImportGroups) Add(alias, pkgPath string) {
	if alias == "_" { //nolint:gocritic//not need
		g.AddUnused(pkgPath)
	} else if stdLibsSet.Has(pkgPath) {
		g.AddStandard(alias, pkgPath)
	} else if strings.HasPrefix(pkgPath, g.goModule) {
		g.AddCurrent(alias, pkgPath)
	} else {
		g.AddVendor(alias, pkgPath)
	}
}

func (g *ImportGroups) SortedImports() [][]GoImport {
	const groupCapacity = 4

	groups := make([][]GoImport, 0, groupCapacity)

	if !g.standard.IsEmpty() {
		groups = append(groups, g.standard.SortedList())
	}

	if !g.vendor.IsEmpty() {
		groups = append(groups, g.vendor.SortedList())
	}

	if !g.current.IsEmpty() {
		groups = append(groups, g.current.SortedList())
	}

	if !g.unused.IsEmpty() {
		groups = append(groups, g.current.SortedList())
	}

	return groups
}

func (g *ImportGroups) IsEmpty() bool {
	return g.standard.IsEmpty() && g.vendor.IsEmpty() && g.current.IsEmpty() && g.unused.IsEmpty()
}

func (g *ImportGroups) Len() int {
	return g.standard.Len() + g.vendor.Len() + g.current.Len() + g.unused.Len()
}
