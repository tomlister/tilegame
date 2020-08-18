package main

import (
	"fmt"
	"runtime"
	"time"

	"net/http"
	_ "net/http/pprof"

	"github.com/hajimehoshi/ebiten"
	"github.com/tomlister/tilegame/assets"
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
	world.Images["missingtexture"] = importImage(assets.MISSING_go)
	world.Images["grass"] = importImage(assets.GRASS_go)
	world.Images["stone"] = importImage(assets.STONE_go)
	world.Images["tree0"] = importImage(assets.TREE2_00_go)
	world.Images["tree1"] = importImage(assets.TREE2_01_go)
	world.Images["tree2"] = importImage(assets.TREE2_02_go)
	world.Images["tree3"] = importImage(assets.TREE2_03_go)
	world.Images["water"] = importImage(assets.WATER_go)
	world.Images["beach"] = importImage(assets.BEACH_SAND_go)
	world.Images["chestopen"] = importImage(assets.CHESTOPEN_go)
	world.Images["chestclosed"] = importImage(assets.CHESTCLOSED_go)
	world.Images["plusone"] = importImage(assets.PLUSONE_go)
	world.Images["spellarcane0"] = importImage(assets.ARCANE_EFFECT_1_go)
	world.Images["spellarcane1"] = importImage(assets.ARCANE_EFFECT_2_go)
	world.Images["spellarcane2"] = importImage(assets.ARCANE_EFFECT_3_go)
	world.Images["spellarcane3"] = importImage(assets.ARCANE_EFFECT_4_go)
	world.Images["spellarcane4"] = importImage(assets.ARCANE_EFFECT_5_go)
	world.Images["spellarcane5"] = importImage(assets.ARCANE_EFFECT_6_go)
	world.Images["spellarcane6"] = importImage(assets.ARCANE_EFFECT_7_go)
	world.Images["popup"] = importImage(assets.POPUP_go)
	world.Images["elementbar"] = importImage(assets.ELEMENTBAR_go)
	world.Images["vial"] = importImage(assets.VIAL_go)
	world.Images["vialmask"] = importImage(assets.VIALMASK_go)
	world.Images["hotbar"] = importImage(assets.HOTBAR_go)
	world.Images["purplewand"] = importImage(assets.PURPLEWAND_go)
	world.Images["ironsword"] = importImage(assets.IRONSWORD_go)
	world.Images["ironaxe"] = importImage(assets.IRONAXE_go)
	world.Images["wooditem"] = importImage(assets.WOOD_go)
	world.Images["titlenormal"] = importImage(assets.NORMALMAP_go)
	world.Images["titlediffuse"] = importImage(assets.DIFFUSE_go)
	world.Images["woodensword"] = importImage(assets.WOODENSWORD_go)
	world.Images["ironpowderitem"] = importImage(assets.IRONPOWDER_go)
	world.Images["wateredgeS"] = importImage(assets.WATEREDGES_go)
	world.Images["wateredgeE"] = importImage(assets.WATEREDGEE_go)
	world.Images["wateredgeW"] = importImage(assets.WATEREDGEW_go)
	world.Images["wateredgeN"] = importImage(assets.WATEREDGEN_go)
	world.Images["wateredgeSE"] = importImage(assets.WATEREDGESE_go)
	world.Images["wateredgeNE"] = importImage(assets.WATEREDGENE_go)
	world.Images["wateredgeSW"] = importImage(assets.WATEREDGESW_go)
	world.Images["wateredgeNW"] = importImage(assets.WATEREDGENW_go)
	world.Images["rock"] = importImage(assets.ROCK_go)
	world.Shaders = make(map[string]*ebiten.Shader)
	world.Shaders["title"] = loadShader(shaders.TitleShader())
	world.Shaders["blur"] = loadShader(shaders.BlurShader())

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
