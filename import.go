package goimports

import (
	"fmt"
	"strings"
)

type GoImport struct {
	Alias   string
	Package Package
}

type Package struct {
	Path     string
	LastName string
}

func newGoImport(alias, pkgPath string) GoImport {
	imp := GoImport{
		Alias: alias,
		Package: Package{
			Path: pkgPath,
		},
	}

	pathParts := strings.Split(pkgPath, "/")
	if len(pathParts) > 0 {
		imp.Package.LastName = pathParts[len(pathParts)-1]
	}

	return imp
}

func (i *GoImport) GoString() string {
	if i.Alias == "" {
		return fmt.Sprintf("%q", i.Package.Path)
	}

	return fmt.Sprintf("%s %q", i.Alias, i.Package.Path)
}
