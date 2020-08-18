package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

type Item struct {
	Name      string
	ImageName string
	Quantity  int
}

//lint:ignore U1000 Stubs
func inventoryActorLogic(actor *Actor, world *World, sceneDidMove bool) {

}

func inventoryRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
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
	text.Draw(screen, "Inventory", (*pipelinewrapper.World.Font[2]), 20, 50, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xff})
	/*
		Draw the inventory grid
	*/
	inventoryGridRenderCode(pipelinewrapper, screen)
}

func inventoryGridRenderCode(pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	i := (*pipelinewrapper.World).TagTable["Player"]
	inventory := (*pipelinewrapper.World).Actors[i].State["inventory"].([]Item)
	x, y := 0, 0
	for j := 0; j < len(inventory); j++ {
		inventory[j].inventoryGridItemRenderCode(160+(x*128), 160+(y*64), pipelinewrapper, screen)
		if x == 2 {
			x = 0
			y++
		} else {
			x++
		}
	}
}

func (i *Item) inventoryGridItemRenderCode(x, y int, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(2, 2)
	opts.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage((*pipelinewrapper.World).getImage((*i).ImageName), opts)
	text.Draw(screen, fmt.Sprintf("%d", (*i).Quantity), (*pipelinewrapper.World.Font[0]), x+32, y+60, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xff})
}
