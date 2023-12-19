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
		want int
	}{
		{
			name: "Begins counter",
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

			c := count.NewCounter()
			c.Input = tt.args.Input
			c.Output = tt.args.Output

			if got := c.Count(); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
