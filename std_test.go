package goimports

import (
	"fmt"
	"github.com/artarts36/gods"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestStdLibsIsActual(t *testing.T) { //nolint:gocognit // not need
	goRoot, exists := os.LookupEnv("GOROOT")
	if !exists {
		t.Skip("GOROOT not set")
	}

	goRoot += "/src"

	entries, err := os.ReadDir(goRoot)
	if err != nil {
		t.Fatalf("failed to read go src: %s", err)
	}

	srcStdLibsSet := gods.NewSet[string]()

	var walkLibs func(entries []os.DirEntry, parent, prefix string) error

	walkLibs = func(entries []os.DirEntry, parent, prefix string) error {
		for _, entry := range entries {
			if entry.Name() == "testdata" || entry.Name() == "internal" || entry.Name() == "vendor" {
				continue
			}

			if entry.IsDir() {
				pkgName := filepath.Base(entry.Name())
				if prefix != "" {
					pkgName = prefix + string(os.PathSeparator) + pkgName
				}

				srcStdLibsSet.Add(pkgName)

				path := entry.Name()
				if parent != "" {
					path = parent + string(os.PathSeparator) + entry.Name()
				}

				children, cerr := os.ReadDir(path)
				if cerr != nil {
					return fmt.Errorf("failed to read children dir %q: %w", entry.Name(), cerr)
				}

				err = walkLibs(children, path, pkgName)
				if err != nil {
					return err
				}
			}
		}

		return nil
	}

	err = walkLibs(entries, goRoot, "")
	if err != nil {
		t.Fatalf("failed to walk libs: %s", err)
	}

	if srcStdLibsSet.Equal(stdLibsSet) {
		return
	}

	content := make([]string, 0, len(stdLibsSet.List()))
	for _, val := range srcStdLibsSet.List() {
		content = append(content, fmt.Sprintf("%q,", val))
	}

	err = os.WriteFile("standard.txt", []byte(strings.Join(content, "\n")), 0755)
	if err != nil {
		t.Fatalf("failed to write standard.txt: %s", err)
	}

	t.Fatal("standard libs set is not actual")
}
