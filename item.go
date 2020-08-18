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
	i := (*world).TagTable["Player"]
	if collided {
		for j, item := range (*world).Actors[i].State["inventory"].([]Item) {
			if item.Name == (*actor).State["item"].(Item).Name {
				tempitem := item
				tempitem.Quantity = tempitem.Quantity + (*actor).State["item"].(Item).Quantity
				(*world).Actors[i].State["inventory"].([]Item)[j] = tempitem
			}
			break
		}
		(*actor).Kill = true
	}

	(*actor).X -= int((*actor).VelocityX)
	(*actor).Y -= int((*actor).VelocityY)
	actor.applyFriction()
}
