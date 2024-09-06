package findgo_test

import (
	"github.com/cassamajor/findgo"
	"github.com/google/go-cmp/cmp"
	"os"
	"testing"
	"testing/fstest"
)

func Test_Files(t *testing.T) {
	t.Run("Test Files Correctly List Files In Tree", func(t *testing.T) {
		t.Parallel()
		fsys := os.DirFS("testdata/tree")
		want := []string{
			"file.go",
			"subfolder/subfolder.go",
			"subfolder2/another.go",
			"subfolder2/file.go",
		}

		got := findgo.Files(fsys)
		if !cmp.Equal(want, got) {
			t.Error(cmp.Diff(want, got))
		}
	})
}

func Test_FilesMap(t *testing.T) {
	t.Run("Test Files Correctly List Files In Tree", func(t *testing.T) {
		t.Parallel()
		fsys := fstest.MapFS{
			"file.go":                {},
			"subfolder/subfolder.go": {},
			"subfolder2/another.go":  {},
			"subfolder2/file.go":     {},
		}

		want := []string{
			"file.go",
			"subfolder/subfolder.go",
			"subfolder2/another.go",
			"subfolder2/file.go",
		}

		got := findgo.Files(fsys)
		if !cmp.Equal(want, got) {
			t.Error(cmp.Diff(want, got))
		}
	})
}
