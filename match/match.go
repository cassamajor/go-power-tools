package match

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type matcher struct {
	input  io.Reader
	output io.Writer
}

type option func(*matcher) error

func WithInput(r io.Reader) option {
	return func(c *matcher) error {
		if r == nil {
			return errors.New("nil is not a valid reader")
		}
		c.input = r
		return nil
	}
}

func WithOutput(w io.Writer) option {
	return func(c *matcher) error {
		if w == nil {
			return errors.New("nil is not a valid writer")
		}
		c.output = w
		return nil
	}
}

func NewMatcher(opts ...option) (*matcher, error) {
	m := &matcher{
		input:  os.Stdin,
		output: os.Stdout,
	}

	for _, opt := range opts {
		err := opt(m)
		if err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (m *matcher) Match() string {
	input := bufio.NewScanner(m.input)
	for input.Scan() {
		text := input.Text()
		if strings.Contains("hello", text) {
			fmt.Fprintln(m.output, text)
		}

	}

	b := new(strings.Builder)
	fmt.Fprintln(b, m.output)
	return b.String()
}

func DefaultMatcher() {
	m, err := NewMatcher()

	if err != nil {
		panic(err)
	}

	fmt.Fprintln(m.output, m.Match())
}
