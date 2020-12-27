package main

//Config constains all the data
type Config struct {
	EnergyLossMultiplier, AverageWaterPercentage, MinWaterRequirement, FoodGrowthSpeedPerWaterPercent, WaterPerFood float64
	MapHeight, MapWidth, Seed, ContinentCount                                                                       uint16
}

func LoadConfig
