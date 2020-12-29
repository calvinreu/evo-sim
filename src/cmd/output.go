package main

import (
	"fmt"

	"../lib/graphic"
)

func configureGraphic(config *Config) graphic.Graphic {
	window, err := graphic.New(config.WinConfig.Title, config.WinConfig.X, config.WinConfig.Y, config.WinConfig.Width, config.WinConfig.Height, config.WinConfig.WindowFlags, config.WinConfig.RendererFlags)

	if err != nil {
		fmt.Println(err)
		return window
	}

	fmt.Println("Window initialized")

	for _, i := range config.ImgConfig {
		mainSpriteID := window.AddSprite(i.ImgPath, i.SrcRects[0])
		fmt.Println("Initialized sprite via file path: ", i.ImgPath)
		for _, iSrcRect := range i.SrcRects[1:] {
			window.AddSpriteByID(mainSpriteID, iSrcRect)
			fmt.Println("Initialized sprite via spriteID: ", mainSpriteID)
		}
	}

	return window
}
