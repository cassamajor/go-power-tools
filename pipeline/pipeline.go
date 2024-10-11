package pipeline

import (
	"errors"
	"io"
	"os"
	"strings"
)

type option func(session *Pipeline)

type Pipeline struct {
	Input  io.Reader
	Output io.Writer
	Error  error
}

func WithString(s string) option {
	return func(p *Pipeline) {
		p.Input = strings.NewReader(s)
	}
}

func WithInput(r io.Reader) option {
	return func(p *Pipeline) {
		p.Input = r
	}
}

func WithOutput(w io.Writer) option {
	return func(p *Pipeline) {
		if w == nil {
			p.Error = errors.New("nil is not a valid writer")
			return
		}
		p.Output = w
	}
}

func NewPipeline(opts ...option) *Pipeline {
	c := &Pipeline{
		Input:  os.Stdin,
		Output: os.Stdout,
		Error:  nil,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func FromString(s string) *Pipeline {
	input := WithString(s)
	return NewPipeline(input)
}

func (p *Pipeline) Stdout() {
	if p.Error != nil {
		return
	}
	io.Copy(p.Output, p.Input)
}
