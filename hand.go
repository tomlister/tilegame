package main

import "github.com/hajimehoshi/ebiten"

func handActorLogic(actor *Actor, world *World, sceneDidMove bool) {
}

func handRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	i := (*pipelinewrapper.World).TagTable["Player"]
	opts := &ebiten.DrawImageOptions{}
	shadowopts := (*opts)
	imagename := (*pipelinewrapper.World).Actors[i].State["hotbar"].(Hotbar).Slots[(*pipelinewrapper.World).Actors[i].State["hotbar"].(Hotbar).Slot].ImageName
	shadowopts.ColorM.Scale(0, 0, 0, 0.5)
	r := float64(0x00)
	g := float64(0x00)
	b := float64(0x00)
	shadowopts.ColorM.Translate(r, g, b, 0)
	opts.GeoM.Scale(1, 1)
	shadowopts.GeoM.Scale(1, 1)
	shadowopts.GeoM.Translate(float64(8+(*pipelinewrapper.World).Actors[i].X+(*pipelinewrapper.World).CameraX+32/8), float64(8+(*pipelinewrapper.World).Actors[i].Y+(*pipelinewrapper.World).CameraY)-32/8)
	opts.GeoM.Translate(float64(8+(*pipelinewrapper.World).Actors[i].X+(*pipelinewrapper.World).CameraX), float64(8+(*pipelinewrapper.World).Actors[i].Y+(*pipelinewrapper.World).CameraY))
	screen.DrawImage((*pipelinewrapper.World).getImage(imagename), &shadowopts)
	screen.DrawImage((*pipelinewrapper.World).getImage(imagename), opts)
}
