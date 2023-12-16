package greet

import (
	"bufio"
	"fmt"
	"io"
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
