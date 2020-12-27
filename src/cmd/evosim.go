package main

import (
	"fmt"
	"os"
)

var config Config

func main() {

	args := os.Args[1:]
	if len(args) > 2 {
		fmt.Println("to many aunch arguments")
	}
	if len(args) == 0 {
		fmt.Println("to few args")
		return
	}
	config.Load(args[0])
}
