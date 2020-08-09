package main

import "github.com/hajimehoshi/ebiten"

func resetAllStrafe(actor *Actor) {
	(*actor).X += (*actor).State["StrafeLeft"].(int)
	(*actor).X -= (*actor).State["StrafeRight"].(int)
	(*actor).Y += (*actor).State["StrafeUp"].(int)
	(*actor).Y -= (*actor).State["StrafeDown"].(int)
	(*actor).State["StrafeLeft"] = 0
	(*actor).State["StrafeRight"] = 0
	(*actor).State["StrafeUp"] = 0
	(*actor).State["StrafeDown"] = 0
}

func crossHairActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	cursorx, cursory := ebiten.CursorPosition()
	(*actor).X = cursorx - 8
	(*actor).Y = cursory - 8
	/*cursorx, cursory := ebiten.CursorPosition()
	playerImageWidth, playerImageHeight := (*world).Actors[i].Image.Size()
	arrowImageWidth, arrowImageHeight := (*actor).Image.Size()
	x := (*world).Actors[i].X + (playerImageWidth / 2) + (*world).CameraX
	y := (*world).Actors[i].Y + (playerImageHeight / 2) + (*world).CameraY
	dx := float64(cursorx) - float64(x)
	dy := float64(cursory) - float64(y)
	dir := math.Atan2(dy, dx) - math.Pi
	newY := int(math.Round(float64(y-(*world).CameraY) + float64(-32)*math.Sin(dir)))
	newX := int(math.Round(float64(x-(*world).CameraX) + float64(-32)*math.Cos(dir)))
	(*actor).X = newX - (arrowImageWidth / 2)
	(*actor).Y = newY - (arrowImageHeight / 2)*/
	//(*actor).Direction = dir + math.Pi
}

func playerSplashActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	i := (*world).TagTable["Player"]
	(*actor).X = (*world).Actors[i].X
	(*actor).Y = (*world).Actors[i].Y + 3
	if (*world).Actors[i].State["Swimming"].(bool) == true {
		(*actor).Disabled = false
	} else {
		(*actor).Disabled = true
	}
}

func playerWandShoot(actor *Actor, world *World) {
	(*world).State["popup"] = true
	(*world).State["popuptimeout"] = 200
	if (*actor).State["mana"].(int) > 0 {
		(*actor).State["mana"] = (*actor).State["mana"].(int) - 1
		spell := Actor{
			Image: (*world).getImage("spellarcane0"),
			AltImages: []*ebiten.Image{
				(*world).getImage("spellarcane0"),
				(*world).getImage("spellarcane1"),
				(*world).getImage("spellarcane2"),
				(*world).getImage("spellarcane3"),
				(*world).getImage("spellarcane4"),
				(*world).getImage("spellarcane5"),
				(*world).getImage("spellarcane6"),
			},
			ActorLogic: spellArcaneActorLogic,
			Shadow:     true,
			Z:          1,
			State:      make(map[string]interface{}),
		}
		cursorx, cursory := ebiten.CursorPosition()
		spell.State["targetx"] = cursorx - (*world).CameraX
		spell.State["targety"] = cursory - (*world).CameraY
		spell.State["AnimCount"] = 0
		spell.State["Interval"] = 0
		world.spawnActor(spell, (*actor).X-32, (*actor).Y-32)
	}
}

func playerAxeUse(actor *Actor, world *World) {
	cursorx, cursory := ebiten.CursorPosition()
	i, collided := world.detectCollisionPointTag(cursorx-(*world).CameraX, cursory-(*world).CameraY, "tree")
	if collided {
		wood := Actor{
			Image:      (*world).getImage("wooditem"),
			ActorLogic: droppedItemActorLogic,
			Shadow:     true,
			Z:          0,
			State:      make(map[string]interface{}),
		}
		wood.State["targetx"] = cursorx - (*world).CameraX + 64 + (4 * (*world).Actors[i].State["health"].(int))
		wood.State["targety"] = cursory - (*world).CameraY
		wood.State["Interval"] = 0
		(*world).Actors[i].State["health"] = (*world).Actors[i].State["health"].(int) - 1
		world.spawnActor(wood, cursorx-(*world).CameraX, cursory-(*world).CameraY-(*world).Actors[i].State["health"].(int)*8)
	}
}

func playerHotbarSwitch(actor *Actor, world *World, hotbarname string) {
	switch hotbarname {
	case "Wand":
		playerWandShoot(actor, world)
	case "Iron Axe":
		playerAxeUse(actor, world)
	}
}

func playerActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	(*actor).VelocityX = (*world).VelocityX
	(*actor).VelocityY = (*world).VelocityY
	(*actor).X -= int((*actor).VelocityX)
	(*actor).Y -= int((*actor).VelocityY)
	if ebiten.IsKeyPressed(ebiten.Key1) {
		(*actor).State["hotbarslot"] = 0
	} else if ebiten.IsKeyPressed(ebiten.Key2) {
		(*actor).State["hotbarslot"] = 1
	} else if ebiten.IsKeyPressed(ebiten.Key3) {
		(*actor).State["hotbarslot"] = 2
	}
	if (*actor).State["tooltimeout"].(int) == 0 {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			(*actor).State["tooltimeout"] = 1
			if (*actor).State["hotbarslot"].(int) == 0 {
				playerHotbarSwitch(actor, world, (*actor).State["hotbar0name"].(string))
			} else if (*actor).State["hotbarslot"].(int) == 1 {
				playerHotbarSwitch(actor, world, (*actor).State["hotbar1name"].(string))
			} else if (*actor).State["hotbarslot"].(int) == 2 {
				playerHotbarSwitch(actor, world, (*actor).State["hotbar2name"].(string))
			}
		}
	}
	if (*actor).State["tooltimeout"].(int) == 5 {
		(*actor).State["tooltimeout"] = 0
	} else if (*actor).State["tooltimeout"].(int) != 0 {
		(*actor).State["tooltimeout"] = (*actor).State["tooltimeout"].(int) + 1
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		(*actor).State["health"] = (*actor).State["health"].(int) - 1
	}
	if sceneDidMove {
		if actor.DetectTileTouch(world, "water") && !(*actor).State["Swimming"].(bool) {
			(*actor).State["Swimming"] = true
			(*actor).Shadow = false
			(*actor).Y += 10
		} else if !actor.DetectTileTouch(world, "water") && (*actor).State["Swimming"].(bool) {
			(*actor).State["Swimming"] = false
			(*actor).Shadow = true
			(*actor).Y -= 10
		}
		fsx, fsy := getActorShift()
		sx, sy := int(fsx), int(fsy)
		if sx > 0 {
			if (*actor).State["StrafeLeft"].(int) == 0 {
				(*actor).State["StrafeLeft"] = 10
				(*actor).X -= (*actor).State["StrafeLeft"].(int)
			}
		} else if sx < 0 {
			if (*actor).State["StrafeRight"].(int) == 0 {
				(*actor).State["StrafeRight"] = 10
				(*actor).X += (*actor).State["StrafeRight"].(int)
			}
		} else {
			(*actor).X += (*actor).State["StrafeLeft"].(int)
			(*actor).X -= (*actor).State["StrafeRight"].(int)
			(*actor).State["StrafeLeft"] = 0
			(*actor).State["StrafeRight"] = 0
		}
		if sy > 0 {
			if (*actor).State["StrafeUp"].(int) == 0 {
				(*actor).State["StrafeUp"] = 10
				(*actor).Y -= (*actor).State["StrafeUp"].(int)
			}
		} else if sy < 0 {
			if (*actor).State["StrafeDown"].(int) == 0 {
				(*actor).State["StrafeDown"] = 10
				(*actor).Y += (*actor).State["StrafeDown"].(int)
			}
		} else {
			(*actor).Y += (*actor).State["StrafeUp"].(int)
			(*actor).Y -= (*actor).State["StrafeDown"].(int)
			(*actor).State["StrafeUp"] = 0
			(*actor).State["StrafeDown"] = 0
		}
	} else {
		resetAllStrafe(actor)
	}
}
