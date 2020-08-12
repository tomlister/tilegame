package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

//lint:ignore U1000 Stubs
func titleActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		(*actor).Kill = true
		actorSetup(world, windowsettings)
		ebiten.SetCursorMode(ebiten.CursorModeHidden)
	}
}

func titleRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	s := (*pipelinewrapper.World).Shaders["title"]
	cx, cy := ebiten.CursorPosition()
	op := &ebiten.DrawRectShaderOptions{}
	op.GeoM.Scale(2, 2)
	op.GeoM.Translate(float64((Width/2)-256), float64(50))
	op.Uniforms = []interface{}{
		[]float32{float32(cx), float32(cy)}, // Cursor
	}
	diffuse := (*pipelinewrapper.World).getImage("titlediffuse")
	dw, dh := diffuse.Size()
	normal := (*pipelinewrapper.World).getImage("titlenormal")
	op.Images[0] = diffuse
	op.Images[1] = normal
	screen.DrawRectShader(dw, dh, s, op)
	/*
		Start Text
	*/
	text.Draw(screen, "Press ENTER to Start", (*pipelinewrapper.World.Font[1]), 175, Height-50, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xff})
}
