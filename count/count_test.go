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

func TestPrompt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Begins counter",
			args: args{
				Input:  bytes.NewBufferString("Steven\nCassamajor"),
				Output: new(bytes.Buffer),
			},
			want: "2\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel()

			p := count.NewCounter()
			p.Input = tt.args.Input
			p.Output = tt.args.Output
			p.Count()

			if got := tt.args.Output.String(); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
