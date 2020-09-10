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
	s, err := ebiten.NewShader(data)
	if err != nil {
		log.Fatalln(err)
	}
	return s
}

func actorSetup(world *World, windowsettings WindowSettings, gs interface{}) {
	hasGS := false
	if gs != nil {
		hasGS = true
	}
	//Create actors.
	//Everything is an actor.
	player := Actor{
		Tag:        "Player",
		Image:      world.getImage("player"),
		ActorLogic: playerActorLogic,
		Shadow:     true,
		Z:          1,
		State:      make(map[string]interface{}),
	}
	//State allows for actor specific variables for use in its logic
	player.State["StrafeLeft"] = 0
	player.State["StrafeRight"] = 0
	player.State["StrafeUp"] = 0
	player.State["StrafeDown"] = 0
	player.State["Swimming"] = false
	if hasGS {
		player.State["xp"] = gs.(GameState).XP
	} else {
		player.State["xp"] = 0
	}
	player.State["mana"] = 200
	player.State["manamax"] = 200
	player.State["health"] = 100
	if hasGS {
		player.State["hotbar"] = gs.(GameState).Hotbar
	} else {
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
	}
	player.State["tooltimeout"] = 0
	if hasGS {
		player.State["inventory"] = gs.(GameState).Inventory
	} else {
		player.State["inventory"] = [9]Item{}
	}
	if hasGS {
		player.State["attributes"] = gs.(GameState).Attributes
	} else {
		player.State["attributes"] = []Attribute{
			{
				Name:   "Run Speed+",
				Amount: 1,
				Cost:   20,
			},
			{
				Name:   "Knockback+",
				Amount: 1,
				Cost:   10,
			},
		}
	}
	//find pos
	for i := 0; i < len(world.Actors); i++ {
		if world.Actors[i].Tag == "Player" {
			world.TagTable["Player"] = i
			break
		}
	}
	playerImageSizeX, playerImageSizeY := world.getImage("player").Size()
	if hasGS {
		world.CameraX = (-gs.(GameState).PlayerX + (Width / 2) - (playerImageSizeX / 2))
		world.CameraY = (-gs.(GameState).PlayerY + (Height / 2) - (playerImageSizeY / 2))
		world.spawnActor(player, gs.(GameState).PlayerX, gs.(GameState).PlayerY)
	} else {
		world.CameraX = (-(200 * 16) / 2) + (Width / 2) - (playerImageSizeX / 2)
		world.CameraY = (-(200 * 16) / 2) + (Height / 2) - (playerImageSizeY / 2)
		world.spawnActor(player, (200*16)/2, (200*16)/2)
	}

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
		{
			Item: Item{
				Name:      "Mana Potion",
				ImageName: "manapotion",
				Quantity:  1,
			},
			Needs: []Item{
				{
					Name:      "Mana Crystal",
					ImageName: "manacrystal",
					Quantity:  2,
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

	arrow := Actor{
		Tag:        "Crosshair",
		Image:      world.getImage("crosshair"),
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
	arrowImageSizeX, arrowImageSizeY := world.getImage("crosshair").Size()
	world.spawnActor(arrow, (windowsettings.Width/2)-arrowImageSizeX/2, (windowsettings.Height/2)-arrowImageSizeY/2)

	xpIcon := Actor{
		Tag:        "xpicon",
		Image:      world.Images["plusone"],
		ActorLogic: backgroundActorLogic,
		Static:     true,
		Shadow:     true,
		Z:          3,
	}
	xpIconImageSizeX, xpIconImageSizeY := world.Images["plusone"].Size()
	world.spawnActor(xpIcon, (xpIconImageSizeX / 2), windowsettings.Height-xpIconImageSizeY-(xpIconImageSizeY/2))

	xpCounter := Actor{
		Tag:        "xpcounter",
		Renderhook: true,
		Rendercode: xpcounterRenderCode,
		ActorLogic: xpcounterActorLogic,
		Static:     true,
		Z:          3,
	}
	world.spawnActor(xpCounter, xpIconImageSizeX*2, windowsettings.Height-(xpIconImageSizeY/2)-1)

	popup := Actor{
		Tag:        "popup",
		Renderhook: true,
		Rendercode: popupRenderCode,
		ActorLogic: popupActorLogic,
		Static:     true,
		Z:          3,
		State:      make(map[string]interface{}),
	}
	popup.State["Interval"] = 0
	world.spawnActor(popup, 0, windowsettings.Height)

	health := Actor{
		Renderhook: true,
		Rendercode: healthRenderCode,
		ActorLogic: healthActorLogic,
		Static:     true,
		Z:          3,
		State:      make(map[string]interface{}),
	}
	world.spawnActor(health, windowsettings.Width-148, windowsettings.Height-148)

	hotbar := Actor{
		Tag:        "hotbar",
		Renderhook: true,
		Rendercode: hotbarRenderCode,
		ActorLogic: hotbarActorLogic,
		Static:     true,
		Z:          3,
		State:      make(map[string]interface{}),
	}
	hotbar.State["Interval"] = 0
	world.spawnActor(hotbar, 32, 0)

	hand := Actor{
		Tag:               "hand",
		Renderhook:        true,
		Rendercode:        handRenderCode,
		ActorLogic:        handActorLogic,
		RenderDestination: (*world).getImage("offscreen"),
		Static:            true,
		Z:                 2,
		State:             make(map[string]interface{}),
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
	kb.State["Pdown"] = false
	world.spawnActor(kb, 0, 0)

	sePlayer, _ := audio.NewPlayerFromBytes((*world).AudioContext, (*world.Sounds["gopherland"]))
	world.State["musicplayer"] = sePlayer
	world.State["musicplayer"].(*audio.Player).SetVolume(0.25)
	world.State["musicplayer"].(*audio.Player).Play()

	world.generateWorld()
	world.generateDungeonWorld()
}
