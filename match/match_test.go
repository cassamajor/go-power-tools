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
		want string
	}{
		{
			name: "Find the match in the provided input",
			args: args{
				Input:  bytes.NewBufferString("hello\nworld\nhello world\n"),
				Output: new(bytes.Buffer),
			},
			want: "hello\nhello world\n",
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
