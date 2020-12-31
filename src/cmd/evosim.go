package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	var command string
	var rendererRunning bool = true
	creatures := list.New()

	args := os.Args[1:]
	if len(args) > 2 {
		fmt.Println("to many aunch arguments")
	}
	if len(args) < 1 {
		fmt.Println("to few arguments")
		return
	}

	fmt.Println(args[0])

	switch args[0] {
	case "--help":
		terminalHelp()
		return
	case "--load":
		loadSaveFile(args[1], &config, &chart, creatures)
		break
	case "--new":
		config.Load(args[1])
		//add loading and creating of all the other parts
		break
	default:
		fmt.Println("type --help for help")
	}

	window.Configure(&config, &chart)

	go window.RunOutput(&rendererRunning)

	for true {
		fmt.Scanln(&command)

		if command == "kill_renderer" {
			rendererRunning = false
		}

		if command == "exit" {
			break
		}
	}

	rendererRunning = false

	sdl.Quit()

}

type Commands struct {
	Commands []string
}

func terminalHelp() {
	//load commands
	var commands Commands
	data, err := ioutil.ReadFile("../../resources/cmd args.json")
	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal(data, &commands)
	if err != nil {
		fmt.Println("error:", err)
	}

	for _, i := range commands.Commands {
		fmt.Println(i)
	}
}
