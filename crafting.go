package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

type Craftable struct {
	Item
	Needs    []Item
	Quantity int
}

//lint:ignore U1000 Stubs
func craftingActorLogic(actor *Actor, world *World, sceneDidMove bool) {

}

func craftingRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	/*
		Draw background
	*/
	sx, sy := screen.Size()
	blackbg, _ := ebiten.NewImage(sx, sy, ebiten.FilterDefault)
	blackbg.Fill(color.RGBA{25, 25, 25, 0xff})
	screen.DrawImage(blackbg, &ebiten.DrawImageOptions{})
	/*
		Draw title
	*/
	text.Draw(screen, "Crafting", (*pipelinewrapper.World.Font[2]), 20, 50, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xff})
	/*
		Draw the crafting list
	*/
	craftingListRenderCode(pipelinewrapper, screen)
}

func craftingListRenderCode(pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	inventory := (*pipelinewrapper.World).State["craftable"].([]Craftable)
	for j := 0; j < len(inventory); j++ {
		inventory[j].craftingListItemRenderCode(160, 100+(j*64), pipelinewrapper, screen)
	}
}

func (i *Item) craftingListItemRenderCode(x, y int, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(2, 2)
	opts.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage((*pipelinewrapper.World).getImage((*i).ImageName), opts)
}
