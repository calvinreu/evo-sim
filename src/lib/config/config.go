package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/veandco/go-sdl2/sdl"
)

//Config constains all the data
type Config struct {
	EnvConfig EnvironmentConfig
	WinConfig WindowConfig
	ImgConfig []ImageConfig
}

//EnvironmentConfig contains info about map and creature
type EnvironmentConfig struct {
	EnergyLossMultiplier, AverageWaterPercentage, MinWaterRequirement, FoodGrowthSpeedPerWaterPercent, WaterPerFood, ReproductionCost float64
	MapHeight, MapWidth, Seed, ContinentCount                                                                                         uint16
}

//WindowConfig Render info for Window
type WindowConfig struct {
	Title                           string
	X, Y, Width, Height             int32
	WindowFlags, RendererFlags, FPS uint32
}

//ImageConfig contains info about one img
type ImageConfig struct {
	ImgPath  string
	SrcRects []sdl.Rect
}

//TempConfig contains strings to indivdual config files
type TempConfig struct {
	EnvironmentConfigFile, WindowConfigFile string
	ImageConfigFiles                        []string
}

//Load sets config from file
func (config *Config) Load(filePath string) {

	fmt.Println("Main config file: ", filePath)
	fmt.Println("----------")

	//load main config file
	var tempConfig TempConfig
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal(data, &tempConfig)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Config files:")
	fmt.Println("Environment config file: ", tempConfig.EnvironmentConfigFile)
	fmt.Println("Window config file: ", tempConfig.WindowConfigFile)
	fmt.Println("Image config files: ", tempConfig.ImageConfigFiles)
	fmt.Println("----------")

	//load Environment config file
	dataEnv, err := ioutil.ReadFile(tempConfig.EnvironmentConfigFile)
	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal(dataEnv, &config.EnvConfig)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Environment config:")
	fmt.Println("EnergyLossMultiplier:", config.EnvConfig.EnergyLossMultiplier)
	fmt.Println("AverageWaterPercentage", config.EnvConfig.AverageWaterPercentage)
	fmt.Println("MinWaterRequirement", config.EnvConfig.MinWaterRequirement)
	fmt.Println("FoodGrowthSpeedPerWaterPercent", config.EnvConfig.FoodGrowthSpeedPerWaterPercent)
	fmt.Println("WaterPerFood", config.EnvConfig.WaterPerFood)
	fmt.Println("ReproductionCost", config.EnvConfig.ReproductionCost)
	fmt.Println("MapHeight", config.EnvConfig.MapHeight)
	fmt.Println("MapWidth", config.EnvConfig.MapWidth)
	fmt.Println("Seed", config.EnvConfig.Seed)
	fmt.Println("ContinentCount", config.EnvConfig.ContinentCount)
	fmt.Println("----------")

	//load Window Config file
	dataWindow, err := ioutil.ReadFile(tempConfig.WindowConfigFile)
	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal(dataWindow, &config.WinConfig)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Window config:")
	fmt.Println("Title: ", config.WinConfig.Title)
	fmt.Println("X: ", config.WinConfig.X)
	fmt.Println("Y: ", config.WinConfig.Y)
	fmt.Println("Width: ", config.WinConfig.Width)
	fmt.Println("Height: ", config.WinConfig.Height)
	fmt.Println("WindowFlags: ", config.WinConfig.WindowFlags)
	fmt.Println("RendererFlags: ", config.WinConfig.RendererFlags)
	fmt.Println("FPS: ", config.WinConfig.FPS)
	fmt.Println("----------")

	//load Image Config Files
	for _, i := range tempConfig.ImageConfigFiles {
		dataImage, err := ioutil.ReadFile(i)
		if err != nil {
			fmt.Print(err)
		}

		var tempImgConfig ImageConfig

		err = json.Unmarshal(dataImage, &tempImgConfig)
		if err != nil {
			fmt.Println("error:", err)
		}

		config.ImgConfig = append(config.ImgConfig, tempImgConfig)

		fmt.Println("Image config: ", i)
		fmt.Println("Image file: ", tempImgConfig.ImgPath)
		for _, iRect := range tempImgConfig.SrcRects {
			fmt.Println("Source Rectangle:")
			fmt.Println("X: ", iRect.X)
			fmt.Println("Y: ", iRect.Y)
			fmt.Println("W: ", iRect.W)
			fmt.Println("H: ", iRect.H)
		}
		fmt.Println("----------")
	}

}
