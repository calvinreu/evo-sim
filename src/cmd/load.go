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
	ConfigFile, MapFile string
	Creatures           []string
}

func loadSaveFile(filename string, config *config.Config, chart *environment.Map, creatures *list.List) {
	//load main save file
	var tempSaveFile SaveFile
	creatureIDs := make(map[uint32]*environment.Creature)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error:", err)
	}
	err = json.Unmarshal(data, &tempSaveFile)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Save Files:")
	fmt.Println("Config file: ", tempSaveFile.ConfigFile)
	fmt.Println("Map file: ", tempSaveFile.MapFile)
	fmt.Println("Creature files: ", tempSaveFile.Creatures)
	fmt.Println("----------")

	config.Load(tempSaveFile.ConfigFile)

	for i := 0; i < len(tempSaveFile.Creatures); i++ {
		var tempCreature environment.Creature
		tempCreature.Load(tempSaveFile.Creatures[i])
		if creature, ok := creatures.PushBack(&tempCreature).Value.(*environment.Creature); ok {
			creatureIDs[uint32(i)] = creature
		} else {
			fmt.Println("type error creatures.PushBack does not return *environment.Creature")
		}
		fmt.Println("-----------")
	}

	chart.Load(tempSaveFile.MapFile, creatureIDs)
}
