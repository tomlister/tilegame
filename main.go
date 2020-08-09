package main

import (
	"fmt"
	"runtime"
	"time"

	"net/http"
	_ "net/http/pprof"

	"github.com/hajimehoshi/ebiten"
)

//var Width = 320
//var Height = 240

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
	(*world).CameraX += int((*world).VelocityX)
	(*world).CameraY += int((*world).VelocityY)
	for i := 0; i < len((*world).Actors); i++ {
		sceneDidMove := false
		if sx != 0 || sy != 0 {
			sceneDidMove = true
		}
		if !(*world).Actors[i].Renderhook {
			imgwidth, imgheight := (*world).Actors[i].Image.Size()
			if (*world).Actors[i].Static || ((*world).CameraX+(*world).Actors[i].X+imgwidth > 0 && (*world).CameraX+(*world).Actors[i].X < Width) && ((*world).CameraY+(*world).Actors[i].Y < Height && (*world).CameraY+(*world).Actors[i].Y+imgheight > 0) {
				(*world).Actors[i].runActorLogic(world, sceneDidMove)
				if (*world).Actors[i].Kill == true {
					(*world).Actors = append((*world).Actors[:i], (*world).Actors[i+1:]...)
					i--
				}
			}
		} else {
			(*world).Actors[i].runActorLogic(world, sceneDidMove)
			if (*world).Actors[i].Kill == true {
				(*world).Actors = append((*world).Actors[:i], (*world).Actors[i+1:]...)
				i--
			}
		}
	}
	//friction
	world.applyFriction()
}

func main() {

	DEBUG := false

	windowsettings := WindowSettings{
		Name:   "hahahahahahahahha lo",
		Width:  Width,
		Height: Height,
	}
	world := NewWorld()
	world.Font = importDefaultFont()
	world.State = make(map[string]interface{})
	world.State["popup"] = false
	world.State["popuptimeout"] = 0
	world.Images = make(map[string]*ebiten.Image)
	world.Images["missingtexture"] = importImage("missing.png")
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
	//Create actors.
	//Everything is an actor.
	playerImage := importImage("player.png")
	player := Actor{
		Tag:        "Player",
		Image:      playerImage,
		AltImages:  []*ebiten.Image{playerImage},
		ActorLogic: playerActorLogic,
		Shadow:     true,
		Z:          1,
		State:      make(map[string]interface{}),
	}
	//State allows for actor specific variables for use in it's logic
	player.State["StrafeLeft"] = 0
	player.State["StrafeRight"] = 0
	player.State["StrafeUp"] = 0
	player.State["StrafeDown"] = 0
	player.State["Swimming"] = false
	player.State["xp"] = 0
	player.State["mana"] = 200
	player.State["manamax"] = 200
	player.State["health"] = 100
	player.State["hotbarslot"] = 0
	player.State["hotbar0name"] = "Wand"
	player.State["hotbar0image"] = "purplewand"
	player.State["hotbar1name"] = "Iron Sword"
	player.State["hotbar1image"] = "ironsword"
	player.State["hotbar2name"] = "Iron Axe"
	player.State["hotbar2image"] = "ironaxe"
	player.State["tooltimeout"] = 0
	//find pos
	for i := 0; i < len(world.Actors); i++ {
		if world.Actors[i].Tag == "Player" {
			world.TagTable["Player"] = i
			break
		}
	}
	playerImageSizeX, playerImageSizeY := playerImage.Size()
	world.CameraX = (-(200 * 16) / 2) + (Width / 2) - (playerImageSizeX / 2)
	world.CameraY = (-(200 * 16) / 2) + (Height / 2) - (playerImageSizeY / 2)
	world.spawnActor(player, (200*16)/2, (200*16)/2)

	wolfImage := importImage("wolf.png")
	wolf := Actor{
		Tag:        "Wolf",
		Image:      wolfImage,
		ActorLogic: wolfActorLogic,
		Shadow:     true,
		Z:          1,
	}

	world.spawnActor(wolf, ((200*16)/2)-100, ((200*16)/2)-100)

	playerSplashImage := importImage("splash.png")
	playerSplash := Actor{
		Image:      playerSplashImage,
		AltImages:  []*ebiten.Image{playerSplashImage},
		ActorLogic: playerSplashActorLogic,
		Z:          -2,
		Disabled:   true,
		State:      make(map[string]interface{}),
	}
	playerSplashImageSizeX, playerSplashImageSizeY := playerImage.Size()
	world.spawnActor(playerSplash, (windowsettings.Width/2)-playerSplashImageSizeX/2, (windowsettings.Height/2)-playerSplashImageSizeY/2)

	arrowImage := importImage("notanarrow.png")
	arrow := Actor{
		Tag:        "Crosshair",
		Image:      arrowImage,
		ActorLogic: crossHairActorLogic,
		Static:     true,
		Shadow:     true,
		Z:          3,
	}
	//find pos
	for i := 0; i < len(world.Actors); i++ {
		if world.Actors[i].Tag == "Crosshair" {
			world.TagTable["Crosshair"] = i
			break
		}
	}
	arrowImageSizeX, arrowImageSizeY := arrowImage.Size()
	world.spawnActor(arrow, (windowsettings.Width/2)-arrowImageSizeX/2, (windowsettings.Height/2)-arrowImageSizeY/2)

	xpIcon := Actor{
		Tag:        "xpicon",
		Image:      world.Images["plusone"],
		ActorLogic: backgroundActorLogic,
		Static:     true,
		Shadow:     true,
		Z:          2,
	}
	xpIconImageSizeX, xpIconImageSizeY := world.Images["plusone"].Size()
	world.spawnActor(xpIcon, (xpIconImageSizeX / 2), windowsettings.Height-xpIconImageSizeY-(xpIconImageSizeY/2))

	xpCounter := Actor{
		Tag:        "xpcounter",
		Renderhook: true,
		Rendercode: xpcounterRenderCode,
		ActorLogic: xpcounterActorLogic,
		Static:     true,
		Z:          2,
	}
	world.spawnActor(xpCounter, xpIconImageSizeX*2, windowsettings.Height-(xpIconImageSizeY/2)-1)

	popup := Actor{
		Tag:        "popup",
		Renderhook: true,
		Rendercode: popupRenderCode,
		ActorLogic: popupActorLogic,
		Static:     true,
		Z:          2,
		State:      make(map[string]interface{}),
	}
	popup.State["Interval"] = 0
	world.spawnActor(popup, 0, windowsettings.Height)

	health := Actor{
		Renderhook: true,
		Rendercode: healthRenderCode,
		ActorLogic: healthActorLogic,
		Static:     true,
		Z:          2,
		State:      make(map[string]interface{}),
	}
	world.spawnActor(health, windowsettings.Width-148, windowsettings.Height-148)

	hotbar := Actor{
		Tag:        "hotbar",
		Renderhook: true,
		Rendercode: hotbarRenderCode,
		ActorLogic: hotbarActorLogic,
		Static:     true,
		Z:          2,
		State:      make(map[string]interface{}),
	}
	hotbar.State["Interval"] = 0
	world.spawnActor(hotbar, 32, 0)

	hand := Actor{
		Tag:        "hand",
		Renderhook: true,
		Rendercode: handRenderCode,
		ActorLogic: handActorLogic,
		Static:     true,
		Z:          3,
		State:      make(map[string]interface{}),
	}
	world.spawnActor(hand, 32, 0)

	inv := Actor{
		Tag:        "inv",
		Renderhook: true,
		Rendercode: inventoryRenderCode,
		ActorLogic: inventoryActorLogic,
		Static:     true,
		Z:          3,
		State:      make(map[string]interface{}),
	}
	world.spawnActor(inv, 0, 0)

	world.generateWorld()
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
