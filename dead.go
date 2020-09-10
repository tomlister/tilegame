package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/text"
)

func deadRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	bounds := text.BoundString((*pipelinewrapper.World.Font[0]), "You died!")
	text.Draw(screen, "You died!", (*pipelinewrapper.World.Font[0]), (Width/2)-(bounds.Dx()/2), 40, color.RGBA{0xff, 0x00, 0x00, 0xff})
	/*
		Restart button
	*/
	mx, my := ebiten.CursorPosition()
	bounds = text.BoundString((*pipelinewrapper.World.Font[3]), "Restart")
	rect := Rect{(Width / 2) - ((bounds.Dx() + 20) / 2), Height - 90, bounds.Dx() + 20, bounds.Dy() + 10}
	opts := &ebiten.DrawImageOptions{}
	itembg, _ := ebiten.NewImage(rect.w, rect.h, ebiten.FilterDefault)
	if detectPointRect(mx, my, rect) {
		itembg.Fill(color.RGBA{0x77, 0x77, 0x77, 0xff})
	} else {
		itembg.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	}
	opts.GeoM.Translate(float64(rect.x), float64(rect.y))
	screen.DrawImage(itembg, opts)
	text.Draw(screen, "Restart", (*pipelinewrapper.World.Font[3]), rect.x+10, rect.y+20, color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff})
}
func deadActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	mx, my := ebiten.CursorPosition()
	bounds := text.BoundString((*world.Font[3]), "Restart")
	rect := Rect{(Width / 2) - ((bounds.Dx() + 20) / 2), Height - 90, bounds.Dx() + 20, bounds.Dy() + 10}
	if detectPointRect(mx, my, rect) {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			(*world).State["musicplayer"].(*audio.Player).Pause()
			(*world).State["pause"] = false
			(*actor).Kill = true
			actorSetup(world, windowsettings, nil)
			ebiten.SetCursorMode(ebiten.CursorModeHidden)
			sePlayer, _ := audio.NewPlayerFromBytes((*world).AudioContext, (*world.Sounds["select1"]))
			sePlayer.Play()
		}
	}
}
