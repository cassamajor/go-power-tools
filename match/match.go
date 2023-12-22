package match

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type counter struct {
	input  io.Reader
	output io.Writer
}

type option func(*counter) error

func WithInput(r io.Reader) option {
	return func(c *counter) error {
		if r == nil {
			return errors.New("nil is not a valid reader")
		}
		c.input = r
		return nil
	}
}

func WithOutput(w io.Writer) option {
	return func(c *counter) error {
		if w == nil {
			return errors.New("nil is not a valid writer")
		}
		c.output = w
		return nil
	}
}

func NewCounter(opts ...option) (*counter, error) {
	c := &counter{
		input:  os.Stdin,
		output: os.Stdout,
	}

	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *counter) Count() string {
	input := bufio.NewScanner(c.input)
	for input.Scan() {
		text := input.Text()
		if strings.Contains("hello", text) {
			fmt.Fprintln(c.output, text)
		}

	}

	b := new(strings.Builder)
	fmt.Fprintln(b, c.output)
	return b.String()
}

func DefaultCounter() {
	c, err := NewCounter()

	if err != nil {
		panic(err)
	}

	fmt.Fprintln(c.output, c.Count())
}
