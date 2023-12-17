package greet

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Prompt(stdin io.Reader, stdout io.Writer) {
	name := "Stranger"

	fmt.Fprintln(stdout, "What is your name?")
	input := bufio.NewScanner(stdin)

	if input.Scan() {
		name = input.Text()
	}

	fmt.Fprintf(stdout, "Hello, %v\n", name)
}

func Main() {
	Prompt(os.Stdin, os.Stdout)
}
