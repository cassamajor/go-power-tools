package greet_test

import (
	"bytes"
	"errors"
	"github.com/cassamajor/greet"
	"io"
	"testing"
	"testing/iotest"
)

type promptTest struct {
	name   string
	option args
	want   string
}

type args struct {
	input  io.Reader
	output *bytes.Buffer
}

func TestPrompt(t *testing.T) {
	t.Parallel()

	tests := []promptTest{
		{
			name: "Prompts user for a name and renders greeting",
			option: args{
				input:  bytes.NewBufferString("Steven"),
				output: new(bytes.Buffer),
			},
			want: "What is your name?\nHello, Steven\n",
		},
		{
			name: "Prints Hello, stranger on read error",
			option: args{
				input:  iotest.ErrReader(errors.New("bad reader")),
				output: new(bytes.Buffer),
			},
			want: "What is your name?\nHello, Stranger\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel()
			// Using tt.option.output is not concurrency-safe.
			// When tests are run in parallel, each test is simultaneously trying to set tt.option.output to a different writer to print the results.

			input := greet.WithInput(tt.option.input)
			output := greet.WithOutput(tt.option.output)
			p, _ := greet.NewPrompter(input, output)

			if got := p.Prompt(); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
