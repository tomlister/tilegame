package main

func wolfActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	i := (*world).TagTable["Player"]
	distancex := (float64((*actor).X - ((*world).Actors[i].X)))
	distancey := (float64((*actor).Y - ((*world).Actors[i].Y)))
	if distancex > 128 {
		(*actor).VelocityX += 1
	} else if distancex < -128 {
		(*actor).VelocityX -= 1
	}
	if distancey > 128 {
		(*actor).VelocityY += 1
	} else if distancey < -128 {
		(*actor).VelocityY -= 1
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
	(*actor).X -= int((*actor).VelocityX)
	(*actor).Y -= int((*actor).VelocityY)
	actor.applyFriction()
}
