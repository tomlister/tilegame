package main

import (
	"github.com/hajimehoshi/ebiten"
)

func backgroundActorLogic(actor *Actor, world *World, sceneDidMove bool) {

}

func backgroundTreeActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	if actor.DetectBehind(world) {
		(*actor).Z = 1
	} else {
		(*actor).Z = 0
	}
	if (*actor).State["Interval"].(int) == 25 {
		(*actor).State["Interval"] = 0
		if (*actor).State["AnimCount"].(int) == len((*actor).AltImages)-1 {
			(*actor).State["AnimCount"] = 0
		} else {
			(*actor).State["AnimCount"] = (*actor).State["AnimCount"].(int) + 1
		}
		(*actor).Image = (*actor).AltImages[(*actor).State["AnimCount"].(int)]
	} else {
		(*actor).State["Interval"] = (*actor).State["Interval"].(int) + 1
	}
	if (*actor).State["health"] == 0 {
		(*actor).Kill = true
	}
}

func backgroundActorRenderHook(actor *Actor, world *World) (*ebiten.Image, *ebiten.DrawImageOptions) {
	img := world.getImage((*actor).State["imagename"].(string))
	//hue := (*actor).State["Hue"].(float64)
	opts := &ebiten.DrawImageOptions{}
	//opts.ColorM.RotateHue(hue)
	return img, opts
}
