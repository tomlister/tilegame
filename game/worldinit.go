package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/tomlister/tilegame/engine/actor"
	"github.com/tomlister/tilegame/engine/asset"
	"github.com/tomlister/tilegame/engine/pipeline"
	"github.com/tomlister/tilegame/engine/world"
)

func actorSetup(world *world.World, windowsettings pipeline.WindowSettings) {
	//Create actors.
	//Everything is an actor.
	playerImage := asset.ImportImage("assets/player.png")
	player := actor.Actor{
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
	player.State["inventory"] = []Item{
		Item{
			Name:      "Wood",
			ImageName: "wooditem",
			Quantity:  10,
		},
		Item{
			Name:      "Wood",
			ImageName: "wooditem",
			Quantity:  10,
		},
		Item{
			Name:      "Wood",
			ImageName: "wooditem",
			Quantity:  10,
		},
		Item{
			Name:      "Wood",
			ImageName: "wooditem",
			Quantity:  10,
		},
		Item{
			Name:      "Wood",
			ImageName: "wooditem",
			Quantity:  10,
		},
		Item{
			Name:      "Wood",
			ImageName: "wooditem",
			Quantity:  10,
		},
		Item{
			Name:      "Wood",
			ImageName: "wooditem",
			Quantity:  10,
		},
		Item{
			Name:      "Wood",
			ImageName: "wooditem",
			Quantity:  10,
		},
		Item{
			Name:      "Wood",
			ImageName: "wooditem",
			Quantity:  10,
		},
		Item{
			Name:      "Wood",
			ImageName: "wooditem",
			Quantity:  10,
		},
		Item{
			Name:      "Wood",
			ImageName: "wooditem",
			Quantity:  10,
		},
		Item{
			Name:      "Wood",
			ImageName: "wooditem",
			Quantity:  10,
		},
		Item{
			Name:      "Wood",
			ImageName: "wooditem",
			Quantity:  10,
		},
	}
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
	world.SpawnActor(player, (200*16)/2, (200*16)/2)

	wolfImage := asset.ImportImage("assets/wolf.png")
	wolf := actor.Actor{
		Tag:        "Wolf",
		Image:      wolfImage,
		ActorLogic: wolfActorLogic,
		Shadow:     true,
		Z:          1,
	}

	world.SpawnActor(wolf, ((200*16)/2)-100, ((200*16)/2)-100)

	playerSplashImage := asset.ImportImage("assets/splash.png")
	playerSplash := actor.Actor{
		Image:      playerSplashImage,
		AltImages:  []*ebiten.Image{playerSplashImage},
		ActorLogic: playerSplashActorLogic,
		Z:          -2,
		Disabled:   true,
		State:      make(map[string]interface{}),
	}
	playerSplashImageSizeX, playerSplashImageSizeY := playerImage.Size()
	world.SpawnActor(playerSplash, (windowsettings.Width/2)-playerSplashImageSizeX/2, (windowsettings.Height/2)-playerSplashImageSizeY/2)

	arrowImage := asset.ImportImage("assets/notanarrow.png")
	arrow := actor.Actor{
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
	world.SpawnActor(arrow, (windowsettings.Width/2)-arrowImageSizeX/2, (windowsettings.Height/2)-arrowImageSizeY/2)

	xpIcon := actor.Actor{
		Tag:        "xpicon",
		Image:      world.Images["plusone"],
		ActorLogic: backgroundActorLogic,
		Static:     true,
		Shadow:     true,
		Z:          2,
	}
	xpIconImageSizeX, xpIconImageSizeY := world.Images["plusone"].Size()
	world.SpawnActor(xpIcon, (xpIconImageSizeX / 2), windowsettings.Height-xpIconImageSizeY-(xpIconImageSizeY/2))

	xpCounter := actor.Actor{
		Tag:        "xpcounter",
		Renderhook: true,
		Rendercode: xpcounterRenderCode,
		ActorLogic: xpcounterActorLogic,
		Static:     true,
		Z:          2,
	}
	world.SpawnActor(xpCounter, xpIconImageSizeX*2, windowsettings.Height-(xpIconImageSizeY/2)-1)

	popup := actor.Actor{
		Tag:        "popup",
		Renderhook: true,
		Rendercode: popupRenderCode,
		ActorLogic: popupActorLogic,
		Static:     true,
		Z:          2,
		State:      make(map[string]interface{}),
	}
	popup.State["Interval"] = 0
	world.SpawnActor(popup, 0, windowsettings.Height)

	health := actor.Actor{
		Renderhook: true,
		Rendercode: healthRenderCode,
		ActorLogic: healthActorLogic,
		Static:     true,
		Z:          2,
		State:      make(map[string]interface{}),
	}
	world.SpawnActor(health, windowsettings.Width-148, windowsettings.Height-148)

	hotbar := actor.Actor{
		Tag:        "hotbar",
		Renderhook: true,
		Rendercode: hotbarRenderCode,
		ActorLogic: hotbarActorLogic,
		Static:     true,
		Z:          2,
		State:      make(map[string]interface{}),
	}
	hotbar.State["Interval"] = 0
	world.SpawnActor(hotbar, 32, 0)

	hand := actor.Actor{
		Tag:        "hand",
		Renderhook: true,
		Rendercode: handRenderCode,
		ActorLogic: handActorLogic,
		Static:     true,
		Z:          3,
		State:      make(map[string]interface{}),
	}
	world.SpawnActor(hand, 32, 0)

	kb := actor.Actor{
		Tag:        "kb",
		Renderhook: true,
		Rendercode: keybinderRenderCode,
		ActorLogic: keybinderActorLogic,
		Static:     true,
		State:      make(map[string]interface{}),
		Unpausable: true,
	}
	kb.State["Idown"] = false
	world.SpawnActor(kb, 0, 0)

	generateWorld(world)
}
