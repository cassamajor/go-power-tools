package count

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type counter struct {
	Input  io.Reader
	Output io.Writer
}

func NewCounter() *counter {
	return &counter{
		Input:  os.Stdin,
		Output: os.Stdout,
	}
}

func (c *counter) Count() int {
	lines := 0
	input := bufio.NewScanner(c.Input)
	for input.Scan() {
		lines++
	}
	return lines
}

func DefaultCounter() {
	c := NewCounter()
	fmt.Fprintln(c.Output, c.Count())
}
