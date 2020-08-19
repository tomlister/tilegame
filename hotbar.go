package main

import (
	"github.com/hajimehoshi/ebiten"
)

func hotbarActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	i := (*world).TagTable["Player"]
	tx := 0
	switch (*world).Actors[i].State["hotbar"].(Hotbar).Slot {
	case 0:
		tx = 32
	case 1:
		tx = 96
	case 2:
		tx = 160
	default:
		tx = -100
	}
	if actor.X > tx {
		actor.X -= 4
	} else if actor.X < tx {
		actor.X += 4
	}
}

func hotbarRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	/*
		Draw the hotbar slider
	*/
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(2, 2)
	opts.GeoM.Translate(float64((*actor).X), float64((*actor).Y))
	screen.DrawImage((*pipelinewrapper.World).getImage("hotbar"), opts)
	/*
		Draw hotbar items
	*/
	i := (*pipelinewrapper.World).TagTable["Player"]
	opts = &ebiten.DrawImageOptions{}
	shadowopts := (*opts)
	/*
		0
	*/
	imagename := (*pipelinewrapper.World).Actors[i].State["hotbar"].(Hotbar).Slots[0].ImageName
	if imagename != "" {
		shadowopts.ColorM.Scale(0, 0, 0, 0.5)
		r := float64(0x00)
		g := float64(0x00)
		b := float64(0x00)
		shadowopts.ColorM.Translate(r, g, b, 0)
		opts.GeoM.Scale(2, 2)
		shadowopts.GeoM.Scale(2, 2)
		shadowopts.GeoM.Translate(float64(float64(32)+(32/16)), float64((*actor).Y+(32/16)))
		opts.GeoM.Translate(float64(32), float64((*actor).Y))
		screen.DrawImage((*pipelinewrapper.World).getImage(imagename), &shadowopts)
		screen.DrawImage((*pipelinewrapper.World).getImage(imagename), opts)
	}
	/*
		1
	*/
	imagename = (*pipelinewrapper.World).Actors[i].State["hotbar"].(Hotbar).Slots[1].ImageName
	if imagename != "" {
		opts = &ebiten.DrawImageOptions{}
		opts.GeoM.Scale(2, 2)
		shadowopts.GeoM.Reset()
		shadowopts.GeoM.Scale(2, 2)
		shadowopts.GeoM.Translate(float64(float64(96)+(32/16)), float64((*actor).Y+(32/16)))
		opts.GeoM.Translate(float64(96), float64((*actor).Y))
		screen.DrawImage((*pipelinewrapper.World).getImage(imagename), &shadowopts)
		screen.DrawImage((*pipelinewrapper.World).getImage(imagename), opts)
	}
	/*
		2
	*/
	imagename = (*pipelinewrapper.World).Actors[i].State["hotbar"].(Hotbar).Slots[2].ImageName
	if imagename != "" {
		opts = &ebiten.DrawImageOptions{}
		opts.GeoM.Scale(2, 2)
		shadowopts.GeoM.Reset()
		shadowopts.GeoM.Scale(2, 2)
		shadowopts.GeoM.Translate(float64(float64(160)+(32/16)), float64((*actor).Y+(32/16)))
		opts.GeoM.Translate(float64(160), float64((*actor).Y))
		screen.DrawImage((*pipelinewrapper.World).getImage(imagename), &shadowopts)
		screen.DrawImage((*pipelinewrapper.World).getImage(imagename), opts)
	}
}
