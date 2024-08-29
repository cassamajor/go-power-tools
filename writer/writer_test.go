package writer_test

import (
	"github.com/cassamajor/writer"
	"github.com/google/go-cmp/cmp"
	"os"
	"testing"
)

func Test_WriteToFile(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "write_test.txt"
	want := []byte{1, 2, 3}
	err := writer.WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestWriteToFile(t *testing.T) {
	t.Parallel()
	path := "does_not_exist/write_test.txt"
	err := writer.WriteToFile(path, []byte{})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func Test_WriteToFileAgain(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "clobber.txt"
	err := os.WriteFile(path, []byte{4, 5, 6}, 0o644)
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile(path, []byte{4, 5, 6}, 0o600)
	if err != nil {
		t.Fatal(err)
	}

	stat, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}

	perm := stat.Mode().Perm()
	if perm != 0o600 {
		t.Errorf("want file mode 0o600, got 0%o", perm)
	}

	want := []byte{1, 2, 3}

	err = writer.WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}

}

func Test_WriteToFilePermissions(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "permissions_test.txt"
	err := os.WriteFile(path, []byte{}, 0o644)
	if err != nil {
		t.Fatal(err)
	}

	err = writer.WriteToFile(path, []byte{1, 2, 3})

	if err != nil {
		t.Fatal(err)
	}

	stat, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}

	perm := stat.Mode().Perm()
	if perm != 0o600 {
		t.Errorf("want file mode 0o600, got 0%o", perm)
	}
}
