package main

import (
	"github.com/tomlister/tilegame/engine/actor"
	"github.com/tomlister/tilegame/engine/world"
)

func droppedItemActorLogic(a *actor.Actor, w *world.World, sceneDidMove bool) {
	tx := (*a).State["targetx"].(int) - 32
	distancex := (float64((*a).X - tx))
	if distancex > 128 {
		(*a).VelocityX++
	} else if distancex < -128 {
		(*a).VelocityX--
	}

	if distancex > 64 {
		(*a).VelocityX += 0.5
	} else if distancex < -64 {
		(*a).VelocityX -= 0.5
	}

	if distancex > 0 {
		(*a).VelocityX += 0.25
	} else if distancex < 0 {
		(*a).VelocityX -= 0.25
	}
	collided, _, _ := a.DetectPlayerCollision(w)
	if collided {
		(*a).Kill = true
	}

	(*a).X -= int((*a).VelocityX)
	(*a).Y -= int((*a).VelocityY)
	a.ApplyFriction()
}
