package environment

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"

	"github.com/veandco/go-sdl2/sdl"

	"../config"
)

//CreatureSave is a struct to convert creature json file
type CreatureSave struct {
	FoodLevel, WaterLevel, EnergyLevel, Speed, WaterSlowdown float64
	X, Y                                                     float32
	Brain                                                    string //filepath
}

//Attributes every creature has
type Attributes struct {
	Speed, WaterSlowdown float64
}

type Creature struct {
	FoodLevel   float64
	WaterLevel  float64
	EnergyLevel float64

	Position        sdl.FPoint
	GraphicInstance *list.Element

	attributes Attributes
	brain      Brain
}

//Generates a new Creature
func NewCreature(attributes Attributes, brain Brain, position sdl.FPoint, graphicInstance *list.Element) Creature {
	return Creature{0, 0, 0, position, graphicInstance, attributes, brain}
}

//Has the creature enough energy to reproduce?
func (c *Creature) CanReproduce(config *config.EnvironmentConfig) bool {
	return c.EnergyLevel >= config.ReproductionCost
}

//Is the creature able to produce more energy?
func (c *Creature) CanProduceEnergy(config *config.EnvironmentConfig) bool {
	return c.FoodLevel >= 1 && c.WaterLevel >= config.WaterPerFood
}

//When the creature has no energy left, it dies
func (c *Creature) IsDead(config *config.EnvironmentConfig) bool {
	return c.EnergyLevel <= 0 && !c.CanProduceEnergy(config)
}

//The creature tries to produce as much energy as possible. Returns how much energy the creature produced
func (c *Creature) ProduceEnergy(config *config.EnvironmentConfig) float64 {
	var producedEnergy float64 = math.Min(c.FoodLevel, math.Floor(c.WaterLevel/config.WaterPerFood))
	c.EnergyLevel += producedEnergy
	c.FoodLevel -= producedEnergy
	c.WaterLevel -= producedEnergy * config.WaterPerFood

	return producedEnergy
}

//Die removes the creature from the map usw but doen't remove the graphics instance
func (c *Creature) Die() {
	c.brain.Delete()
}

//Getter for attributes
func (c *Creature) GetAttributes() Attributes {
	return c.attributes
}

//Load creature from file
func (c *Creature) Load(filename string) {
	var creatureSaveFile CreatureSave
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error:", err)
	}
	err = json.Unmarshal(data, &creatureSaveFile)
	if err != nil {
		fmt.Println("error:", err)
	}

	c.attributes.Speed = creatureSaveFile.Speed
	c.attributes.WaterSlowdown = creatureSaveFile.WaterSlowdown
	c.FoodLevel = creatureSaveFile.FoodLevel
	c.WaterLevel = creatureSaveFile.WaterLevel
	c.EnergyLevel = creatureSaveFile.EnergyLevel
	c.Position.X = creatureSaveFile.X
	c.Position.Y = creatureSaveFile.Y

	c.brain.Load(creatureSaveFile.Brain)
}
