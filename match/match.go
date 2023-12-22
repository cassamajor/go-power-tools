package match

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type Matcher struct {
	Input  io.Reader
	Output io.Writer
	Text   string
}

type option func(*Matcher) error

func WithInput(r io.Reader) option {
	return func(c *Matcher) error {
		if r == nil {
			return errors.New("nil is not a valid reader")
		}
		c.Input = r
		return nil
	}
}

func WithOutput(w io.Writer) option {
	return func(c *Matcher) error {
		if w == nil {
			return errors.New("nil is not a valid writer")
		}
		c.Output = w
		return nil
	}
}

func WithText(t string) option {
	return func(m *Matcher) error {
		if t == "" {
			return errors.New("text cannot be empty")
		}
		m.Text = t
		return nil
	}
}

func NewMatcher(opts ...option) (*Matcher, error) {
	m := &Matcher{
		Input:  os.Stdin,
		Output: os.Stdout,
	}

	for _, opt := range opts {
		err := opt(m)
		if err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (m *Matcher) Match() string {
	input := bufio.NewScanner(m.Input)
	for input.Scan() {
		text := input.Text()
		if strings.Contains(text, m.Text) {
			fmt.Fprintln(m.Output, text)
		}

	}

	b := new(strings.Builder)
	fmt.Fprintln(b, m.Output)
	return b.String()
}

func DefaultMatcher() {
	m, err := NewMatcher()

	if err != nil {
		panic(err)
	}

	fmt.Fprintln(m.Output, m.Match())
}
