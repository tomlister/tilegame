package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/text"
)

//lint:ignore U1000 Stubs
func traderActorLogic(actor *Actor, world *World, sceneDidMove bool) {

}

//lint:ignore U1000 Stubs
func tradeOfferActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	if (*actor).State["interval"].(int) == 5 {
		if len((*actor).State["text"].(string)) > (*actor).State["pos"].(int) {
			(*actor).State["pos"] = (*actor).State["pos"].(int) + 1
			sePlayer, _ := audio.NewPlayerFromBytes((*world).AudioContext, (*world.Sounds["text"]))
			sePlayer.SetVolume(0.2)
			sePlayer.Play()
		}
		(*actor).State["interval"] = 0
	}
	(*actor).State["interval"] = (*actor).State["interval"].(int) + 1
}

//lint:ignore U1000 Stubs
func tradeOfferRenderLogic(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	/*
		Draw the speech background
	*/
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(4, 4)
	opts.GeoM.Translate(float64((*actor).X), float64((*actor).Y))
	screen.DrawImage((*pipelinewrapper.World).getImage("speech"), opts)
	text.Draw(screen, (*actor).State["text"].(string)[:(*actor).State["pos"].(int)], (*pipelinewrapper.World.Font[0]), (*actor).X+10, (*actor).Y+22, color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff})
}
