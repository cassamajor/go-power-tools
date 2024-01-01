package count_test

import (
	"bytes"
	"github.com/cassamajor/count"
	"github.com/rogpeppe/go-internal/testscript"
	"io"
	"os"
	"testing"
)

type Commands map[string]func() int

func TestCounter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		args  []string
		input io.Reader
		want  int
	}{
		{
			name:  "Count the newlines from a file",
			args:  []string{"testdata/three_lines.txt"},
			input: bytes.NewBufferString(""),
			want:  3,
		},
		{
			name:  "Count the newlines from stdin when no file is provided",
			args:  []string{},
			input: bytes.NewBufferString("1\n2\n3\n"),
			want:  3,
		},
		{
			name:  "When when no file is provided, and no input is provided, there are no newlines to count",
			args:  []string{},
			input: bytes.NewBufferString(""),
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			inputArgs := count.WithInputFromArgs(tt.args)
			input := count.WithInput(tt.input)
			c, err := count.NewCounter(input, inputArgs) // inputArgs is only applied when len(args) > 0

			if err != nil {
				t.Fatal(err)
			}

			if got := c.Count(); got != tt.want {
				t.Errorf("want %v, got = %v", got, tt.want)
			}
		})
	}
}

func Test(t *testing.T) {
	t.Parallel()

	dir := testscript.Params{Dir: "testdata/scripts"}
	testscript.Run(t, dir)
}

func TestMain(m *testing.M) {
	commands := Commands{"count": count.DefaultCounter}
	status := testscript.RunMain(m, commands)
	os.Exit(status)
}
