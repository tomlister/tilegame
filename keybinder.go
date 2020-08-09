package main

import (
	"github.com/hajimehoshi/ebiten"
)

func keybinderActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	if ebiten.IsKeyPressed(ebiten.KeyI) {
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

func keybinderRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
}
