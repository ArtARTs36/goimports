package goimports

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGoImport(t *testing.T) {
	goimp := newGoImport("ds", "github.com/artarts36/gds")

	assert.Equal(t, GoImport{
		Alias: "ds",
		Package: Package{
			Path:     "github.com/artarts36/gds",
			LastName: "gds",
		},
	}, goimp)
}
