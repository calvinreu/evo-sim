package environment

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//MapSave contains map json data
type MapSave struct {
	Fields [][]FieldSave
}

//FieldSave contains Field json data
type FieldSave struct {
	Food, WaterPercentage float64
	Creatures             []uint32
}

//Field contains the information of one square of the map
type Field struct {
	Creatures             []*Creature
	Food, WaterPercentage float64
}

//Map contains the data about every part of the map
type Map struct {
	Fields [][]Field //[x][y]
}

//Load a map from file
func (m *Map) Load(filename string, creatureIDs map[uint32]*Creature) {
	var tempSaveFile MapSave

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error:", err)
	}
	err = json.Unmarshal(data, &tempSaveFile)
	if err != nil {
		fmt.Println("error:", err)
	}

	if len(tempSaveFile.Fields) == 0 {
		fmt.Println("Map json file empty")
		return
	}

	fmt.Println("Map size x/y: ", len(tempSaveFile.Fields), len(tempSaveFile.Fields[0]))

	fmt.Println("----------")

	for i := 0; i < len(tempSaveFile.Fields); i++ {
		var colunm []Field
		m.Fields = append(m.Fields, colunm)
		for _, j := range tempSaveFile.Fields[i] {
			tempField := j.CreateField(creatureIDs)
			m.Fields[i] = append(m.Fields[i], tempField)
		}
	}
}

//CreateField create a field from field save info object
func (fieldSave *FieldSave) CreateField(creatureIDs map[uint32]*Creature) Field {
	var field Field
	field.Food = fieldSave.Food
	field.WaterPercentage = fieldSave.WaterPercentage

	for _, i := range fieldSave.Creatures {
		field.Creatures = append(field.Creatures, creatureIDs[i])
	}

	fmt.Println("Field:")
	fmt.Println("Water percentage: ", field.WaterPercentage)
	fmt.Println("Food: ", field.Food)
	fmt.Println("Creatures: ", field.Creatures)
	fmt.Println("----------")

	return field
}
