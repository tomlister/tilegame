package main

func (world *World) applyFriction() {
	(*world).VelocityX -= (*world).VelocityX / 10
	(*world).VelocityY -= (*world).VelocityY / 10
}

func (actor *Actor) applyFriction() {
	(*actor).VelocityX -= (*actor).VelocityX / 10
	(*actor).VelocityY -= (*actor).VelocityY / 10
}
