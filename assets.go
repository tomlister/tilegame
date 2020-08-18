package main

import (
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"golang.org/x/image/font"
)

func importImage(path string) *ebiten.Image {
	importedImage, _, _ := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	return importedImage
}

func importFont(size float64) *font.Face {
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 100
	mplusNormalFont := truetype.NewFace(tt, &truetype.Options{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	return &mplusNormalFont
}

//lint:ignore U1000 Stubs
func importSound(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()
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
	world.spawnActor(kb, 0, 0)

	world.generateWorld()
}
