package greet_test

import (
	"bytes"
	"errors"
	"github.com/cassamajor/greet"
	"testing"
	"testing/iotest"
)

func TestPrompt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		option greet.Prompter
		want   string
	}{
		{
			name: "Prompts user for a name and renders greeting",
			option: greet.Prompter{
				Input:  bytes.NewBufferString("Steven"),
				Output: new(bytes.Buffer),
			},
			want: "What is your name?\nHello, Steven\n",
		},
		{
			name: "Prints Hello, stranger on read error",
			option: greet.Prompter{
				Input:  iotest.ErrReader(errors.New("bad reader")),
				Output: new(bytes.Buffer),
			},
			want: "What is your name?\nHello, Stranger\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel()
			// Using tt.option.Output is not concurrency-safe.
			// When tests are run in parallel, each test is simultaneously trying to set tt.option.Output to a different writer to print the results.

			input := greet.WithInput(tt.option.Input)
			output := greet.WithOutput(tt.option.Output)
			p, _ := greet.NewPrompter(input, output)

			if got := p.Prompt(); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
