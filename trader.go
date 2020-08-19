package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/text"
)

func traderActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	i := (*world).TagTable["Player"]
	rect := Rect{(*actor).X - 100, (*actor).Y - 100, (*actor).X + 100, (*actor).Y + 100}
	if detectPointRect((*world).Actors[i].X, (*world).Actors[i].Y, rect) {
		if !(*actor).State["inspeech"].(bool) {
			(*actor).State["inspeech"] = true
			speech := Actor{
				Tag:        "speech",
				Renderhook: true,
				Rendercode: tradeOfferRenderLogic,
				ActorLogic: tradeOfferActorLogic,
				Static:     true,
				Z:          3,
				State:      make(map[string]interface{}),
			}
			speech.State["interval"] = 0
			speech.State["text"] = "Would you like\nto trade?"
			speech.State["pos"] = 0
			speech.State["arrowyoffset"] = 0
			speech.State["time"] = 0
			world.spawnActor(speech, Width/2, Height-128)
		}
	}
}

//lint:ignore U1000 Stubs
func tradeOfferActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		choices := Actor{
			Tag:        "choices",
			Renderhook: true,
			Rendercode: tradeChoiceRenderLogic,
			ActorLogic: tradeChoiceActorLogic,
			Static:     true,
			Z:          3,
			State:      make(map[string]interface{}),
		}
		choices.State["choice1text"] = "Sure."
		choices.State["choice2text"] = "No thanks!"
		world.spawnActor(choices, Width/2, Height-128)
		(*actor).Kill = true
	}
	(*actor).State["time"] = (*actor).State["time"].(int) + 1
	if (*actor).State["interval"].(int) == 5 {
		if len((*actor).State["text"].(string)) > (*actor).State["pos"].(int) {
			(*actor).State["pos"] = (*actor).State["pos"].(int) + 1
			sePlayer, _ := audio.NewPlayerFromBytes((*world).AudioContext, (*world.Sounds["text"]))
			sePlayer.SetVolume(0.2)
			sePlayer.Play()
		}
		(*actor).State["arrowyoffset"] = int(math.Sin(float64((*actor).State["time"].(int))) * 4)
		(*actor).State["interval"] = 0
	}
	(*actor).State["interval"] = (*actor).State["interval"].(int) + 1
}

func tradeOfferRenderLogic(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	/*
		Draw the speech background
	*/
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(4, 4)
	opts.GeoM.Translate(float64((*actor).X), float64((*actor).Y))
	screen.DrawImage((*pipelinewrapper.World).getImage("speech"), opts)
	text.Draw(screen, (*actor).State["text"].(string)[:(*actor).State["pos"].(int)], (*pipelinewrapper.World.Font[0]), (*actor).X+10, (*actor).Y+22, color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff})
	opts = &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(2, 2)
	opts.GeoM.Translate(float64(((*actor).X + 320 - 32)), float64(((*actor).Y + 128 - 32 - (*actor).State["arrowyoffset"].(int))))
	screen.DrawImage((*pipelinewrapper.World).getImage("arrow"), opts)
}

func tradeChoiceActorLogic(actor *Actor, world *World, sceneDidMove bool) {

}

func tradeChoiceRenderLogic(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	/*
		Draw choice 1 background
	*/
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(4, 4)
	opts.GeoM.Translate(float64((*actor).X), float64((*actor).Y))
	screen.DrawImage((*pipelinewrapper.World).getImage("choice"), opts)
	text.Draw(screen, (*actor).State["choice1text"].(string), (*pipelinewrapper.World.Font[0]), (*actor).X+10, (*actor).Y+22, color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff})
	/*
		Draw choice 2 background
	*/
	opts = &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(4, 4)
	opts.GeoM.Translate(float64((*actor).X), float64((*actor).Y+64))
	screen.DrawImage((*pipelinewrapper.World).getImage("choice"), opts)
	text.Draw(screen, (*actor).State["choice2text"].(string), (*pipelinewrapper.World.Font[0]), (*actor).X+10, (*actor).Y+64+22, color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff})
}
