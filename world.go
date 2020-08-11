package main

import (
	"math/rand"

	"golang.org/x/image/font"

	"github.com/faiface/beep"
	"github.com/hajimehoshi/ebiten"
)

//World Stores all the things accessable by the rendering and logic pipelines
type World struct {
	Actors    []Actor
	Text      []Text //Text is ephemeral, only lasts one frame.
	Templates map[string]Actor
	CameraX   int
	CameraY   int
	VelocityX float64
	VelocityY float64
	Debug     bool
	Font      []*font.Face
	State     map[string]interface{}
	Images    map[string]*ebiten.Image
	TagTable  map[string]int
	Sounds    map[string]*beep.Streamer
}

//Text Stores text data
type Text struct {
	Text       string
	X          int
	Y          int
	Background bool
	Width      int
}

//NewWorld Creates a new world with no actors
func NewWorld() World {
	nw := World{}
	return nw
}

func (world *World) createText(textstr string, x, y, width int, bg bool) {
	newText := Text{
		Text:       textstr,
		X:          x,
		Y:          y,
		Background: bg,
		Width:      width,
	}
	(*world).Text = append((*world).Text, newText)
}

func (world *World) spawnActor(actor Actor, x, y int) {
	actorMod := actor
	actorMod.X = x
	actorMod.Y = y
	(*world).Actors = append((*world).Actors, actorMod)
}

func (world *World) spawnActorRepeat(actor Actor, x, y, repeatx, repeaty int) {
	sx, sy := actor.Image.Size()
	for yp := 0; yp < repeaty; yp++ {
		for xp := 0; xp < repeatx; xp++ {
			actorMod := actor
			actorMod.X = x + (xp * sx)
			actorMod.Y = y + (yp * sy)
			(*world).Actors = append((*world).Actors, actorMod)
		}
	}
}

func (world *World) spawnActorRepeatSizeDefined(actor Actor, x, y, sx, sy, repeatx, repeaty int) {
	for yp := 0; yp < repeaty; yp++ {
		for xp := 0; xp < repeatx; xp++ {
			actorMod := actor
			actorMod.X = x + (xp * sx)
			actorMod.Y = y + (yp * sy)
			(*world).Actors = append((*world).Actors, actorMod)
		}
	}
}

func (world *World) spawnActorRandom(actor Actor, x, y, maxx, maxy, chance int) {
	//sx, sy := actor.Image.Size()
	yes := false
	for i := 0; i < chance+1; i++ {
		if rand.Int63()&(1<<62) != 0 {
			break
		}
		if chance == i {
			yes = true
		}
	}
	if yes {
		actorMod := actor
		actorMod.X = x
		actorMod.Y = y
		(*world).Actors = append((*world).Actors, actorMod)
	}
}

func (world *World) getImage(name string) *ebiten.Image {
	if (*world).Images[name] != nil {
		return (*world).Images[name]
	} else {
		return (*world).Images["missingtexture"]
	}
}

func getActorShift() (x, y float64) {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		y++
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		y--
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		x++
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		x--
	}
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		x = x * 2
		y = y * 2
	}
	return x, y
}

func (world *World) detectCollisionPointTag(x, y int, tag string) (int, bool) {
	for i := 0; i < len((*world).Actors); i++ {
		if (*world).Actors[i].Tag == tag {
			w, h := (*world).Actors[i].Image.Size()
			if (*world).Actors[i].X < x &&
				(*world).Actors[i].X+w > x &&
				(*world).Actors[i].Y < y &&
				(*world).Actors[i].Y+h > y {
				return i, true
			}
		}
	}
	return 0, false
}
