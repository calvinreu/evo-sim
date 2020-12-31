package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../lib/config"
	"../lib/environment"
)

//SaveFile contains json info of the main save file
type SaveFile struct {
	configFile, mapFile string
	creatures           []string
}

func loadSaveFile(filename string, config *config.Config, chart *environment.Map, creatures *list.List) map[uint32]*environment.Creature {
	//load main save file
	var tempSaveFile SaveFile
	creatureIDs := 

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal(data, &tempSaveFile)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Save Files:")
	fmt.Println("Config file: ", tempSaveFile.configFile)
	fmt.Println("Map file: ", tempSaveFile.mapFile)
	fmt.Println("Creature files: ", tempSaveFile.creatures)

	config.Load(tempSaveFile.configFile)
	chart.Load(tempSaveFile.mapFile)

	for i := 0; i < len(tempSaveFile.creatures); i++ {
		var tempCreature environment.Creature
		tempCreature.Load(tempSaveFile.creatures[i])
		if creature, ok := creatures.PushBack(&tempCreature).Value.(*Instance); ok {

		} else {
			fmt.Println("type error creatures.PushBack does not return *Creature")
		}
	}

}
