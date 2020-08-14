package main

import (
	"github.com/tomlister/tilegame/engine/actor"
	"github.com/tomlister/tilegame/engine/world"
)

func backgroundActorLogic(a *actor.Actor, w *world.World, sceneDidMove bool) {

}

func backgroundTreeActorLogic(a *actor.Actor, w *world.World, sceneDidMove bool) {
	if a.DetectBehind(w) {
		(*a).Z = 1
	} else {
		(*a).Z = 0
	}
	if (*a).State["Interval"].(int) == 25 {
		(*a).State["Interval"] = 0
		if (*a).State["AnimCount"].(int) == len((*a).AltImages)-1 {
			(*a).State["AnimCount"] = 0
		} else {
			(*a).State["AnimCount"] = (*a).State["AnimCount"].(int) + 1
		}
		(*a).Image = (*a).AltImages[(*a).State["AnimCount"].(int)]
	} else {
		(*a).State["Interval"] = (*a).State["Interval"].(int) + 1
	}
	if (*a).State["health"] == 0 {
		(*a).Kill = true
	}
}
