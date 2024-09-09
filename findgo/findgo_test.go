package findgo_test

import (
	"archive/zip"
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

func Benchmark_Files(b *testing.B) {
	fsys := os.DirFS("testdata/tree")
	b.ResetTimer()
	for range b.N {
		_ = findgo.Files(fsys)
	}
}

func Test_ZipFiles(t *testing.T) {
	t.Run("Test Files Correctly Opens Zip Files", func(t *testing.T) {
		t.Parallel()
		fsys, err := zip.OpenReader("testdata/files.zip")

		if err != nil {
			t.Fatal(err)
		}

		want := []string{
			"tree/file.go",
			"tree/subfolder/subfolder.go",
			"tree/subfolder2/another.go",
			"tree/subfolder2/file.go",
		}

		got := findgo.Files(fsys)

		if !cmp.Equal(got, want) {
			t.Error(cmp.Diff(want, got))
		}
	})
}
