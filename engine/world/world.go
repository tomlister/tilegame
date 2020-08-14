package engine

import (
	"math/rand"

	"golang.org/x/image/font"

	"github.com/faiface/beep"
	"github.com/hajimehoshi/ebiten"
	"github.com/tomlister/tilegame/engine/actor"
)

//World Stores all the things accessable by the rendering and logic pipelines
type World struct {
	Actors    []actor.Actor
	Text      []Text //Text is ephemeral, only lasts one frame.
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
	Shaders   map[string]*ebiten.Shader
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
func New() World {
	nw := World{}
	return nw
}

func (w *World) CreateText(textstr string, x, y, width int, bg bool) {
	newText := Text{
		Text:       textstr,
		X:          x,
		Y:          y,
		Background: bg,
		Width:      width,
	}
	(*w).Text = append((*w).Text, newText)
}

func (w *World) SpawnActor(a actor.Actor, x, y int) {
	actorMod := a
	actorMod.X = x
	actorMod.Y = y
	(*w).Actors = append((*w).Actors, actorMod)
}

//lint:ignore U1000 Engine function
func (w *World) SpawnActorRepeat(a actor.Actor, x, y, repeatx, repeaty int) {
	sx, sy := a.Image.Size()
	for yp := 0; yp < repeaty; yp++ {
		for xp := 0; xp < repeatx; xp++ {
			actorMod := a
			actorMod.X = x + (xp * sx)
			actorMod.Y = y + (yp * sy)
			(*w).Actors = append((*w).Actors, actorMod)
		}
	}
}

func (w *World) SpawnActorRepeatSizeDefined(a actor.Actor, x, y, sx, sy, repeatx, repeaty int) {
	for yp := 0; yp < repeaty; yp++ {
		for xp := 0; xp < repeatx; xp++ {
			actorMod := a
			actorMod.X = x + (xp * sx)
			actorMod.Y = y + (yp * sy)
			(*w).Actors = append((*w).Actors, actorMod)
		}
	}
}

func (w *World) SpawnActorRandom(a actor.Actor, x, y, maxx, maxy, chance int) {
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
		actorMod := a
		actorMod.X = x
		actorMod.Y = y
		(*w).Actors = append((*w).Actors, actorMod)
	}
}

func (w *World) GetImage(name string) *ebiten.Image {
	if (*w).Images[name] != nil {
		return (*w).Images[name]
	} else {
		return (*w).Images["missingtexture"]
	}
}

func (w *World) GetActorShift() (x, y float64) {
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

func (w *World) DetectCollisionPointTag(x, y int, tag string) (int, bool) {
	for i := 0; i < len((*w).Actors); i++ {
		if (*w).Actors[i].Tag == tag {
			width, height := (*w).Actors[i].Image.Size()
			if (*w).Actors[i].X < x &&
				(*w).Actors[i].X+width > x &&
				(*w).Actors[i].Y < y &&
				(*w).Actors[i].Y+height > y {
				return i, true
			}
		}
	}
	return 0, false
}
