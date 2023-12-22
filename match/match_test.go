package match_test

import (
	"bytes"
	"github.com/cassamajor/match"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args match.Matcher
		want string
	}{
		{
			name: "Find a single match in the provided input",
			args: match.Matcher{
				Input:  bytes.NewBufferString("hello\n"),
				Output: new(bytes.Buffer),
				Text:   "hello",
			},
			want: "hello\n",
		},
		{
			name: "Find multiple matches in the provided input",
			args: match.Matcher{
				Input:  bytes.NewBufferString("hello\nworld\nhello world\n"),
				Output: new(bytes.Buffer),
				Text:   "hello",
			},
			want: "hello\nhello world\n",
		},
		{
			name: "No match in the provided input",
			args: match.Matcher{
				Input:  bytes.NewBufferString("that's crazy"),
				Output: new(bytes.Buffer),
				Text:   "hello",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel()

			input := match.WithInput(tt.args.Input)
			output := match.WithOutput(tt.args.Output)
			text := match.WithText(tt.args.Text)
			m, _ := match.NewMatcher(input, output, text)

			if got := m.Match(); got != tt.want {
				t.Errorf("got =\n %v, want =\n %v", got, tt.want)
			}
		})
	}
}
