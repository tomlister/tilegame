package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/audio"
)

func spellArcaneActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	if (*actor).State["Interval"].(int) == 5 {
		(*actor).State["Interval"] = 0
		if (*actor).State["AnimCount"].(int) == len((*actor).AltImages)-1 {
			(*actor).State["AnimCount"] = 0
			(*actor).Kill = true
		} else {
			(*actor).State["AnimCount"] = (*actor).State["AnimCount"].(int) + 1
		}
		(*actor).Image = (*actor).AltImages[(*actor).State["AnimCount"].(int)]
		i, collided := world.detectCollisionPointTag(actor.X+63, actor.Y+54, "enemy")
		if collided {
			profile := (*world).Actors[i].State["profile"].(Enemy)
			profile.Health = profile.Health - 10
			(*world).Actors[i].State["profile"] = profile
			minusten := Actor{
				Image:      (*world).getImage("minusten"),
				ActorLogic: floaterActorLogic,
				State:      make(map[string]interface{}),
			}
			minusten.State["Interval"] = 0
			minusten.State["AnimCount"] = 0
			world.spawnActor(minusten, (*world).Actors[i].X, (*world).Actors[i].Y-16)
			soundnames := []string{"hit1", "hit2", "hit3", "hit4"}
			soundindex := rand.Intn(len(soundnames))
			soundname := soundnames[soundindex]
			sePlayer, _ := audio.NewPlayerFromBytes((*world).AudioContext, (*world.Sounds[soundname]))
			sePlayer.Play()
		}
	} else {
		(*actor).State["Interval"] = (*actor).State["Interval"].(int) + 1
	}

	tx := (*actor).State["targetx"].(int) - 64
	ty := (*actor).State["targety"].(int) - 64
	distancex := (float64((*actor).X - tx))
	distancey := (float64((*actor).Y - ty))
	if distancex > 128 {
		(*actor).VelocityX++
	} else if distancex < -128 {
		(*actor).VelocityX--
	}
	if distancey > 128 {
		(*actor).VelocityY++
	} else if distancey < -128 {
		(*actor).VelocityY--
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

	if distancex > 0 {
		(*actor).VelocityX += 0.25
	} else if distancex < 0 {
		(*actor).VelocityX -= 0.25
	}

	if distancey > 0 {
		(*actor).VelocityY += 0.25
	} else if distancey < 0 {
		(*actor).VelocityY -= 0.25
	}

	(*actor).X -= int((*actor).VelocityX)
	(*actor).Y -= int((*actor).VelocityY)
	actor.applyFriction()
}
