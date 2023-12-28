package count_test

import (
	"bytes"
	"github.com/cassamajor/count"
	"io"
	"testing"
)

type args struct {
	Input  io.Reader
	Output *bytes.Buffer
}

func TestCounter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Count the newlines in the provided input",
			args: args{
				Input:  bytes.NewBufferString("1\n2\n3\n"),
				Output: new(bytes.Buffer),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			input := count.WithInput(tt.args.Input)
			c, _ := count.NewCounter(input)

			if got := c.Count(); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCounterWithArgs(t *testing.T) {
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
