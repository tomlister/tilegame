package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/text"
)

//lint:ignore U1000 Stubs
func titleActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		(*actor).Kill = true
		actorSetup(world, windowsettings)
		ebiten.SetCursorMode(ebiten.CursorModeHidden)
	}
	mx, my := ebiten.CursorPosition()
	rect := Rect{(Width / 2) - 75, Height - 100, 140, 50}
	if detectPointRect(mx, my, rect) {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			(*actor).Kill = true
			actorSetup(world, windowsettings)
			ebiten.SetCursorMode(ebiten.CursorModeHidden)
			sePlayer, _ := audio.NewPlayerFromBytes((*world).AudioContext, (*world.Sounds["select1"]))
			sePlayer.Play()
		}
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
	rect := Rect{(Width / 2) - 75, Height - 100, 140, 50}
	opts := &ebiten.DrawImageOptions{}
	itembg, _ := ebiten.NewImage(rect.w, rect.h, ebiten.FilterDefault)
	itembg.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	opts.GeoM.Translate(float64(rect.x), float64(rect.y))
	screen.DrawImage(itembg, opts)
	text.Draw(screen, "Start", (*pipelinewrapper.World.Font[3]), rect.x+32, rect.y+32, color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff})
}
