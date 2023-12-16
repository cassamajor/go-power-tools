package greet_test

import (
	"bytes"
	"errors"
	"github.com/cassamajor/greet"
	"io"
	"testing"
	"testing/iotest"
)

func TestPrompt(t *testing.T) {
	t.Parallel()

	type args struct {
		stdin  io.Reader
		stdout io.Writer
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Prompts user for a name and renders greeting",
			args: args{
				stdin: bytes.NewBufferString("Steven"),
				//stdout: new(bytes.Buffer),
			},
			want: "What is your name?\nHello, Steven\n",
		},
		{
			name: "Prints Hello, stranger on read error",
			args: args{
				stdin: iotest.ErrReader(errors.New("bad reader")),
				//stdout: new(bytes.Buffer),
			},
			want: "What is your name?\nHello, Stranger\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// Using tt.args.stdout is not concurrency-safe.
			// When tests are run in parallel, each test is simultaneously trying to set tt.args.stdout to a different writer to print the results.

			stdout := new(bytes.Buffer)
			greet.Prompt(tt.args.stdin, stdout)
			if got := stdout.String(); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
