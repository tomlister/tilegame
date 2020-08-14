package main

import (
	"github.com/tomlister/tilegame/engine/actor"
	"github.com/tomlister/tilegame/engine/world"
)

func spellArcaneActorLogic(a *actor.Actor, world *world.World, sceneDidMove bool) {
	if (*a).State["Interval"].(int) == 5 {
		(*a).State["Interval"] = 0
		if (*a).State["AnimCount"].(int) == len((*a).AltImages)-1 {
			(*a).State["AnimCount"] = 0
			(*a).Kill = true
		} else {
			(*a).State["AnimCount"] = (*a).State["AnimCount"].(int) + 1
		}
		(*a).Image = (*a).AltImages[(*a).State["AnimCount"].(int)]
	} else {
		(*a).State["Interval"] = (*a).State["Interval"].(int) + 1
	}

	tx := (*a).State["targetx"].(int) - 64
	ty := (*a).State["targety"].(int) - 64
	distancex := (float64((*a).X - tx))
	distancey := (float64((*a).Y - ty))
	if distancex > 128 {
		(*a).VelocityX++
	} else if distancex < -128 {
		(*a).VelocityX--
	}
	if distancey > 128 {
		(*a).VelocityY++
	} else if distancey < -128 {
		(*a).VelocityY--
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

	if distancex > 0 {
		(*a).VelocityX += 0.25
	} else if distancex < 0 {
		(*a).VelocityX -= 0.25
	}

	if distancey > 0 {
		(*a).VelocityY += 0.25
	} else if distancey < 0 {
		(*a).VelocityY -= 0.25
	}

	(*a).X -= int((*a).VelocityX)
	(*a).Y -= int((*a).VelocityY)
	a.ApplyFriction()
}
