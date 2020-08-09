package main

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

func healthActorLogic(actor *Actor, world *World, sceneDidMove bool) {
}

func healthRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	/*
		Draw the blood
	*/
	i := (*pipelinewrapper.World).TagTable["Player"]
	player := (*pipelinewrapper.World).Actors[i]
	healthpercentage := float64(player.State["health"].(int)) / 100.0
	pxhealth := int(healthpercentage * 32)
	if pxhealth > 0 {
		mask := (*pipelinewrapper.World).getImage("vialmask").SubImage(image.Rect(0, 32-pxhealth, 32, 32))
		opts2 := &ebiten.DrawImageOptions{}
		opts2.GeoM.Scale(4, 4)
		opts2.GeoM.Translate(float64((*actor).X), float64((*actor).Y+((32-pxhealth)*4)))
		opts2.ColorM.Scale(0, 0, 0, 1)
		opts2.ColorM.Translate(float64(0x8a)/0xff, float64(0x03)/0xff, float64(0x03)/0xff, 0)
		screen.DrawImage(mask.(*ebiten.Image), opts2)
		mask.(*ebiten.Image).Dispose()
	}
	/*
		Draw the blood vial
	*/
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(4, 4)
	opts.GeoM.Translate(float64((*actor).X), float64((*actor).Y))
	screen.DrawImage((*pipelinewrapper.World).getImage("vial"), opts)
}
