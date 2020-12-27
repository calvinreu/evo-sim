package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//Config constains all the data
type Config struct {
	EnergyLossMultiplier, AverageWaterPercentage, MinWaterRequirement, FoodGrowthSpeedPerWaterPercent, WaterPerFood float64
	MapHeight, MapWidth, Seed, ContinentCount, ReproductionCost                                                     uint16
}

//TempConfig is the struct to contain json info
type TempConfig struct {
	EnergyLossMultiplier, AverageWaterPercentage, MinWaterRequirement, FoodGrowthSpeedPerWaterPercent, WaterPerFood, MapHeight, MapWidth, Seed, ContinentCount, ReproductionCost float64
}

//Load sets config from file
func (config *Config) Load(filePath string) {
	var tempConfig TempConfig
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal(data, &tempConfig)
	if err != nil {
		fmt.Println("error:", err)
	}

	config.EnergyLossMultiplier = tempConfig.EnergyLossMultiplier
	config.AverageWaterPercentage = tempConfig.AverageWaterPercentage
	config.MinWaterRequirement = tempConfig.MinWaterRequirement
	config.FoodGrowthSpeedPerWaterPercent = tempConfig.FoodGrowthSpeedPerWaterPercent
	config.WaterPerFood = tempConfig.WaterPerFood
	config.MapHeight = (uint16)(tempConfig.MapHeight)
	config.MapWidth = (uint16)(tempConfig.MapWidth)
	config.Seed = (uint16)(tempConfig.Seed)
	config.ContinentCount = (uint16)(tempConfig.ContinentCount)
	config.ReproductionCost = (uint16)(tempConfig.ReproductionCost)

	fmt.Println("Config:")
	fmt.Println("EnergyLossMultiplier:", config.EnergyLossMultiplier)
	fmt.Println("AverageWaterPercentage", config.AverageWaterPercentage)
	fmt.Println("MinWaterRequirement", config.MinWaterRequirement)
	fmt.Println("FoodGrowthSpeedPerWaterPercent", config.FoodGrowthSpeedPerWaterPercent)
	fmt.Println("WaterPerFood", config.WaterPerFood)
	fmt.Println("MapHeight", config.MapHeight)
	fmt.Println("MapWidth", config.MapWidth)
	fmt.Println("Seed", config.Seed)
	fmt.Println("ContinentCount", config.ContinentCount)
	fmt.Println("ReproductionCost", config.ReproductionCost)

}
