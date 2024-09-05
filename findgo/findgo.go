package findgo

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func Files(path string) (paths []string) {
	fsys := os.DirFS(path)
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("an error encountered while walking %v: %v\n", path, err)
		}

		if filepath.Ext(p) == ".go" {
			paths = append(paths, p)
		}
		return nil
	})
	return paths
}
