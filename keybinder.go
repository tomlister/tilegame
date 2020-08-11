package main

import (
	"github.com/hajimehoshi/ebiten"
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
				world.spawnActor(inv, 0, 0)
			} else {
				ebiten.SetCursorVisibility(false)
				for i := 0; i < len((*world).Actors); i++ {
					if (*world).Actors[i].Tag == "inv" {
						(*world).Actors[i].Kill = true
					}
				}
				(*world).State["pause"] = false
			}
		}
	} else {
		(*actor).State["Idown"] = false
	}
}

func keybinderRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
}
