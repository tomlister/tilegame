package engine

import "github.com/tomlister/tilegame/engine/world"

/*
	TODO:
	Fix DetectPlayerCollision
*/

func (a *Actor) DetectPlayerCollision(world *world.World) (bool, int, int) {
	aw, ah := (*a).Image.Size()
	i := (*world).TagTable["Player"]
	collided := false
	shiftx := 0
	shifty := 0
	if (*world).Actors[i].Y < (*a).Y+ah && (*world).Actors[i].Y > (*a).Y && (*world).Actors[i].X < (*a).X+aw && (*world).Actors[i].X > (*a).X {
		collided = true
		shifty = 1
	}
	return collided, shiftx, shifty
}

func (a *Actor) DetectPickup(world *world.World) bool {
	aw, ah := (*a).Image.Size()
	i := (*world).TagTable["Player"]
	paw, pah := (*world).Actors[i].Image.Size()
	if (*world).Actors[i].X < (*a).X+aw &&
		(*world).Actors[i].X+paw > (*a).X &&
		(*world).Actors[i].Y < (*a).Y+ah &&
		(*world).Actors[i].Y+pah > (*a).Y {
		return true
	}
	return false
}

func (a *Actor) DetectTileTouch(world *world.World, tag string) bool {
	aw, ah := (*a).Image.Size()
	for i := 0; i < len((*world).Actors); i++ {
		if (*world).Actors[i].Tag == tag {
			paw, pah := 32, 32
			if (*world).Actors[i].X < (*a).X+aw &&
				(*world).Actors[i].X+paw > (*a).X &&
				(*world).Actors[i].Y < (*a).Y+ah &&
				(*world).Actors[i].Y+pah > (*a).Y {
				return true
			}
		}
	}
	return false
}

func (a *Actor) DetectBehind(world *world.World) bool {
	aw, ah := (*a).Image.Size()
	i := (*world).TagTable["Player"]
	paw, pah := (*world).Actors[i].Image.Size()
	if (*world).Actors[i].X < (*a).X+aw &&
		(*world).Actors[i].X+paw > (*a).X &&
		(*world).Actors[i].Y < (*a).Y+(ah/2) &&
		(*world).Actors[i].Y+pah > (*a).Y {
		return true
	}
	return false
}
