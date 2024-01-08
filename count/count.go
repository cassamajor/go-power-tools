package count

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type counter struct {
	files  []io.Reader
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
		argLength := len(args)

		if argLength < 1 {
			return nil
		}

		c.files = make([]io.Reader, argLength)
		for index, file := range args {
			read, err := os.Open(file)
			if err != nil {
				return err
			}
			c.files[index] = read
		}
		c.input = io.MultiReader(c.files...)
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

func (c *counter) CountLines() int {
	lines := 0
	input := bufio.NewScanner(c.input)
	for input.Scan() {
		lines++
	}

	for _, file := range c.files {
		file.(io.Closer).Close()
	}

	return lines
}

func (c *counter) CountWords() int {
	words := 0
	input := bufio.NewScanner(c.input)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		words++
	}

	for _, file := range c.files {
		file.(io.Closer).Close()
	}

	return words
}

func LineCounter() int {
	input := WithInputFromArgs(os.Args[1:])
	c, err := NewCounter(input)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Fprintln(c.output, c.CountLines())
	return 0
}

func WordCounter() int {
	input := WithInputFromArgs(os.Args[1:])
	c, err := NewCounter(input)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Fprintln(c.output, c.CountWords())
	return 0
}
