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
