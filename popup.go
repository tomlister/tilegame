package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

func popupActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	if (*actor).State["Interval"].(int) == 1 {
		(*actor).State["Interval"] = 0
		if (*world).State["popup"].(bool) == true {
			if actor.Y > 400 {
				actor.Y -= 5
			}
		} else {
			if actor.Y < 480 {
				actor.Y += 5
			}
		}
	} else {
		(*actor).State["Interval"] = (*actor).State["Interval"].(int) + 1
	}
	if (*world).State["popuptimeout"].(int) == 0 {
		(*world).State["popup"] = false
	} else {
		(*world).State["popuptimeout"] = (*world).State["popuptimeout"].(int) - 1
	}
}

func popupRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	/*
		Draw the popup background
	*/
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(4, 4)
	opts.GeoM.Translate(float64((*actor).X), float64((*actor).Y))
	screen.DrawImage((*pipelinewrapper.World).getImage("popup"), opts)

	/*
		Draw the liquid inside the element bar
	*/
	i := (*pipelinewrapper.World).TagTable["Player"]
	player := (*pipelinewrapper.World).Actors[i]
	manapercentage := float64(player.State["mana"].(int)) / float64(player.State["manamax"].(int))
	pxmana := int(manapercentage * 280)
	if pxmana > 0 {
		element, _ := ebiten.NewImage(pxmana, 24, ebiten.FilterDefault)
		element.Fill(color.RGBA{R: 0x67, G: 0x00, B: 0x74, A: 0xFF})
		opts2 := &ebiten.DrawImageOptions{}
		opts2.GeoM.Translate(float64((*actor).X+20), float64((*actor).Y+47))
		screen.DrawImage(element, opts2)
		element.Dispose()
	}

	/*
		Draw the element bar
	*/
	opts = &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(4, 4)
	opts.GeoM.Translate(float64((*actor).X), float64((*actor).Y)-45)
	screen.DrawImage((*pipelinewrapper.World).getImage("elementbar"), opts)
	text.Draw(screen, "Mana", (*pipelinewrapper.World.Font), actor.X+20, actor.Y+30, color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff})
}
