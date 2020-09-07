package main

type EnemyBehaviour struct {
	Strafe bool
	Spin   bool
	Melee  bool
}

type Enemy struct {
	Name           string
	Health         int
	Behaviour      EnemyBehaviour
	AttackInterval int
	Speed          float64
	ImageName      string
}

func (e Enemy) enemyBehaviourProvider(actor *Actor, world *World) {
	if e.Behaviour.Melee {
		i := (*world).TagTable["Player"]
		rect := Rect{actor.X - 64, actor.Y - 64, 96, 96}
		if detectPointRect((*world).Actors[i].X+16, (*world).Actors[i].Y+16, rect) {
			(*world).Actors[i].State["health"] = (*world).Actors[i].State["health"].(int) - 1
		}
	}
}

func (e Enemy) enemyMovementProvider(actor *Actor, world *World) {
	i := (*world).TagTable["Player"]
	distancex := (float64((*actor).X - ((*world).Actors[i].X)))
	distancey := (float64((*actor).Y - ((*world).Actors[i].Y)))

	if distancex > 2 {
		(*actor).VelocityX += e.Speed
	} else if distancex < -2 {
		(*actor).VelocityX -= e.Speed
	}

	if distancey > 2 {
		(*actor).VelocityY += e.Speed
	} else if distancey < -2 {
		(*actor).VelocityY -= e.Speed
	}
	(*actor).X -= int((*actor).VelocityX)
	(*actor).Y -= int((*actor).VelocityY)
	actor.applyFriction()
}

func enemyActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	profile := (*actor).State["profile"].(Enemy)
	profile.enemyMovementProvider(actor, world)
	if profile.AttackInterval == 30 {
		profile.AttackInterval = 0
		profile.enemyBehaviourProvider(actor, world)
	} else {
		profile.AttackInterval++
	}
	if profile.Health <= 0 {
		(*actor).Kill = true
	}
	(*actor).State["profile"] = profile
}
