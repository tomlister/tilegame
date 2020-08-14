package main

import (
	"github.com/tomlister/tilegame/engine/actor"
	"github.com/tomlister/tilegame/engine/world"
)

func wolfActorLogic(a *actor.Actor, w *world.World, sceneDidMove bool) {
	i := (*w).TagTable["Player"]
	distancex := (float64((*a).X - ((*w).Actors[i].X)))
	distancey := (float64((*a).Y - ((*w).Actors[i].Y)))
	if distancex > 128 {
		(*a).VelocityX += 1
	} else if distancex < -128 {
		(*a).VelocityX -= 1
	}
	if distancey > 128 {
		(*a).VelocityY += 1
	} else if distancey < -128 {
		(*a).VelocityY -= 1
	}

	if distancex > 64 {
		(*a).VelocityX += 0.5
	} else if distancex < -64 {
		(*a).VelocityX -= 0.5
	}

	if distancey > 64 {
		(*a).VelocityY += 0.5
	} else if distancey < -64 {
		(*a).VelocityY -= 0.5
	}
	(*a).X -= int((*a).VelocityX)
	(*a).Y -= int((*a).VelocityY)
	a.ApplyFriction()
}
