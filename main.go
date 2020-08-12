package main

import (
	"fmt"
	"runtime"
	"time"

	"net/http"
	_ "net/http/pprof"

	"github.com/hajimehoshi/ebiten"
	"github.com/tomlister/tilegame/shaders"
	_ "github.com/tomlister/tilegame/shaders"
)

var Width = 640
var Height = 480

func logic(world *World) {
	sx, sy := getActorShift()
	(*world).VelocityX += sx / 4
	(*world).VelocityY += sy / 4
	//collision pass
	for i := 0; i < len((*world).Actors); i++ {
		if !(*world).Actors[i].Static {
			if (*world).Actors[i].Collidable {
				didCollide, rx, ry := (*world).Actors[i].DetectPlayerCollision(world)
				if didCollide {
					if rx != 0 {
						(*world).VelocityX = 0
					} else if ry != 0 {
						(*world).VelocityY = 0
					}
				}
			}
		}
	}
	//set world camera offsets
	if !world.State["pause"].(bool) {
		(*world).CameraX += int((*world).VelocityX)
		(*world).CameraY += int((*world).VelocityY)
	}
	for i := 0; i < len((*world).Actors); i++ {
		sceneDidMove := false
		if sx != 0 || sy != 0 {
			sceneDidMove = true
		}
		if !(*world).Actors[i].Renderhook {
			imgwidth, imgheight := (*world).Actors[i].Image.Size()
			if (*world).Actors[i].Static || ((*world).CameraX+(*world).Actors[i].X+imgwidth > 0 && (*world).CameraX+(*world).Actors[i].X < Width) && ((*world).CameraY+(*world).Actors[i].Y < Height && (*world).CameraY+(*world).Actors[i].Y+imgheight > 0) {
				(*world).Actors[i].runActorLogic(world, sceneDidMove)
				if (*world).Actors[i].Kill {
					(*world).Actors = append((*world).Actors[:i], (*world).Actors[i+1:]...)
					i--
				}
			}
		} else {
			(*world).Actors[i].runActorLogic(world, sceneDidMove)
			if (*world).Actors[i].Kill {
				(*world).Actors = append((*world).Actors[:i], (*world).Actors[i+1:]...)
				i--
			}
		}
	}
	//friction
	if !world.State["pause"].(bool) {
		world.applyFriction()
	}
}

var windowsettings = WindowSettings{
	Name:   "hahahahahahahahha lo",
	Width:  Width,
	Height: Height,
}

func main() {

	DEBUG := false

	world := NewWorld()
	world.Font = append(world.Font, importFont(14))
	world.Font = append(world.Font, importFont(20))
	world.Font = append(world.Font, importFont(25))
	world.State = make(map[string]interface{})
	world.State["popup"] = false
	world.State["popuptimeout"] = 0
	world.State["pause"] = false
	world.Images = make(map[string]*ebiten.Image)
	world.Images["missingtexture"] = importImage("assets/smissing.png")
	world.Images["grass"] = importImage("assets/grass.png")
	world.Images["stone"] = importImage("assets/stone.png")
	world.Images["tree0"] = importImage("assets/tree2/tree2_00.png")
	world.Images["tree1"] = importImage("assets/tree2/tree2_01.png")
	world.Images["tree2"] = importImage("assets/tree2/tree2_02.png")
	world.Images["tree3"] = importImage("assets/tree2/tree2_03.png")
	world.Images["water"] = importImage("assets/water.png")
	world.Images["beach"] = importImage("assets/beach_sand.png")
	world.Images["chestopen"] = importImage("assets/chests/chestopen.png")
	world.Images["chestclosed"] = importImage("assets/chests/chestclosed.png")
	world.Images["plusone"] = importImage("assets/plusone.png")
	world.Images["spellarcane0"] = importImage("assets/fx/arcane/04/Arcane_Effect_1.png")
	world.Images["spellarcane1"] = importImage("assets/fx/arcane/04/Arcane_Effect_2.png")
	world.Images["spellarcane2"] = importImage("assets/fx/arcane/04/Arcane_Effect_3.png")
	world.Images["spellarcane3"] = importImage("assets/fx/arcane/04/Arcane_Effect_4.png")
	world.Images["spellarcane4"] = importImage("assets/fx/arcane/04/Arcane_Effect_5.png")
	world.Images["spellarcane5"] = importImage("assets/fx/arcane/04/Arcane_Effect_6.png")
	world.Images["spellarcane6"] = importImage("assets/fx/arcane/04/Arcane_Effect_7.png")
	world.Images["popup"] = importImage("assets/popup.png")
	world.Images["elementbar"] = importImage("assets/elementbar.png")
	world.Images["vial"] = importImage("assets/vial.png")
	world.Images["vialmask"] = importImage("assets/vialmask.png")
	world.Images["hotbar"] = importImage("assets/hotbar.png")
	world.Images["purplewand"] = importImage("assets/wands/purplewand.png")
	world.Images["ironsword"] = importImage("assets/swords/ironsword.png")
	world.Images["ironaxe"] = importImage("assets/axes/ironaxe.png")
	world.Images["wooditem"] = importImage("assets/items/wood.png")
	world.Images["titlenormal"] = importImage("assets/title/normalmap.png")
	world.Images["titlediffuse"] = importImage("assets/title/diffuse.png")
	world.Shaders = make(map[string]*ebiten.Shader)
	world.Shaders["title"] = loadShader(shaders.TitleShader())

	title := Actor{
		Tag:        "title",
		Renderhook: true,
		Rendercode: titleRenderCode,
		ActorLogic: titleActorLogic,
		Static:     true,
		State:      make(map[string]interface{}),
		Unpausable: true,
	}

	world.spawnActor(title, 0, 0)

	go func() {
		for {
			PrintMemUsage()
			time.Sleep(5 * time.Second)
		}
	}()
	if DEBUG {
		go http.ListenAndServe("localhost:8080", nil)
	}
	//start the bruh engine
	StartEngine(logic, &world, windowsettings)
}

/*
	From: https://gist.github.com/j33ty/79e8b736141be19687f565ea4c6f4226
*/

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
