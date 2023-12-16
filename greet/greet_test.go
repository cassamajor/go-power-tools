package greet_test

import (
	"bytes"
	"errors"
	"github.com/cassamajor/greet"
	"io"
	"testing"
	"testing/iotest"
)

type promptArgs struct {
	stdin  io.Reader
	stdout *bytes.Buffer
}

type promptTest struct {
	name string
	args promptArgs
	want string
}

func TestPrompt(t *testing.T) {
	t.Parallel()

	tests := []promptTest{
		{
			name: "Prompts user for a name and renders greeting",
			args: promptArgs{
				stdin:  bytes.NewBufferString("Steven"),
				stdout: new(bytes.Buffer),
			},
			want: "What is your name?\nHello, Steven\n",
		},
		{
			name: "Prints Hello, stranger on read error",
			args: promptArgs{
				stdin:  iotest.ErrReader(errors.New("bad reader")),
				stdout: new(bytes.Buffer),
			},
			want: "What is your name?\nHello, Stranger\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// Using tt.args.stdout is not concurrency-safe.
			// When tests are run in parallel, each test is simultaneously trying to set tt.args.stdout to a different writer to print the results.

			//stdout := new(bytes.Buffer)
			greet.Prompt(tt.args.stdin, tt.args.stdout)
			if got := tt.args.stdout.String(); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
