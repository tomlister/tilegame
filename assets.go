package main

import (
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"golang.org/x/image/font"
)

func importImage(path string) *ebiten.Image {
	importedImage, _, _ := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	return importedImage
}

func importFont(size float64) *font.Face {
	tt, err := truetype.Parse(fonts.ArcadeN_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 100
	rfont := truetype.NewFace(tt, &truetype.Options{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	return &rfont
}

func importSound(ctx *audio.Context, path string) *[]byte {
	f, err := ebitenutil.OpenFile(path)
	if err != nil {
		log.Fatal(err)
	}
	stream, err := wav.Decode(ctx, f)
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(stream)
	return &b
}

func loadShader(data []byte) *ebiten.Shader {
	s, _ := ebiten.NewShader(data)
	return s
}

func actorSetup(world *World, windowsettings WindowSettings) {
	//Create actors.
	//Everything is an actor.
	playerImage := importImage("assets/player.png")
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
	player.State["hotbar"] = Hotbar{
		Slot: 0,
		Slots: [3]Item{
			{
				Name:      "Wooden Axe",
				ImageName: "woodenaxe",
				Quantity:  1,
			},
		},
	}
	player.State["tooltimeout"] = 0
	player.State["inventory"] = [9]Item{
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
	world.spawnActor(player, (200*16)/2, (200*16)/2)

	(*world).State["craftable"] = []Craftable{
		{
			Item: Item{
				Name:      "Wooden Sword",
				ImageName: "woodensword",
				Quantity:  1,
			},
			Needs: []Item{
				{
					Name:      "Wood",
					ImageName: "wooditem",
					Quantity:  2,
				},
			},
			Quantity: 1,
		},
		{
			Item: Item{
				Name:      "Wooden Axe",
				ImageName: "woodenaxe",
				Quantity:  1,
			},
			Needs: []Item{
				{
					Name:      "Wood",
					ImageName: "wooditem",
					Quantity:  3,
				},
			},
			Quantity: 1,
		},
		{
			Item: Item{
				Name:      "Iron Sword",
				ImageName: "ironsword",
				Quantity:  1,
			},
			Needs: []Item{
				{
					Name:      "Wood",
					ImageName: "wooditem",
					Quantity:  2,
				},
				{
					Name:      "Iron Powder",
					ImageName: "ironpowderitem",
					Quantity:  3,
				},
			},
			Quantity: 1,
		},
	}

	(*world).State["tradable"] = []Tradable{
		{
			Item: Item{
				Name:      "Wand",
				ImageName: "purplewand",
				Quantity:  1,
			},
			Needs: []Item{
				{
					Name:      "Mana Crystal",
					ImageName: "manacrystal",
					Quantity:  20,
				},
			},
			Quantity: 1,
		},
	}

	wolfImage := importImage("assets/wolf.png")
	wolf := Actor{
		Tag:        "Wolf",
		Image:      wolfImage,
		ActorLogic: wolfActorLogic,
		Shadow:     true,
		Z:          1,
	}

	world.spawnActor(wolf, ((200*16)/2)-100, ((200*16)/2)-100)

	playerSplashImage := importImage("assets/splash.png")
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

	arrowImage := importImage("assets/notanarrow.png")
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

	kb := Actor{
		Tag:        "kb",
		Renderhook: true,
		Rendercode: keybinderRenderCode,
		ActorLogic: keybinderActorLogic,
		Static:     true,
		State:      make(map[string]interface{}),
		Unpausable: true,
	}
	kb.State["Idown"] = false
	kb.State["Cdown"] = false
	kb.State["Jdown"] = false
	world.spawnActor(kb, 0, 0)

	world.generateWorld()
	world.generateDungeonWorld()
}
