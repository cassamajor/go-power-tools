package greet

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/cassamajor/convert"
	"io"
	"os"
)

type Prompter struct {
	Input  io.Reader
	Output io.Writer
}

type option func(*Prompter) error

func WithInput(r io.Reader) option {
	return func(p *Prompter) error {
		if r == nil {
			return errors.New("nil is not a valid reader")
		}
		p.Input = r
		return nil
	}
}

func WithOutput(w io.Writer) option {
	return func(p *Prompter) error {
		if w == nil {
			return errors.New("nil is not a valid writer")
		}
		p.Output = w
		return nil
	}
}

func NewPrompter(option ...option) (*Prompter, error) {
	p := &Prompter{
		Input:  os.Stdin,
		Output: os.Stdout,
	}

	for _, opt := range option {
		err := opt(p)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (p *Prompter) Prompt() string {
	name := "Stranger"

	fmt.Fprintln(p.Output, "What is your name?")
	input := bufio.NewScanner(p.Input)

	if input.Scan() {
		name = input.Text()
	}

	fmt.Fprintf(p.Output, "Hello, %v\n", name)

	return convert.String(p.Output)
}

func DefaultPrompt() {
	c, err := NewPrompter()

	if err != nil {
		panic(err)
	}

	c.Prompt()
}
