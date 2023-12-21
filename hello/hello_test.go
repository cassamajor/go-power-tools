package hello_test

import (
	"bytes"
	"github.com/cassamajor/hello"
	"testing"
)

func TestPrintTo_PrintsHelloMessageToGivenWriter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		option hello.Printer
		want   string
	}{
		{
			name: "Print Hello Message To Given Writer",
			option: hello.Printer{
				Output: new(bytes.Buffer),
			},
			want: "Hello, world\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			output := hello.WithOutput(tt.option.Output)
			p, _ := hello.NewPrinter(output)

			if got := p.Print(); got != tt.want {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}
