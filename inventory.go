package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type Item struct {
	Name      string
	ImageName string
	Quantity  int
}

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
		Draw the inventory grid
	*/
	inventoryGridRenderCode(pipelinewrapper, screen)
}

func inventoryGridRenderCode(pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	i := (*pipelinewrapper.World).TagTable["Player"]
	inventory := (*pipelinewrapper.World).Actors[i].State["inventory"].([]Item)
	x, y := 0, 0
	for j := 0; j < len(inventory); j++ {
		inventory[j].inventoryGridItemRenderCode(160+(x*32), 160+(y*32), pipelinewrapper, screen)
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
	opts.GeoM.Scale(1, 1)
	opts.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage((*pipelinewrapper.World).getImage((*i).ImageName), opts)
}
