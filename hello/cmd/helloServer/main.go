package main

import (
	"fmt"
	"github.com/cassamajor/hello"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Listening on http://localhost:9001")
	http.ListenAndServe(":9001", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			p := &hello.Printer{Output: os.Stdout}
			p.Print()
		}))
}
