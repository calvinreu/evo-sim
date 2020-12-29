package main

import (
	"fmt"
	"os"
)

func main() {

	var config Config

	args := os.Args[1:]
	if len(args) > 2 {
		fmt.Println("to many aunch arguments")
	}
	if len(args) == 0 {
		fmt.Println("to few arguments")
		return
	}
	config.Load(args[0])

	configureGraphic(&config)

	var endProgramm string

	fmt.Scanln(&endProgramm)

}
