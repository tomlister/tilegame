package main

import (
	"fmt"
	"runtime"
	"time"

	"net/http"
	_ "net/http/pprof"

	"github.com/hajimehoshi/ebiten"
	"github.com/tomlister/tilegame/engine/actor"
	"github.com/tomlister/tilegame/engine/asset"
	"github.com/tomlister/tilegame/engine/pipeline"
	"github.com/tomlister/tilegame/engine/world"
	"github.com/tomlister/tilegame/shaders"
)

var Width = 640
var Height = 480

func logic(world *world.World) {
	sx, sy := world.GetActorShift()
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
				(*world).Actors[i].RunActorLogic(world, sceneDidMove)
				if (*world).Actors[i].Kill {
					(*world).Actors = append((*world).Actors[:i], (*world).Actors[i+1:]...)
					i--
				}
			}
		} else {
			(*world).Actors[i].RunActorLogic(world, sceneDidMove)
			if (*world).Actors[i].Kill {
				(*world).Actors = append((*world).Actors[:i], (*world).Actors[i+1:]...)
				i--
			}
		}
	}
	//friction
	if !world.State["pause"].(bool) {
		world.ApplyFriction()
	}
}

var windowsettings = pipeline.WindowSettings{
	Name:   "hahahahahahahahha lo",
	Width:  Width,
	Height: Height,
}

func main() {

	DEBUG := false

	world := world.New()
	world.Font = append(world.Font, asset.ImportFont(14))
	world.Font = append(world.Font, asset.ImportFont(20))
	world.Font = append(world.Font, asset.ImportFont(25))
	world.State = make(map[string]interface{})
	world.State["popup"] = false
	world.State["popuptimeout"] = 0
	world.State["pause"] = false
	world.Images = make(map[string]*ebiten.Image)
	world.Images["missingtexture"] = asset.ImportImage("assets/smissing.png")
	world.Images["grass"] = asset.ImportImage("assets/grass.png")
	world.Images["stone"] = asset.ImportImage("assets/stone.png")
	world.Images["tree0"] = asset.ImportImage("assets/tree2/tree2_00.png")
	world.Images["tree1"] = asset.ImportImage("assets/tree2/tree2_01.png")
	world.Images["tree2"] = asset.ImportImage("assets/tree2/tree2_02.png")
	world.Images["tree3"] = asset.ImportImage("assets/tree2/tree2_03.png")
	world.Images["water"] = asset.ImportImage("assets/water.png")
	world.Images["beach"] = asset.ImportImage("assets/beach_sand.png")
	world.Images["chestopen"] = asset.ImportImage("assets/chests/chestopen.png")
	world.Images["chestclosed"] = asset.ImportImage("assets/chests/chestclosed.png")
	world.Images["plusone"] = asset.ImportImage("assets/plusone.png")
	world.Images["spellarcane0"] = asset.ImportImage("assets/fx/arcane/04/Arcane_Effect_1.png")
	world.Images["spellarcane1"] = asset.ImportImage("assets/fx/arcane/04/Arcane_Effect_2.png")
	world.Images["spellarcane2"] = asset.ImportImage("assets/fx/arcane/04/Arcane_Effect_3.png")
	world.Images["spellarcane3"] = asset.ImportImage("assets/fx/arcane/04/Arcane_Effect_4.png")
	world.Images["spellarcane4"] = asset.ImportImage("assets/fx/arcane/04/Arcane_Effect_5.png")
	world.Images["spellarcane5"] = asset.ImportImage("assets/fx/arcane/04/Arcane_Effect_6.png")
	world.Images["spellarcane6"] = asset.ImportImage("assets/fx/arcane/04/Arcane_Effect_7.png")
	world.Images["popup"] = asset.ImportImage("assets/popup.png")
	world.Images["elementbar"] = asset.ImportImage("assets/elementbar.png")
	world.Images["vial"] = asset.ImportImage("assets/vial.png")
	world.Images["vialmask"] = asset.ImportImage("assets/vialmask.png")
	world.Images["hotbar"] = asset.ImportImage("assets/hotbar.png")
	world.Images["purplewand"] = asset.ImportImage("assets/wands/purplewand.png")
	world.Images["ironsword"] = asset.ImportImage("assets/swords/ironsword.png")
	world.Images["ironaxe"] = asset.ImportImage("assets/axes/ironaxe.png")
	world.Images["wooditem"] = asset.ImportImage("assets/items/wood.png")
	world.Images["titlenormal"] = asset.ImportImage("assets/title/normalmap.png")
	world.Images["titlediffuse"] = asset.ImportImage("assets/title/diffuse.png")
	world.Shaders = make(map[string]*ebiten.Shader)
	world.Shaders["title"] = asset.LoadShader(shaders.TitleShader())

	title := actor.Actor{
		Tag:        "title",
		Renderhook: true,
		Rendercode: titleRenderCode,
		ActorLogic: titleActorLogic,
		Static:     true,
		State:      make(map[string]interface{}),
		Unpausable: true,
	}

	world.SpawnActor(title, 0, 0)

	go func() {
		for {
			PrintMemUsage()
			time.Sleep(5 * time.Second)
		}
	}()
	if DEBUG {
		go http.ListenAndServe("localhost:8080", nil)
	}
	//start the engine
	pipeline.StartEngine(logic, &world, windowsettings)
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
