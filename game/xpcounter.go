package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/tomlister/tilegame/engine/actor"
	"github.com/tomlister/tilegame/engine/pipeline"
	"github.com/tomlister/tilegame/engine/world"
)

func xpcounterRenderCode(a *actor.Actor, pipelinewrapper pipeline.PipelineWrapper, screen *ebiten.Image) {
	playerid := (*pipelinewrapper.World).TagTable["Player"]
	player := (*pipelinewrapper.World).Actors[playerid]
	text.Draw(screen, fmt.Sprintf("%d", player.State["xp"].(int)), (*pipelinewrapper.World.Font[0]), actor.X, actor.Y, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xff})
}

func xpcounterActorLogic(a *actor.Actor, world *world.World, sceneDidMove bool) {
}
