package greet

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type prompter struct {
	input  io.Reader
	output io.Writer
}

type option func(*prompter) error

func WithInput(r io.Reader) option {
	return func(p *prompter) error {
		if r == nil {
			return errors.New("nil is not a valid reader")
		}
		p.input = r
		return nil
	}
}

func WithOutput(w io.Writer) option {
	return func(p *prompter) error {
		if w == nil {
			return errors.New("nil is not a valid writer")
		}
		p.output = w
		return nil
	}
}

func NewPrompter(option ...option) (*prompter, error) {
	p := &prompter{
		input:  os.Stdin,
		output: os.Stdout,
	}

	for _, opt := range option {
		err := opt(p)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (p *prompter) Prompt() string {
	name := "Stranger"

	fmt.Fprintln(p.output, "What is your name?")
	input := bufio.NewScanner(p.input)

	if input.Scan() {
		name = input.Text()
	}

	fmt.Fprintf(p.output, "Hello, %v\n", name)

	s := new(strings.Builder)
	fmt.Fprint(s, p.output) // Write the content of io.Writer to the string builder.

	return s.String()
}

func DefaultPrompt() {
	c, err := NewPrompter()

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(c.output, c.Prompt())
}
