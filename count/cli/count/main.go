package main

import (
	"fmt"
	"github.com/cassamajor/count"
	"os"
)

func main() {
	c, err := count.NewCounter(count.WithInputFromArgs(os.Args[1:]))

	if err != nil {
		panic(err)
	}

	fmt.Println(c.Count())
}
