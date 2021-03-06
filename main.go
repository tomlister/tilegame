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
	sx, sy := 0.0, 0.0
	if !(*world).State["pause"].(bool) {
		sx, sy = world.getActorShift()
		(*world).VelocityX += sx / 4
		(*world).VelocityY += sy / 4
	}
	//collision pass
	for i := 0; i < len((*world).Actors); i++ {
		if !(*world).Actors[i].Static {
			if (*world).Actors[i].Collidable {
				didCollide, rx, ry := (*world).Actors[i].DetectPlayerCollision(world)
				if didCollide {
					if rx != 0 {
						(*world).VelocityX = 0
					}
					if ry != 0 {
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
	Name:   "GOexplore!",
	Width:  Width,
	Height: Height,
}

func main() {

	DEBUG := false

	world := NewWorld()
	world.Seed = 69
	world.Debug = DEBUG
	world.Font = append(world.Font, importFont(14))
	world.Font = append(world.Font, importFont(20))
	world.Font = append(world.Font, importFont(25))
	world.Font = append(world.Font, importFont(12))
	world.Font = append(world.Font, importFont(10))
	world.State = make(map[string]interface{})
	world.State["popup"] = false
	world.State["popuptimeout"] = 0
	world.State["pause"] = true
	world.Images = make(map[string]*ebiten.Image)
	world.Images["missingtexture"] = importImage("assets/missing.png")
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
	world.Images["woodensword"] = importImage("assets/swords/woodensword.png")
	world.Images["ironpowderitem"] = importImage("assets/items/ironpowder.png")
	world.Images["wateredgeS"] = importImage("assets/wateredgeS.png")
	world.Images["wateredgeE"] = importImage("assets/wateredgeE.png")
	world.Images["wateredgeW"] = importImage("assets/wateredgeW.png")
	world.Images["wateredgeN"] = importImage("assets/wateredgeN.png")
	world.Images["wateredgeSE"] = importImage("assets/wateredgeSE.png")
	world.Images["wateredgeNE"] = importImage("assets/wateredgeNE.png")
	world.Images["wateredgeSW"] = importImage("assets/wateredgeSW.png")
	world.Images["wateredgeNW"] = importImage("assets/wateredgeNW.png")
	world.Images["rock"] = importImage("assets/rock.png")
	world.Images["speech"] = importImage("assets/speech.png")
	world.Images["trader"] = importImage("assets/trader.png")
	world.Images["arrow"] = importImage("assets/arrow.png")
	world.Images["choice"] = importImage("assets/choice.png")
	world.Images["woodenaxe"] = importImage("assets/axes/woodenaxe.png")
	world.Images["manacrystal"] = importImage("assets/items/manacrystal.png")
	world.Images["cavefloor"] = importImage("assets/cave/cavefloor.png")
	world.Images["cavewall"] = importImage("assets/cave/cavewall.png")
	world.Images["cavewallS"] = importImage("assets/cave/cavewallS.png")
	world.Images["cavewallW"] = importImage("assets/cave/cavewallW.png")
	world.Images["cavewallN"] = importImage("assets/cave/cavewallN.png")
	world.Images["cavewallE"] = importImage("assets/cave/cavewallE.png")
	world.Images["cavewallEFullCorner"] = importImage("assets/cave/cavewallEFullCorner.png")
	world.Images["caveblack"] = importImage("assets/cave/caveblack.png")
	world.Images["button"] = importImage("assets/button.png")
	world.Images["buttonhover"] = importImage("assets/buttonhover.png")
	world.Images["overworld"] = importImage("assets/overworld.png")
	world.Images["cavemask"] = importImage("assets/cavemask.png")
	world.Images["hole"] = importImage("assets/hole.png")
	world.Images["enemy1"] = importImage("assets/enemy1.png")
	world.Images["minusone"] = importImage("assets/minusone.png")
	world.Images["minusten"] = importImage("assets/minusten.png")
	world.Images["manapotion"] = importImage("assets/items/manapotion.png")
	world.Images["minustwo"] = importImage("assets/minustwo.png")
	world.Images["player"] = importImage("assets/player.png")
	world.Images["crosshair"] = importImage("assets/notanarrow.png")
	world.Images["enemy2"] = importImage("assets/enemy2.png")
	world.Sounds = make(map[string]*[]byte)
	world.Sounds["hover"] = importSound(world.AudioContext, "assets/hover.wav")
	world.Sounds["select1"] = importSound(world.AudioContext, "assets/select1.wav")
	world.Sounds["back1"] = importSound(world.AudioContext, "assets/back1.wav")
	world.Sounds["text"] = importSound(world.AudioContext, "assets/text.wav")
	world.Sounds["back2"] = importSound(world.AudioContext, "assets/back2.wav")
	world.Sounds["select2"] = importSound(world.AudioContext, "assets/select2.wav")
	world.Sounds["sword1"] = importSound(world.AudioContext, "assets/swordsounds/swish-1.wav")
	world.Sounds["sword2"] = importSound(world.AudioContext, "assets/swordsounds/swish-2.wav")
	world.Sounds["sword3"] = importSound(world.AudioContext, "assets/swordsounds/swish-3.wav")
	world.Sounds["sword4"] = importSound(world.AudioContext, "assets/swordsounds/swish-4.wav")
	world.Sounds["hit1"] = importSound(world.AudioContext, "assets/swordsounds/hit1.wav")
	world.Sounds["hit2"] = importSound(world.AudioContext, "assets/swordsounds/hit2.wav")
	world.Sounds["hit3"] = importSound(world.AudioContext, "assets/swordsounds/hit3.wav")
	world.Sounds["hit4"] = importSound(world.AudioContext, "assets/swordsounds/hit4.wav")
	world.Sounds["gopherland"] = importSound(world.AudioContext, "assets/gopherland.wav")
	world.Sounds["cave"] = importSound(world.AudioContext, "assets/cave.wav")
	world.Shaders = make(map[string]*ebiten.Shader)
	world.Shaders["title"] = loadShader(shaders.TitleShader())
	world.Shaders["lighting"] = loadShader(shaders.LightShader())
	world.TagTable = make(map[string]int)
	world.State["musicplayer"] = nil
	offscreen, _ := ebiten.NewImage(Width, Height, ebiten.FilterDefault)
	world.Images["offscreen"] = offscreen

	title := Actor{
		Tag:        "title",
		Renderhook: true,
		Rendercode: titleRenderCode,
		ActorLogic: titleActorLogic,
		Static:     true,
		State:      make(map[string]interface{}),
		Unpausable: true,
	}
	title.State["player"] = nil
	title.State["playing"] = false
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
