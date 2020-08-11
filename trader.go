package main

//lint:ignore U1000 Stubs
func traderActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	i := (*world).TagTable["Player"]
	distancex := (float64((*actor).X - ((*world).Actors[i].X)))
	distancey := (float64((*actor).Y - ((*world).Actors[i].Y)))
	if distancex > 128 {
		(*actor).VelocityX += 0.5
	} else if distancex < -128 {
		(*actor).VelocityX -= 0.5
	}
	if distancey > 128 {
		(*actor).VelocityY += 0.5
	} else if distancey < -128 {
		(*actor).VelocityY -= 0.5
	}
	if !(*world).State["FoundHenny"].(bool) {
		world.UIcreateSpeechBubble("Hello!", (*actor).X, (*actor).Y, 265)
	} else {
		world.UIcreateSpeechBubble("Grouse", (*actor).X, (*actor).Y, 50)
	}
	(*actor).X -= int((*actor).VelocityX)
	(*actor).Y -= int((*actor).VelocityY)
	actor.applyFriction()
}
