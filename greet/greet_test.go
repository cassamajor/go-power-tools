package greet_test

import (
	"bytes"
	"errors"
	"github.com/cassamajor/greet"
	"testing"
	"testing/iotest"
)

type promptTest struct {
	name string
	args greet.Prompter
	want string
}

func TestPrompt(t *testing.T) {
	t.Parallel()

	tests := []promptTest{
		{
			name: "Prompts user for a name and renders greeting",
			args: greet.Prompter{
				Input:  bytes.NewBufferString("Steven"),
				Output: new(bytes.Buffer),
			},
			want: "What is your name?\nHello, Steven\n",
		},
		{
			name: "Prints Hello, stranger on read error",
			args: greet.Prompter{
				Input:  iotest.ErrReader(errors.New("bad reader")),
				Output: new(bytes.Buffer),
			},
			want: "What is your name?\nHello, Stranger\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel()
			// Using tt.args.stdout is not concurrency-safe.
			// When tests are run in parallel, each test is simultaneously trying to set tt.args.stdout to a different writer to print the results.

			p := greet.NewPrompter()
			p.Input = tt.args.Input
			p.Output = tt.args.Output
			p.Prompt()

			if got := p.Output.String(); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
