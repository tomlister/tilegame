package engine

func (world *World) ApplyFriction() {
	(*world).VelocityX -= (*world).VelocityX / 10
	(*world).VelocityY -= (*world).VelocityY / 10
}
