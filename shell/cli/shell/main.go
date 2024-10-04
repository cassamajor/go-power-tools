package main

import (
	"bufio"
	"fmt"
	"github.com/cassamajor/shell"
	"os"
)

func main() {
	fmt.Print("> ")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		cmd, err := shell.CmdFromString(line)

		if err != nil {
			continue
		}

		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("error", err)
		}
		fmt.Printf("%s", out)
		fmt.Print("\n> ")
	}
	fmt.Println("\n Be seeing you!")
}
