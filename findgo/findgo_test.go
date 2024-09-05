package findgo_test

import (
	"github.com/cassamajor/findgo"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_Files(t *testing.T) {
	t.Run("Test Files Correctly List Files In Tree", func(t *testing.T) {
		t.Parallel()
		want := []string{
			"file.go",
			"subfolder/subfolder.go",
			"subfolder2/another.go",
			"subfolder2/file.go",
		}

		got := findgo.Files("testdata/tree")
		if !cmp.Equal(want, got) {
			t.Error(cmp.Diff(want, got))
		}
	})
}
