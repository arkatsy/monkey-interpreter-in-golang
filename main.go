package main

import (
	"fmt"
	"monkey-go/repl"
	"os"
)

func main() {
	fmt.Println("monkey-go REPL")
	repl.Start(os.Stdin, os.Stdout)
}
