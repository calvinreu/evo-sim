//Package graphic is using the sdl2 go interface from (c)https://github.com/veandco/go-sdl2/ under the BSD 3 License
package graphic

import (
	"container/list"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

//Instance position Angle and the Center of an instance of a sprite
type Instance struct {
	DestRect sdl.FRect
	Angle    float64
	Center   sdl.FPoint
}

//CreatureInstance is type used by renderer for every instance of a creature
type CreatureInstance *sdl.FPoint

//Sprite contains the texture a list of instances and a srcRect
type Sprite struct {
	texture   *sdl.Texture
	instances *list.List
	srcRect   sdl.Rect
}

//NewSprite creates a sprite based on a renderer, the image path and a src rectAngle
func NewSprite(renderer *sdl.Renderer, imgPath string, srcRect sdl.Rect) (Sprite, error) {
	var sprite Sprite
	var err error
	sprite.texture, err = img.LoadTexture(renderer, imgPath)

	if err != nil {
		return sprite, err
	}

	sprite.srcRect = srcRect
	sprite.instances = list.New()

	return sprite, err
}

//NewInstance adds a instance to the sprite and initializes the width and height of the dest rectAngle with the src rectAngle
func (sprite *Sprite) NewInstance(Angle float64, Center *sdl.FPoint) *Instance {
	var instance Instance

	instance.DestRect.W = (float32)(sprite.srcRect.W)
	instance.DestRect.H = (float32)(sprite.srcRect.H)
	instance.NewPosition(Center)
	instance.Angle = Angle

	if instance, ok := sprite.instances.PushBack(&instance).Value.(*Instance); ok {
		return instance
	}

	println("fatal error &instance in func NewInstace is no *Instance")
	return nil
}

//NewPosition sets the position of this instance Center is the Center of the instances new position
func (instance *Instance) NewPosition(Center *sdl.FPoint) {
	instance.Center.X = Center.X
	instance.Center.Y = Center.Y
	instance.DestRect.X = Center.X - (instance.DestRect.W / 2)
	instance.DestRect.Y = Center.Y - (instance.DestRect.H / 2)
}

//ChangePosition moves the instance by x, y
func (instance *Instance) ChangePosition(x, y float32) {
	instance.Center.X += x
	instance.Center.Y += y
	instance.DestRect.X = instance.Center.X - (instance.DestRect.W / 2)
	instance.DestRect.Y = instance.Center.Y - (instance.DestRect.H / 2)
}
