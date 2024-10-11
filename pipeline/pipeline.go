package pipeline

import (
	"errors"
	"fmt"
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
	return WithInput(strings.NewReader(s))
}

func WithInput(r io.Reader) option {
	return func(p *Pipeline) {
		if r == nil {
			p.Error = errors.New("nil is not a valid reader")
			return
		}
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

func FromFile(path string) *Pipeline {
	p := NewPipeline()

	content, err := os.Open(path)

	if err != nil {
		p.Error = fmt.Errorf("%w", err)
		return p
	}

	p.Input = content

	return p
}

func (p *Pipeline) Stdout() {
	if p.Error != nil {
		return
	}
	io.Copy(p.Output, p.Input)
}
