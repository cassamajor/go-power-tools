package pipeline_test

import (
	"bytes"
	"errors"
	"github.com/cassamajor/convert"
	"github.com/cassamajor/pipeline"
	"github.com/google/go-cmp/cmp"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_Stdout(t *testing.T) {
	t.Run("Stdout prints message to output", func(t *testing.T) {
		t.Parallel()

		want := "Hello, world\n"

		input := pipeline.WithInput(strings.NewReader(want))
		output := pipeline.WithOutput(new(bytes.Buffer))

		p := pipeline.NewPipeline(input, output)
		p.Stdout()

		if p.Error != nil {
			t.Fatal(p.Error)
		}

		got := convert.String(p.Output)

		if !cmp.Equal(want, got) {
			t.Error(cmp.Diff(want, got))
			//t.Errorf("want %q, got %q", want, got)
		}
	})
}

func Test_StdoutErrorSafe(t *testing.T) {
	t.Run("Stdout prints nothing on Error", func(t *testing.T) {
		t.Parallel()

		want := "Hello, world\n"

		input := pipeline.WithInput(strings.NewReader(want))
		output := pipeline.WithOutput(new(bytes.Buffer))

		p := pipeline.NewPipeline(input, output)
		p.Error = errors.New("oh no")
		p.Stdout()

		got := convert.String(p.Output)

		if got != "" {
			t.Errorf("want no output from Stdout after error, but got %q", got)
		}
	})
}

func Test_FromFile(t *testing.T) {
	t.Run("Reads all data from file", func(t *testing.T) {
		t.Parallel()

		path := t.TempDir() + "hello.txt"
		perm := os.FileMode(0o600)
		want := []byte("Hello, world")

		err := os.WriteFile(path, want, perm)

		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}

		p := pipeline.FromFile(path)

		if p.Error != nil {
			t.Fatal(p.Error)
		}

		got, err := io.ReadAll(p.Input)

		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(want, got) {
			t.Error(cmp.Diff(want, got))
		}
	})
}

func Test_FromFileErrorSafe(t *testing.T) {
	t.Run("Sets Error Given Non-existent file", func(t *testing.T) {
		t.Parallel()

		p := pipeline.FromFile("doesnt-exist.txt")

		if p.Error == nil {
			t.Fatal("want error opening non-existent file, got nil")
		}

	})
}
