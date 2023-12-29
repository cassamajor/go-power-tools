package count

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
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

func WithInputFromArgs(args []string) option {
	return func(c *counter) error {
		if len(args) > 0 {
			file, err := os.Open(args[0])
			if err != nil {
				return err
			}
			c.input = file
		}
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

func (c *counter) Count() int {
	lines := 0
	input := bufio.NewScanner(c.input)
	for input.Scan() {
		lines++
	}
	return lines
}

func DefaultCounter() int {
	input := WithInputFromArgs(os.Args[1:])
	c, err := NewCounter(input)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Fprintln(c.output, c.Count())
	return 0
}
