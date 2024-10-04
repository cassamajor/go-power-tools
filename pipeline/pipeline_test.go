package pipeline_test

import (
	"bytes"
	"errors"
	"github.com/cassamajor/pipeline"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_Stdout(t *testing.T) {
	t.Run("Stdout prints message to output", func(t *testing.T) {
		t.Parallel()

		want := "Hello, world\n"
		p := pipeline.FromString(want)

		buf := new(bytes.Buffer)
		p.Output = buf
		p.Stdout()

		if p.Error != nil {
			t.Fatal(p.Error)
		}

		got := buf.String()

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
		p := pipeline.FromString(want)
		p.Error = errors.New("oh no")

		buf := new(bytes.Buffer)
		p.Output = buf
		p.Stdout()

		got := buf.String()
		if got != "" {
			t.Errorf("want no output from Stdout after error, but got %q", got)
		}
	})
}
