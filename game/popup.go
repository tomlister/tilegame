package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/tomlister/tilegame/engine/actor"
	"github.com/tomlister/tilegame/engine/pipeline"
	"github.com/tomlister/tilegame/engine/world"
)

func popupActorLogic(a *actor.Actor, world *world.World, sceneDidMove bool) {
	if (*a).State["Interval"].(int) == 1 {
		(*a).State["Interval"] = 0
		if (*world).State["popup"].(bool) {
			if a.Y > 400 {
				a.Y -= 5
			}
		} else {
			if a.Y < 480 {
				a.Y += 5
			}
		}
	} else {
		(*a).State["Interval"] = (*a).State["Interval"].(int) + 1
	}
	if (*world).State["popuptimeout"].(int) == 0 {
		(*world).State["popup"] = false
	} else {
		(*world).State["popuptimeout"] = (*world).State["popuptimeout"].(int) - 1
	}
}

func popupRenderCode(a *actor.Actor, pipelinewrapper pipeline.PipelineWrapper, screen *ebiten.Image) {
	/*
		Draw the popup background
	*/
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(4, 4)
	opts.GeoM.Translate(float64((*a).X), float64((*a).Y))
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
		opts2.GeoM.Translate(float64((*a).X+20), float64((*a).Y+47))
		screen.DrawImage(element, opts2)
		element.Dispose()
	}

	/*
		Draw the element bar
	*/
	opts = &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(4, 4)
	opts.GeoM.Translate(float64((*a).X), float64((*a).Y)-45)
	screen.DrawImage((*pipelinewrapper.World).getImage("elementbar"), opts)
	text.Draw(screen, "Mana", (*pipelinewrapper.World.Font[0]), a.X+20, a.Y+30, color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff})
}
