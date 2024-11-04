package goimports

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImportGroups_SortedImports(t *testing.T) {
	cases := []struct {
		Title string

		GoModule string

		Imports []string

		Expected [][]GoImport
	}{
		{
			Title: "test with standard, vendor, current",

			GoModule: "github.com/artarts36/goimports",

			Imports: []string{
				"github.com/artarts36/goimports/a",

				"os",

				"github.com/artarts36/goimports/b",

				"github.com/vendor/package1",

				"strings",

				"github.com/vendor/package2",
			},

			Expected: [][]GoImport{
				[]GoImport{
					{Package: Package{Path: "os", LastName: "os"}},
					{Package: Package{Path: "strings", LastName: "strings"}},
				},
				[]GoImport{
					{Package: Package{Path: "github.com/vendor/package1", LastName: "package1"}},
					{Package: Package{Path: "github.com/vendor/package2", LastName: "package2"}},
				},
				[]GoImport{
					{Package: Package{Path: "github.com/artarts36/goimports/a", LastName: "a"}},
					{Package: Package{Path: "github.com/artarts36/goimports/b", LastName: "b"}},
				},
			},
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.Title, func(t *testing.T) {
			g := NewImportGroups(tCase.GoModule)

			g.Add(tCase.Imports...)

			assert.Equal(t, tCase.Expected, g.SortedImports())
		})
	}
}
