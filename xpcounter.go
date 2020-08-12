package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

func xpcounterRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	playerid := (*pipelinewrapper.World).TagTable["Player"]
	player := (*pipelinewrapper.World).Actors[playerid]
	text.Draw(screen, fmt.Sprintf("%d", player.State["xp"].(int)), (*pipelinewrapper.World.Font[0]), actor.X, actor.Y, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xff})
}

func xpcounterActorLogic(actor *Actor, world *World, sceneDidMove bool) {
}
