package goimports

import "strings"

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
