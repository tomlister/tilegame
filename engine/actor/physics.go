package engine

func (a *Actor) ApplyFriction() {
	(*a).VelocityX -= (*a).VelocityX / 10
	(*a).VelocityY -= (*a).VelocityY / 10
}
