/*
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠟⠛⠛⠛⠉⠉⠉⠋⠛⠛⠛⠻⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠛⠉⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠉⠙⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⠋⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠏⠄⠄⠄⠄⠄⠄⠄⠂⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠹⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⠛⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠠⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠘⢻⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⠃⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⠄⢠⠄⠄⡀⠄⠄⢀⠂⠄⠄⠄⠄⠄⠄⠄⠄⠄⡁⠄⠄⢛⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠐⡈⢔⠸⣐⢕⢕⢵⢰⢱⢰⢐⢤⡡⡢⣕⢄⢢⢠⠄⠄⠄⠄⠄⠄⠙⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⡁⠂⠅⢕⠌⡎⡎⣎⢎⢮⢮⣳⡳⣝⢮⢺⢜⢕⢕⢍⢎⠪⡐⠄⠁⠄⠸⣿⣿
⣿⣿⣿⣿⣿⣿⠏⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠐⠄⠄⢅⠣⡡⡣⣣⡳⡵⣝⡮⣗⣗⡯⣗⣟⡮⡮⣳⣣⣳⢱⢱⠱⣐⠄⠂⠄⢿⣿
⣿⣿⣿⣿⣿⣿⠄⠄⠄⠄⠄⠄⠄⠂⠄⠄⠄⠄⠄⠄⢂⢈⠢⡱⡱⡝⣮⣿⣟⣿⣽⣷⣿⣯⣿⣷⣿⣿⣿⣾⣯⣗⡕⡇⡇⠄⠂⡀⢹⣿
⣿⣿⣿⣿⣿⡟⠄⠄⠄⠄⠄⠄⠂⠄⠄⠄⠄⠄⠄⠐⢀⢂⢕⢸⢨⢪⢳⡫⣟⣿⣻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡺⡮⡣⡣⠠⢂⠒⢸⣿
⣿⣿⣿⣿⣿⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠠⠐⠄⡂⠆⡇⣗⣝⢮⢾⣻⣞⣿⣿⣿⣿⣿⣿⣿⣿⢿⣽⣯⡯⣺⢸⢘⠨⠔⡅⢨⣿
⣿⣿⠋⠉⠙⠃⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠁⠄⠄⠄⡂⡪⡪⡪⡮⡮⡯⣻⣽⣾⣿⣿⣿⣟⣿⣿⣿⣽⣿⣿⡯⣯⡺⡸⡰⡱⢐⡅⣼⣿
⣿⠡⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠠⠈⠆⠱⠑⠝⠜⠕⡝⡝⣞⢯⢿⣿⣿⡿⣟⣿⣿⣿⡿⡿⣽⣷⣽⡸⡨⡪⣂⠊⣿⣿
⣿⠡⠄⡨⣢⠐⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠐⠍⡓⣗⡽⣝⠽⠍⠅⠑⠁⠉⠘⠘⠘⠵⡑⢜⢀⢀⢉⢽
⣿⠁⠠⢱⢘⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠈⠱⣁⠜⡘⠌⠄⠄⡪⣳⣟⡮⢅⠤⠠⠄⠄⣀⣀⡀⡀⠄⠈⡂⢲⡪⡠⣿
⣿⡇⠨⣺⢐⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⡀⠄⠄⠄⠤⡠⡢⢒⠦⠠⠄⠄⠄⡸⢽⣟⢮⠢⡂⡐⠄⡈⡀⠤⡀⠄⠑⢄⠨⢸⡺⣐⣿
⣿⣿⠈⠕⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⡂⡪⡐⡥⢤⣰⣰⣰⡴⡮⠢⠂⠄⠄⡊⢮⢺⢕⢵⢥⡬⣌⣒⡚⣔⢚⢌⢨⢚⠌⣾⡪⣾⣿
⣿⣿⣆⠄⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⡑⢕⢕⡯⡷⣕⢧⢓⢭⠨⡀⠄⡂⠨⡨⣪⡳⣝⢝⡽⣻⣻⣞⢽⣲⢳⢱⢡⠱⠨⣟⢺⣿⣿
⣿⣿⣿⡆⠄⡅⠇⡄⠄⠄⠄⠄⠄⠄⠄⠐⠨⢪⢹⢽⢽⣺⢝⠉⠁⠁⠄⠄⠄⢌⢎⡖⡯⡎⡗⢝⠜⣶⣯⣻⢮⡻⣟⣳⡕⠅⣷⣿⣿⣿
⣿⣿⣿⣿⣶⣶⣿⣷⠄⠄⠄⠄⠄⠄⠄⠄⠈⠔⡑⠕⠝⠄⡀⠄⠄⠊⢆⠂⠨⡪⣺⣮⣿⡾⡜⣜⡜⣄⠙⢞⣿⢿⡿⣗⢝⢸⣾⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⠄⠄⠄⠄⠄⡀⠄⠄⠄⠄⢀⠄⠠⠄⠠⠄⠄⠄⠄⠄⠄⠊⠺⡹⠳⡙⡜⡓⡭⡺⡀⠄⠣⡻⡹⡸⠨⣣⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⠄⠄⠄⠄⠄⠠⠄⠄⣂⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢄⠤⡤⡄⡆⡯⡢⡣⡣⡓⢕⠽⣄⠄⠨⡂⢌⣼⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⡆⠄⠄⠄⠄⠈⠆⠄⠸⡂⠄⠄⠄⢀⠄⢀⠈⠄⠂⠁⠙⠝⠼⠭⠣⠣⠣⠑⠌⠢⠣⡣⡠⡘⣰⣱⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⡇⠄⠄⠄⠄⠄⢑⠄⠈⡱⠄⢘⠄⡀⠨⢐⣧⣳⣷⣶⣦⣤⣴⣶⣶⣶⡶⠄⡠⡢⡕⣜⠎⡮⣣⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⡇⠄⠄⠄⠄⠄⠄⠢⠄⠨⠄⠄⠣⡀⠄⢀⢀⢙⠃⡿⢿⠿⡿⡿⢟⢋⢔⡱⣝⢜⡜⡪⡪⣵⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⡁⠄⠄⠄⠄⠄⠄⠄⠅⠄⠡⠄⠄⠡⢀⢂⠢⡡⠡⠣⡑⣏⢯⡻⡳⣹⡺⡪⢎⠎⡆⢣⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣇⠄⠄⠄⠄⠄⠄⠄⠐⠄⠄⠁⠄⢈⠄⢂⠕⡕⡝⢕⢎⢎⢮⢎⢯⢺⢸⢬⠣⢃⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠠⠨⡐⠌⢆⢇⢧⢭⣣⡳⣵⢫⣳⢱⠱⢑⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣆⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠁⡊⢌⢢⢡⢣⢪⡺⡪⡎⡎⡎⡚⣨⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣄⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠕⡅⢗⢕⡳⡭⣳⢕⠕⡱⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠌⠄⠑⠩⢈⢂⣱⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⡀⢄⠄⣀⠄⡀⣀⢠⢄⣖⣖⣞⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⣱⡐⡕⡕⡽⣝⣟⣮⣾⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣵⣽⣸⣃⣧⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
*/

//Package graphic is using the sdl2 go interface from (c)https://github.com/veandco/go-sdl2/ under the BSD 3 License
package graphic

import (
	"container/list"
	"fmt"
	"sync/atomic"

	"../config"
	"../environment"
	"github.com/veandco/go-sdl2/sdl"
)

//Graphic contains the information required to render a window with diffrent Sprites
type Graphic struct {
	running      *atomic.Value
	fps          uint32
	chartSRect   sdl.Rect
	chartDRect   sdl.Rect
	creatures    Sprite
	sprites      []Sprite // use for multiple stationary textures
	chartTexture *sdl.Texture
	renderer     *sdl.Renderer
	window       *sdl.Window
	tileW, tileH int32
}

//Configure graphic for use with info from config file
func (graphic *Graphic) Configure(config *config.Config, chart *environment.Map) {

	err := graphic.New(config.WinConfig.Title, config.WinConfig.X, config.WinConfig.Y, config.WinConfig.Width, config.WinConfig.Height, config.WinConfig.WindowFlags, config.WinConfig.RendererFlags, config.WinConfig.FPS)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Window initialized")

	graphic.tileW, graphic.tileH = config.ImgConfig[0].SrcRects[0].W, config.ImgConfig[0].SrcRects[0].H
	graphic.AddMap(&config.ImgConfig[0], chart)

	var creatureSRect sdl.Rect
	creatureSRect.X, creatureSRect.Y, creatureSRect.W, creatureSRect.H = config.ImgConfig[1].SrcRects[0].X, config.ImgConfig[1].SrcRects[0].Y, config.ImgConfig[1].SrcRects[0].W, config.ImgConfig[1].SrcRects[0].H

	graphic.creatures, err = NewSprite(graphic.renderer, config.ImgConfig[1].ImgPath, creatureSRect)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Map Texture initialized")

	for _, i := range config.ImgConfig[2:] {
		mainSpriteID := graphic.AddSprite(i.ImgPath, i.SrcRects[0])
		fmt.Println("Initialized sprite via file path: ", i.ImgPath)
		for _, iSrcRect := range i.SrcRects[1:] {
			graphic.AddSpriteByID(mainSpriteID, iSrcRect)
			fmt.Println("Initialized sprite via spriteID: ", mainSpriteID)
		}
	}
}

//AddMap adds the map for rendering to the graphic
func (graphic *Graphic) AddMap(tileConfig *config.ImageConfig, chart *environment.Map) {
	var err error

	if len(chart.Fields) == 0 {
		fmt.Println("map not initialized before function call: graphic.AddMap")
		return
	}

	MapWidth, MapHeight := (int32)(len(chart.Fields))*(graphic.tileW), (int32)(len(chart.Fields[0]))*(graphic.tileH)

	graphic.chartTexture, err = graphic.renderer.CreateTexture(sdl.PIXELFORMAT_RGB888, sdl.TEXTUREACCESS_TARGET, MapWidth, MapHeight)
	if err != nil {
		fmt.Println(err)
		return
	}

	//spriteID 0 = desert 1 = land 2 = water

	spriteID := graphic.AddSprite(tileConfig.ImgPath, tileConfig.SrcRects[2])

	if spriteID != 0 {
		fmt.Println("tile sprite has to be the first sprite but it is sprite nr: ", spriteID)
	}

	graphic.AddSpriteByID(spriteID, tileConfig.SrcRects[1])
	graphic.AddSpriteByID(spriteID, tileConfig.SrcRects[0])

	err = graphic.renderer.SetRenderTarget(graphic.chartTexture)
	if err != nil {
		fmt.Println(err)
		return
	}

	var dRect sdl.Rect

	for x := 0; x < len(chart.Fields); x++ {
		for y := 0; y < len(chart.Fields[0]); y++ {

			dRect.X = int32(x) * tileConfig.SrcRects[0].W
			dRect.Y = int32(y) * tileConfig.SrcRects[0].H
			dRect.W = tileConfig.SrcRects[0].W
			dRect.H = tileConfig.SrcRects[0].H

			if chart.Fields[x][y].WaterPercentage < 10 {
				graphic.renderer.Copy(graphic.sprites[0].texture, &graphic.sprites[0].srcRect, &dRect)
			} else if chart.Fields[x][y].WaterPercentage < 55 {
				graphic.renderer.Copy(graphic.sprites[0].texture, &graphic.sprites[1].srcRect, &dRect)
			} else {
				graphic.renderer.Copy(graphic.sprites[0].texture, &graphic.sprites[2].srcRect, &dRect)
			}
		}
	}

	graphic.sprites = nil

	err = graphic.renderer.SetRenderTarget(nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	windowWidth, windowHeight := graphic.window.GetSize()

	if windowHeight > MapHeight {
		graphic.chartDRect.H, graphic.chartSRect.H = MapHeight, MapHeight
		graphic.chartDRect.Y = (windowHeight - MapHeight) / 2
	} else {
		graphic.chartDRect.H, graphic.chartSRect.H = windowHeight, windowHeight
	}

	if windowWidth > MapWidth {
		graphic.chartDRect.W, graphic.chartSRect.W = MapWidth, MapWidth
		graphic.chartDRect.X = (windowWidth - MapWidth) / 2
	} else {
		graphic.chartDRect.W, graphic.chartSRect.W = windowWidth, windowWidth
	}
}

//AddCreature instance to the renderer the creature won't be changed
func (graphic *Graphic) AddCreature(creaturePosition CreatureInstance) *list.Element {
	e := graphic.creatures.instances.PushBack(&creaturePosition)
	if _, ok := e.Value.(*CreatureInstance); ok {
		return e
	}
	println("fatal error &instance in func AddCreature is no *CreatureInstance")
	return nil

}

//RemoveCreature Removes a graphics instance of a creature with the pointer creature
func (graphic *Graphic) RemoveCreature(creature *environment.Creature) {
	if _, ok := graphic.creatures.instances.Remove(creature.GraphicInstance).(*CreatureInstance); !ok {
		println("fatal error creature.GraphicInstance *list.Element in func RemoveCreature does not contain type CreatureInstance")
	}
}

//SetRunningBool setter for thread independent acces to the graphic.running value
func (graphic *Graphic) SetRunningBool(running *atomic.Value) {
	graphic.running = running
}

//RunOutput will render FPS frames every second until running is false
func (graphic Graphic) RunOutput(done chan bool) {
	var timeStamp, frameTime uint32
	frameTime = (1000 / graphic.fps) - 1

	fmt.Println("Graphics system running")

	for graphic.running.Load().(bool) {
		timeStamp = sdl.GetTicks()
		graphic.Render()
		if sdl.GetTicks()-timeStamp < frameTime {
			sdl.Delay(frameTime - (sdl.GetTicks() - timeStamp))
		}
	}

	fmt.Println("Graphics system paused")
	done <- true
}

//Render renders the information from the graphic object to the screen
func (graphic *Graphic) Render() {

	var dRect sdl.FRect
	dRect.W, dRect.H = (float32)(graphic.creatures.srcRect.W), (float32)(graphic.creatures.srcRect.H)

	graphic.renderer.SetDrawColor(10, 10, 10, 1)
	graphic.renderer.Clear()

	graphic.renderer.Copy(graphic.chartTexture, &graphic.chartSRect, &graphic.chartDRect)

	//render creatures
	for i := graphic.creatures.instances.Front(); i != nil; i = i.Next() {
		if instance, ok := i.Value.(*CreatureInstance); ok {
			dRect.X, dRect.Y = ((*instance).X*(float32)(graphic.tileW))+(float32)(graphic.chartDRect.X), ((*instance).Y*(float32)(graphic.tileH))+(float32)(graphic.chartDRect.Y)
			graphic.renderer.CopyF(graphic.creatures.texture, &graphic.creatures.srcRect, &dRect)
		} else {
			fmt.Println("list of sprite does not contain CreatureInstances")
		}
	}

	for _, i := range graphic.sprites {
		for j := i.instances.Front(); j != nil; j = j.Next() {
			if instance, ok := j.Value.(*Instance); ok {
				graphic.renderer.CopyExF(i.texture, &i.srcRect, &instance.DestRect, instance.Angle, &instance.Center, sdl.FLIP_HORIZONTAL)
			} else {
				fmt.Println("list of sprite does not contain Instances")
			}
		}
	}
	graphic.renderer.Present()
}

//New returns a Graphic object with initialized renderer and window note that Sprites have to be added manual
func (graphic *Graphic) New(title string, x, y, width, heigh int32, WindowFlags, RendererFlags, FPS uint32) error {
	var err = sdl.Init(sdl.INIT_VIDEO | sdl.INIT_TIMER)
	if err != nil {
		return err
	}

	graphic.window, err = sdl.CreateWindow(title, x, y, width, heigh, WindowFlags)
	if err != nil {
		sdl.QuitSubSystem(sdl.INIT_VIDEO | sdl.INIT_TIMER)
		return err
	}

	graphic.renderer, err = sdl.CreateRenderer(graphic.window, -1, RendererFlags)

	if err != nil {
		sdl.QuitSubSystem(sdl.INIT_VIDEO)
		return err
	}

	graphic.fps = FPS

	return nil
}

//AddSprite adds another sprite which can be used be creating a instance of it see Sprite.NewInstance
func (graphic *Graphic) AddSprite(imgPath string, srcRect sdl.Rect) uint32 {
	var err error
	var sprite Sprite
	retIndex := len(graphic.sprites)

	sprite, err = NewSprite(graphic.renderer, imgPath, srcRect)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	graphic.sprites = append(graphic.sprites, sprite)

	return uint32(retIndex)
}

//AddSpriteByID adds another sprite with the same texture as sprite with id spriteID
func (graphic *Graphic) AddSpriteByID(spriteID uint32, srcRect sdl.Rect) uint32 {
	var sprite Sprite

	if len(graphic.sprites)-1 < (int)(spriteID) {
		fmt.Println("sprite: ", spriteID, " does not exist")
	}

	sprite.texture = graphic.sprites[spriteID].texture
	sprite.srcRect = srcRect

	retIndex := len(graphic.sprites)

	graphic.sprites = append(graphic.sprites, sprite)

	return uint32(retIndex)
}
