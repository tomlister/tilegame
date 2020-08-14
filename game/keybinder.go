package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/tomlister/tilegame/engine/actor"
	"github.com/tomlister/tilegame/engine/pipeline"
	"github.com/tomlister/tilegame/engine/world"
)

func keybinderActorLogic(a *actor.Actor, world *world.World, sceneDidMove bool) {
	if ebiten.IsKeyPressed(ebiten.KeyI) {
		if !(*a).State["Idown"].(bool) {
			(*a).State["Idown"] = true
			if (*world).State["pause"] == false {
				ebiten.SetCursorVisibility(true)
				(*world).State["pause"] = true
				inv := actor.Actor{
					Tag:        "inv",
					Renderhook: true,
					Rendercode: inventoryRenderCode,
					ActorLogic: inventoryActorLogic,
					Static:     true,
					Z:          3,
					State:      make(map[string]interface{}),
					Unpausable: true,
				}
				world.SpawnActor(inv, 0, 0)
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
		(*a).State["Idown"] = false
	}
}

func keybinderRenderCode(a *actor.Actor, pipelinewrapper pipeline.PipelineWrapper, screen *ebiten.Image) {
}
