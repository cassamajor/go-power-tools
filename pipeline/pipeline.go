package pipeline

import (
	"io"
	"os"
	"strings"
)

type option func(session *Pipeline)

type Pipeline struct {
	Reader io.Reader
	Output io.Writer
	Error  error
}

func WithString(s string) option {
	return func(p *Pipeline) {
		p.Reader = strings.NewReader(s)
	}
}

func NewPipeline(opts ...option) *Pipeline {
	c := &Pipeline{
		Reader: os.Stdin,
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
	io.Copy(p.Output, p.Reader)
}

type option func(session *Pipeline)

func WithReader(r io.Reader) option {
	return func(p *Pipeline) {
		p.Reader = r
	}
}

func WithString(s string) option {
	return func(p *Pipeline) {
		p.Reader = strings.NewReader(s)
	}
}

func NewPipeline(opts ...option) *Pipeline {
	c := &Pipeline{
		Reader: os.Stdin,
		Output: os.Stdout,
		Error:  nil,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}
