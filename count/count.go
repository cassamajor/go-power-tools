package count

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Counter struct {
	Input  io.Reader
	Output io.Writer
}

func NewCounter() *Counter {
	return &Counter{
		Input:  os.Stdin,
		Output: os.Stdout,
	}
}

func (c *Counter) Count() {
	lines := 0
	input := bufio.NewScanner(c.Input)
	for input.Scan() {
		lines++
	}
	fmt.Fprintln(c.Output, lines)
}

func Count() {
	NewCounter().Count()
}
