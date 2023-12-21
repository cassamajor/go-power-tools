package hello

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type Printer struct {
	Output io.Writer
}

type option func(*Printer) error

func WithOutput(w io.Writer) option {
	return func(p *Printer) error {
		if w == nil {
			return errors.New("nil is not a valid writer")
		}
		p.Output = w
		return nil
	}
}

func NewPrinter(option ...option) (*Printer, error) {
	p := &Printer{Output: os.Stdout}

	for _, opt := range option {
		err := opt(p)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (p *Printer) Print() string {
	fmt.Fprintln(p.Output, "Hello, world")

	s := new(strings.Builder)
	fmt.Fprint(s, p.Output) // Write the content of io.Writer to the string builder.

	return s.String()
}

func DefaultPrinter() {
	p, err := NewPrinter()

	if err != nil {
		panic(err)
	}

	p.Print()
}
