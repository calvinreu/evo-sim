package graphic

import (
	"fmt"

	"../config"
	"../environment"
)

//Configure graphic for use with info from config file
func (graphic *Graphic) Configure(config *config.Config, chart *environment.Map) {

	err := graphic.New(config.WinConfig.Title, config.WinConfig.X, config.WinConfig.Y, config.WinConfig.Width, config.WinConfig.Height, config.WinConfig.WindowFlags, config.WinConfig.RendererFlags, config.WinConfig.FPS)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Window initialized")

	graphic.AddMap(&config.ImgConfig[0], chart)

	fmt.Println("Map Texture initialized")

	for _, i := range config.ImgConfig[1:] {
		mainSpriteID := graphic.AddSprite(i.ImgPath, i.SrcRects[0])
		fmt.Println("Initialized sprite via file path: ", i.ImgPath)
		for _, iSrcRect := range i.SrcRects[1:] {
			graphic.AddSpriteByID(mainSpriteID, iSrcRect)
			fmt.Println("Initialized sprite via spriteID: ", mainSpriteID)
		}
	}
}

func runOutput(chart *environment.Map, window *Graphic, running *bool) {

}
