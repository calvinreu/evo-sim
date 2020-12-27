package main

//MapConfig contains all the config values from the map json file
type MapConfig struct {
	FoodGrowthSpeedPerWaterPercent, MinWaterRequirement, AverageWaterPercentage, Width, Height, ContinentCount, Seed float64 // if Seed = 0 generate random seed
}

//Field contains the information of one square of the map
type Field struct {
	Creatures             []uint64
	Food, WaterPercentage float64
}

//Map contains the data about every part of the map
type Map struct {
	Fields [][]Field //[x][y]
	Config MapConfig
}
