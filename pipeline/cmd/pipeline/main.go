package main

import "github.com/cassamajor/pipeline"

func main() {
	pipeline.FromString("hello, world\n").Stdout()
}
