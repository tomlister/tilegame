package main

import (
	"github.com/tomlister/tilegame/engine/actor"
	"github.com/tomlister/tilegame/engine/world"
)

func plusOneActorLogic(a *actor.Actor, world *world.World, sceneDidMove bool) {
	if (*a).State["Interval"].(int) == 5 {
		(*a).State["Interval"] = 0
		/*if (*a).State["AnimCount"].(int) == 20 {
			(*a).State["AnimCount"] = 0
			(*a).Kill = true
		} else {
			(*a).State["AnimCount"] = (*a).State["AnimCount"].(int) + 1
		}*/
		i := (*world).TagTable["Player"]
		sx, sy := (*world).Actors[i].Image.Size()
		distancex := (float64((*a).X - ((*world).Actors[i].X + (sx / 2))))
		distancey := (float64((*a).Y - ((*world).Actors[i].Y + (sy / 2))))
		if distancex > 8 && (distancex < 32 || distancey < 32) {
			(*a).VelocityX++
			if distancex > 16 {
				(*a).VelocityX++
			}
		} else if distancex < -8 && (distancey > -32 || distancex > -32) {
			(*a).VelocityX--
			if distancex < -16 {
				(*a).VelocityX--
			}
		}
		if distancey > 8 && (distancex < 32 || distancey < 32) {
			(*a).VelocityY++
			if distancey > 16 {
				(*a).VelocityY++
			}
		} else if distancey < -8 && (distancey > -32 || distancex > -32) {
			(*a).VelocityY--
			if distancey < -16 {
				(*a).VelocityY--
			}
		}

		if distancex < 8 && distancex > -8 {
			if distancey < 8 && distancey > -8 {
				(*world).Actors[i].State["xp"] = (*world).Actors[i].State["xp"].(int) + 1
				(*a).Kill = true
			}
		}
	} else {
		(*a).State["Interval"] = (*a).State["Interval"].(int) + 1
	}
	(*a).X -= int((*a).VelocityX)
	(*a).Y -= int((*a).VelocityY)
	a.ApplyFriction()
}

func chestActorLogic(a *actor.Actor, w *world.World, sceneDidMove bool) {
	if sceneDidMove {
		if !(*a).State["Opened"].(bool) {
			if a.DetectPickup(w) {
				plusone := actor.Actor{
					Image:      (*w).Images["plusone"],
					ActorLogic: plusOneActorLogic,
					Z:          1,
					State:      make(map[string]interface{}),
				}
				plusone.State["AnimCount"] = 0
				plusone.State["Interval"] = 0
				w.SpawnActor(plusone, (*a).X, (*a).Y)
				(*a).Image = (*a).AltImages[1]
				(*a).State["Opened"] = true
			}
		}
	}
}
