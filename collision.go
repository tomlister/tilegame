package main

/*
	TODO:
	Fix DetectPlayerCollision
*/

func (actor *Actor) DetectPlayerCollision(world *World) (bool, int, int) {
	aw, ah := (*actor).Image.Size()
	i := (*world).TagTable["Player"]
	collided := false
	shiftx := 0
	shifty := 0
	if (*world).Actors[i].Y < (*actor).Y+ah && (*world).Actors[i].Y > (*actor).Y && (*world).Actors[i].X < (*actor).X+aw && (*world).Actors[i].X > (*actor).X {
		collided = true
		shifty = 1
	}
	return collided, shiftx, shifty
}

func (actor *Actor) DetectPickup(world *World) bool {
	aw, ah := (*actor).Image.Size()
	i := (*world).TagTable["Player"]
	paw, pah := (*world).Actors[i].Image.Size()
	if (*world).Actors[i].X < (*actor).X+aw &&
		(*world).Actors[i].X+paw > (*actor).X &&
		(*world).Actors[i].Y < (*actor).Y+ah &&
		(*world).Actors[i].Y+pah > (*actor).Y {
		return true
	}
	return false
}

func (actor *Actor) DetectTileTouch(world *World, tag string) bool {
	aw, ah := (*actor).Image.Size()
	for i := 0; i < len((*world).Actors); i++ {
		if (*world).Actors[i].Tag == tag {
			paw, pah := 32, 32
			if (*world).Actors[i].X < (*actor).X+aw &&
				(*world).Actors[i].X+paw > (*actor).X &&
				(*world).Actors[i].Y < (*actor).Y+ah &&
				(*world).Actors[i].Y+pah > (*actor).Y {
				return true
			}
		}
	}
	return false
}

func (actor *Actor) DetectBehind(world *World) bool {
	aw, ah := (*actor).Image.Size()
	i := (*world).TagTable["Player"]
	paw, pah := (*world).Actors[i].Image.Size()
	if (*world).Actors[i].X < (*actor).X+aw &&
		(*world).Actors[i].X+paw > (*actor).X &&
		(*world).Actors[i].Y < (*actor).Y+(ah/2) &&
		(*world).Actors[i].Y+pah > (*actor).Y {
		return true
	}
	return false
}
