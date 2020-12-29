package main

import (
	"fmt"
	"os"

	"../lib/config"
	"../lib/environment"
	"../lib/graphic"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {

	var config config.Config
	var chart environment.Map
	var window graphic.Graphic

	args := os.Args[1:]
	if len(args) > 2 {
		fmt.Println("to many aunch arguments")
	}
	if len(args) == 0 {
		fmt.Println("to few arguments")
		return
	}
	config.Load(args[0])

	window.Configure(&config, &chart)

	var endProgramm string

	fmt.Scanln(&endProgramm)

	sdl.Quit()

}
