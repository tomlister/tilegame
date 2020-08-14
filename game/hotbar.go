package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/tomlister/tilegame/engine/actor"
	"github.com/tomlister/tilegame/engine/pipeline"
	"github.com/tomlister/tilegame/engine/world"
)

func hotbarActorLogic(a *actor.Actor, world *world.World, sceneDidMove bool) {
	i := (*world).TagTable["Player"]
	tx := 0
	switch (*world).Actors[i].State["hotbarslot"].(int) {
	case 0:
		tx = 32
	case 1:
		tx = 96
	case 2:
		tx = 160
	default:
		tx = -100
	}
	if a.X > tx {
		a.X -= 4
	} else if a.X < tx {
		a.X += 4
	}
}

func hotbarRenderCode(a *actor.Actor, pipelinewrapper pipeline.PipelineWrapper, screen *ebiten.Image) {
	/*
		Draw the hotbar slider
	*/
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(2, 2)
	opts.GeoM.Translate(float64((*a).X), float64((*a).Y))
	screen.DrawImage((*pipelinewrapper.World).getImage("hotbar"), opts)
	/*
		Draw hotbar items
	*/
	/*
		0
	*/
	i := (*pipelinewrapper.World).TagTable["Player"]
	opts = &ebiten.DrawImageOptions{}
	shadowopts := (*opts)
	shadowopts.ColorM.Scale(0, 0, 0, 0.5)
	r := float64(0x00)
	g := float64(0x00)
	b := float64(0x00)
	shadowopts.ColorM.Translate(r, g, b, 0)
	opts.GeoM.Scale(2, 2)
	shadowopts.GeoM.Scale(2, 2)
	shadowopts.GeoM.Translate(float64(float64(32)+(32/16)), float64((*a).Y+(32/16)))
	opts.GeoM.Translate(float64(32), float64((*a).Y))
	screen.DrawImage((*pipelinewrapper.World).getImage((*pipelinewrapper.World).Actors[i].State["hotbar0image"].(string)), &shadowopts)
	screen.DrawImage((*pipelinewrapper.World).getImage((*pipelinewrapper.World).Actors[i].State["hotbar0image"].(string)), opts)
	/*
		1
	*/
	opts = &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(2, 2)
	shadowopts.GeoM.Reset()
	shadowopts.GeoM.Scale(2, 2)
	shadowopts.GeoM.Translate(float64(float64(96)+(32/16)), float64((*a).Y+(32/16)))
	opts.GeoM.Translate(float64(96), float64((*a).Y))
	screen.DrawImage((*pipelinewrapper.World).getImage((*pipelinewrapper.World).Actors[i].State["hotbar1image"].(string)), &shadowopts)
	screen.DrawImage((*pipelinewrapper.World).getImage((*pipelinewrapper.World).Actors[i].State["hotbar1image"].(string)), opts)
	/*
		2
	*/
	opts = &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(2, 2)
	shadowopts.GeoM.Reset()
	shadowopts.GeoM.Scale(2, 2)
	shadowopts.GeoM.Translate(float64(float64(160)+(32/16)), float64((*a).Y+(32/16)))
	opts.GeoM.Translate(float64(160), float64((*a).Y))
	screen.DrawImage((*pipelinewrapper.World).getImage((*pipelinewrapper.World).Actors[i].State["hotbar2image"].(string)), &shadowopts)
	screen.DrawImage((*pipelinewrapper.World).getImage((*pipelinewrapper.World).Actors[i].State["hotbar2image"].(string)), opts)
}
