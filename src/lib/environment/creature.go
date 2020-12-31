package environment

import (
	"math"

	"../config"
)

type Brain struct {
}

//Attributes every creature has
type Attributes struct {
	Speed, WaterSlowdown float64
}

type Creature struct {
	FoodLevel   float64
	WaterLevel  float64
	EnergyLevel float64

	attributes Attributes
	brain      Brain
}

//Generates a new Creature
func NewCreature(attributes Attributes, brain Brain) Creature {
	return Creature{0, 0, 0, attributes, brain}
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

//Getter for attributes
func (c *Creature) GetAttributes() Attributes {
	return c.attributes
}

//Load creature from file
func (c *Creature) Load(filename string) {

}
