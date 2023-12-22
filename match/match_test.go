package match_test

import (
	"bytes"
	"github.com/cassamajor/match"
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

			input := match.WithInput(tt.args.Input)
			output := match.WithOutput(tt.args.Output)
			c, _ := match.NewCounter(input, output)

			if got := c.Count(); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
