package main

func spellArcaneActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	if (*actor).State["Interval"].(int) == 5 {
		(*actor).State["Interval"] = 0
		if (*actor).State["AnimCount"].(int) == len((*actor).AltImages)-1 {
			(*actor).State["AnimCount"] = 0
			(*actor).Kill = true
		} else {
			(*actor).State["AnimCount"] = (*actor).State["AnimCount"].(int) + 1
		}
		(*actor).Image = (*actor).AltImages[(*actor).State["AnimCount"].(int)]
	} else {
		(*actor).State["Interval"] = (*actor).State["Interval"].(int) + 1
	}

	tx := (*actor).State["targetx"].(int) - 64
	ty := (*actor).State["targety"].(int) - 64
	distancex := (float64((*actor).X - tx))
	distancey := (float64((*actor).Y - ty))
	if distancex > 128 {
		(*actor).VelocityX++
	} else if distancex < -128 {
		(*actor).VelocityX--
	}
	if distancey > 128 {
		(*actor).VelocityY++
	} else if distancey < -128 {
		(*actor).VelocityY--
	}

	if distancex > 64 {
		(*actor).VelocityX += 0.5
	} else if distancex < -64 {
		(*actor).VelocityX -= 0.5
	}

	if distancey > 64 {
		(*actor).VelocityY += 0.5
	} else if distancey < -64 {
		(*actor).VelocityY -= 0.5
	}

	if distancex > 0 {
		(*actor).VelocityX += 0.25
	} else if distancex < 0 {
		(*actor).VelocityX -= 0.25
	}

	if distancey > 0 {
		(*actor).VelocityY += 0.25
	} else if distancey < 0 {
		(*actor).VelocityY -= 0.25
	}

	(*actor).X -= int((*actor).VelocityX)
	(*actor).Y -= int((*actor).VelocityY)
	actor.applyFriction()
}
