package main

func droppedItemActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	tx := (*actor).State["targetx"].(int) - 32
	distancex := (float64((*actor).X - tx))
	if distancex > 128 {
		(*actor).VelocityX++
	} else if distancex < -128 {
		(*actor).VelocityX--
	}

	if distancex > 64 {
		(*actor).VelocityX += 0.5
	} else if distancex < -64 {
		(*actor).VelocityX -= 0.5
	}

	if distancex > 0 {
		(*actor).VelocityX += 0.25
	} else if distancex < 0 {
		(*actor).VelocityX -= 0.25
	}
	collided, _, _ := actor.DetectPlayerCollision(world)
	if collided {
		(*actor).Kill = true
	}

	(*actor).X -= int((*actor).VelocityX)
	(*actor).Y -= int((*actor).VelocityY)
	actor.applyFriction()
}
