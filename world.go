package main

import (
	"math/rand"

	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
)

//World Stores all the things accessible by the rendering and logic pipelines
type World struct {
	Actors       []Actor
	Text         []Text //Text is ephemeral, only lasts one frame.
	Templates    map[string]Actor
	CameraX      int
	CameraY      int
	VelocityX    float64
	VelocityY    float64
	Debug        bool
	Font         []*font.Face
	State        map[string]interface{}
	Images       map[string]*ebiten.Image
	TagTable     map[string]int
	Sounds       map[string]*[]byte
	Shaders      map[string]*ebiten.Shader
	AudioContext *audio.Context
	Seed         int
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
	audioContext, _ := audio.NewContext(44100)
	nw.AudioContext = audioContext
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

//lint:ignore U1000 Engine function
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

func (w *World) getActorShift() (x, y float64) {
	i := (*w).TagTable["Player"]
	vactor := (*w).Actors[i]
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		vactor.Y++
		if !vactor.DetectCollision(w) {
			vactor.Y--
			y++
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		vactor.Y--
		if !vactor.DetectCollision(w) {
			vactor.Y++
			y--
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		vactor.X++
		if !vactor.DetectCollision(w) {
			vactor.X--
			x++
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		vactor.X--
		if !vactor.DetectCollision(w) {
			vactor.X++
			x--
		}
	}
	speed, _ := getAttribute(&vactor, "Run Speed+")
	if speed.Amount != 0 {
		x *= 1 + float64(speed.Amount/4)
		y *= 1 + float64(speed.Amount/4)
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

type Rect struct {
	x, y, w, h int
}

func detectPointRect(x, y int, rect Rect) bool {
	if rect.x < x &&
		rect.x+rect.w > x &&
		rect.y < y &&
		rect.y+rect.h > y {
		return true
	} else {
		return false
	}
}

func (world *World) killAll() {
	(*world).Actors = make([]Actor, 0)
}
