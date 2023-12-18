package greet

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Prompter struct {
	Input  io.Reader
	Output io.Writer
}

func NewPrompter() *Prompter {
	return &Prompter{
		Input:  os.Stdin,
		Output: os.Stdout,
	}
}

func (p *Prompter) Prompt() {
	name := "Stranger"

	fmt.Fprintln(p.Output, "What is your name?")
	input := bufio.NewScanner(p.Input)

	if input.Scan() {
		name = input.Text()
	}

	fmt.Fprintf(p.Output, "Hello, %v\n", name)
}

func Main() {
	NewPrompter().Prompt()
}
