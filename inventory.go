package main

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

func inventoryActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	//animation controllers
}

func inventoryRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	//blur the background
	//box blur from: https://ebiten.org/examples/blur.html
	screencopy, _ := ebiten.NewImageFromImage(screen.SubImage(image.Rect(320-160, 240-120, 320+160, 240+120)), ebiten.FilterDefault)
	for j := -3; j <= 3; j++ {
		for i := -3; i <= 3; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(2, 2)
			op.GeoM.Translate(float64(i), float64(j))
			op.ColorM.Scale(1, 1, 1, 1.0/25.0)
			screen.DrawImage(screencopy, op)
		}
	}
}
