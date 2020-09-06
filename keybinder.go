package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
)

func keybinderActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	if ebiten.IsKeyPressed(ebiten.KeyI) {
		if !(*actor).State["Idown"].(bool) {
			(*actor).State["Idown"] = true
			if (*world).State["pause"] == false {
				ebiten.SetCursorVisibility(true)
				(*world).State["pause"] = true
				inv := Actor{
					Tag:        "inv",
					Renderhook: true,
					Rendercode: inventoryRenderCode,
					ActorLogic: inventoryActorLogic,
					Static:     true,
					Z:          3,
					State:      make(map[string]interface{}),
					Unpausable: true,
				}
				inv.State["move"] = nil
				inv.State["mousedown"] = false
				world.spawnActor(inv, 0, 0)
			}
		}
	} else {
		(*actor).State["Idown"] = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyC) {
		if !(*actor).State["Cdown"].(bool) {
			(*actor).State["Cdown"] = true
			if (*world).State["pause"] == false {
				ebiten.SetCursorVisibility(true)
				(*world).State["pause"] = true
				crafting := Actor{
					Tag:        "crafting",
					Renderhook: true,
					Rendercode: craftingRenderCode,
					ActorLogic: craftingActorLogic,
					Static:     true,
					Z:          3,
					State:      make(map[string]interface{}),
					Unpausable: true,
				}
				crafting.State["scrolloffset"] = 0.0
				crafting.State["hoveroffset"] = 0
				crafting.State["buttondown"] = false
				world.spawnActor(crafting, 0, 0)
			}
		}
	} else {
		(*actor).State["Cdown"] = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyEscape) && (*world).State["pause"].(bool) == true {
		for i := 0; i < len((*world).Actors); i++ {
			switch (*world).Actors[i].Tag {
			case "crafting", "inv", "trade", "character":
				(*world).Actors[i].Kill = true
			}
		}
		(*world).State["pause"] = false
		sePlayer, _ := audio.NewPlayerFromBytes((*world).AudioContext, (*world.Sounds["back1"]))
		sePlayer.Play()
		ebiten.SetCursorVisibility(false)
	}
	if ebiten.IsKeyPressed(ebiten.KeyI) {
		if !(*actor).State["Idown"].(bool) {
			(*actor).State["Idown"] = true
			if (*world).State["pause"] == false {
				ebiten.SetCursorVisibility(true)
				(*world).State["pause"] = true
				inv := Actor{
					Tag:        "inv",
					Renderhook: true,
					Rendercode: inventoryRenderCode,
					ActorLogic: inventoryActorLogic,
					Static:     true,
					Z:          3,
					State:      make(map[string]interface{}),
					Unpausable: true,
				}
				world.spawnActor(inv, 0, 0)
			}
		}
	} else {
		(*actor).State["Idown"] = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		if !(*actor).State["Jdown"].(bool) {
			(*actor).State["Jdown"] = true
			speech := Actor{
				Tag:        "speech",
				Renderhook: true,
				Rendercode: tradeOfferRenderLogic,
				ActorLogic: tradeOfferActorLogic,
				Static:     true,
				Z:          3,
				State:      make(map[string]interface{}),
			}
			speech.State["interval"] = 0
			speech.State["text"] = "Would you like\nto trade?"
			speech.State["pos"] = 0
			speech.State["arrowyoffset"] = 0
			speech.State["time"] = 0
			speech.State["keydown"] = false
			world.spawnActor(speech, Width/2, Height-128)
		}
	} else {
		(*actor).State["Jdown"] = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyP) {
		if !(*actor).State["Pdown"].(bool) {
			(*actor).State["Pdown"] = true
			if (*world).State["pause"] == false {
				ebiten.SetCursorVisibility(true)
				(*world).State["pause"] = true
				character := Actor{
					Tag:        "character",
					Renderhook: true,
					Rendercode: characterRenderCode,
					ActorLogic: characterActorLogic,
					Static:     true,
					Z:          3,
					State:      make(map[string]interface{}),
					Unpausable: true,
				}
				character.State["scrolloffset"] = 0.0
				character.State["hoveroffset"] = 0
				character.State["buttondown"] = false
				world.spawnActor(character, 0, 0)
			}
		}
	} else {
		(*actor).State["Pdown"] = false
	}
}

func keybinderRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
}
