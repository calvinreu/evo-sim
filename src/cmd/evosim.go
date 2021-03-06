package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync/atomic"

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
	creatures := list.New()
	rendererRunning := true
	var renderSync atomic.Value

	renderChannel := make(chan bool)

	renderSync.Store(rendererRunning)

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

	for i := creatures.Front(); i != nil; i = i.Next() {
		if creature, ok := i.Value.(*environment.Creature); ok {
			window.AddCreature(&creature.Position)
		} else {
			fmt.Println("list of sprite does not contain creatures")
		}
	}

	window.SetRunningBool(&renderSync)

	go window.RunOutput(renderChannel)

	for true {
		fmt.Scanln(&command)

		if command == "pause_renderer" {
			rendererRunning = false
			renderSync.Store(rendererRunning)
		}

		if command == "exit" {
			break
		}
	}

	rendererRunning = false
	renderSync.Store(rendererRunning)

	<-renderChannel
	sdl.Quit()

}

//Commands leave me alone i don't want warnings
type Commands struct {
	Commands  []string
	StartArgs []string
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

	fmt.Println("start arguments")

	for _, i := range commands.StartArgs {
		fmt.Println(i)
	}

	fmt.Println("commands")

	for _, i := range commands.Commands {
		fmt.Println(i)
	}
}
