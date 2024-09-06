package findgo

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func Files(fsys fs.FS) (paths []string) {
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("an error encountered: %v\n", err)
		}

		if filepath.Ext(p) == ".go" {
			paths = append(paths, p)
		}
		return nil
	})
	return paths
}
