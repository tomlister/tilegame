package main

func plusOneActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	if (*actor).State["Interval"].(int) == 5 {
		(*actor).State["Interval"] = 0
		/*if (*actor).State["AnimCount"].(int) == 20 {
			(*actor).State["AnimCount"] = 0
			(*actor).Kill = true
		} else {
			(*actor).State["AnimCount"] = (*actor).State["AnimCount"].(int) + 1
		}*/
		i := (*world).TagTable["Player"]
		sx, sy := (*world).Actors[i].Image.Size()
		distancex := (float64((*actor).X - ((*world).Actors[i].X + (sx / 2))))
		distancey := (float64((*actor).Y - ((*world).Actors[i].Y + (sy / 2))))
		if distancex > 8 && (distancex < 32 || distancey < 32) {
			(*actor).VelocityX++
			if distancex > 16 {
				(*actor).VelocityX++
			}
		} else if distancex < -8 && (distancey > -32 || distancex > -32) {
			(*actor).VelocityX--
			if distancex < -16 {
				(*actor).VelocityX--
			}
		}
		if distancey > 8 && (distancex < 32 || distancey < 32) {
			(*actor).VelocityY++
			if distancey > 16 {
				(*actor).VelocityY++
			}
		} else if distancey < -8 && (distancey > -32 || distancex > -32) {
			(*actor).VelocityY--
			if distancey < -16 {
				(*actor).VelocityY--
			}
		}

		if distancex < 8 && distancex > -8 {
			if distancey < 8 && distancey > -8 {
				(*world).Actors[i].State["xp"] = (*world).Actors[i].State["xp"].(int) + 1
				(*actor).Kill = true
			}
		}
	} else {
		(*actor).State["Interval"] = (*actor).State["Interval"].(int) + 1
	}
	(*actor).X -= int((*actor).VelocityX)
	(*actor).Y -= int((*actor).VelocityY)
	actor.applyFriction()
}

func chestActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	if sceneDidMove {
		if !(*actor).State["Opened"].(bool) {
			if actor.DetectPickup(world) {
				plusone := Actor{
					Image:      (*world).Images["plusone"],
					ActorLogic: plusOneActorLogic,
					Z:          1,
					State:      make(map[string]interface{}),
				}
				plusone.State["AnimCount"] = 0
				plusone.State["Interval"] = 0
				world.spawnActor(plusone, (*actor).X, (*actor).Y)
				(*actor).Image = (*actor).AltImages[1]
				(*actor).State["Opened"] = true
			}
		}
	}
}
