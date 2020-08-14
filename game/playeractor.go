package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/tomlister/tilegame/engine/actor"
	"github.com/tomlister/tilegame/engine/world"
)

func resetAllStrafe(a *actor.Actor) {
	(*a).X += (*a).State["StrafeLeft"].(int)
	(*a).X -= (*a).State["StrafeRight"].(int)
	(*a).Y += (*a).State["StrafeUp"].(int)
	(*a).Y -= (*a).State["StrafeDown"].(int)
	(*a).State["StrafeLeft"] = 0
	(*a).State["StrafeRight"] = 0
	(*a).State["StrafeUp"] = 0
	(*a).State["StrafeDown"] = 0
}

func crossHairActorLogic(a *actor.Actor, w *world.World, sceneDidMove bool) {
	cursorx, cursory := ebiten.CursorPosition()
	(*a).X = cursorx - 8
	(*a).Y = cursory - 8
	/*cursorx, cursory := ebiten.CursorPosition()
	playerImageWidth, playerImageHeight := (*w).Actors[i].Image.Size()
	arrowImageWidth, arrowImageHeight := (*a).Image.Size()
	x := (*w).Actors[i].X + (playerImageWidth / 2) + (*w).CameraX
	y := (*w).Actors[i].Y + (playerImageHeight / 2) + (*w).CameraY
	dx := float64(cursorx) - float64(x)
	dy := float64(cursory) - float64(y)
	dir := math.Atan2(dy, dx) - math.Pi
	newY := int(math.Round(float64(y-(*w).CameraY) + float64(-32)*math.Sin(dir)))
	newX := int(math.Round(float64(x-(*w).CameraX) + float64(-32)*math.Cos(dir)))
	(*a).X = newX - (arrowImageWidth / 2)
	(*a).Y = newY - (arrowImageHeight / 2)*/
	//(*a).Direction = dir + math.Pi
}

func playerSplashActorLogic(a *actor.Actor, w *world.World, sceneDidMove bool) {
	i := (*w).TagTable["Player"]
	(*a).X = (*w).Actors[i].X
	(*a).Y = (*w).Actors[i].Y + 3
	if (*w).Actors[i].State["Swimming"].(bool) {
		(*a).Disabled = false
	} else {
		(*a).Disabled = true
	}
}

func playerWandShoot(a *actor.Actor, w *world.World) {
	(*w).State["popup"] = true
	(*w).State["popuptimeout"] = 200
	if (*a).State["mana"].(int) > 0 {
		(*a).State["mana"] = (*a).State["mana"].(int) - 1
		spell := actor.Actor{
			Image: (*w).GetImage("spellarcane0"),
			AltImages: []*ebiten.Image{
				(*w).GetImage("spellarcane0"),
				(*w).GetImage("spellarcane1"),
				(*w).GetImage("spellarcane2"),
				(*w).GetImage("spellarcane3"),
				(*w).GetImage("spellarcane4"),
				(*w).GetImage("spellarcane5"),
				(*w).GetImage("spellarcane6"),
			},
			ActorLogic: spellArcaneActorLogic,
			Shadow:     true,
			Z:          1,
			State:      make(map[string]interface{}),
		}
		cursorx, cursory := ebiten.CursorPosition()
		spell.State["targetx"] = cursorx - (*w).CameraX
		spell.State["targety"] = cursory - (*w).CameraY
		spell.State["AnimCount"] = 0
		spell.State["Interval"] = 0
		w.SpawnActor(spell, (*a).X-32, (*a).Y-32)
	}
}

func playerAxeUse(a *actor.Actor, w *world.World) {
	cursorx, cursory := ebiten.CursorPosition()
	i, collided := w.DetectCollisionPointTag(cursorx-(*w).CameraX, cursory-(*w).CameraY, "tree")
	if collided {
		wood := actor.Actor{
			Image:      (*w).GetImage("wooditem"),
			ActorLogic: droppedItemActorLogic,
			Shadow:     true,
			Z:          0,
			State:      make(map[string]interface{}),
		}
		wood.State["targetx"] = cursorx - (*w).CameraX + 64 + (4 * (*w).Actors[i].State["health"].(int))
		wood.State["targety"] = cursory - (*w).CameraY
		wood.State["Interval"] = 0
		(*w).Actors[i].State["health"] = (*w).Actors[i].State["health"].(int) - 1
		w.SpawnActor(wood, cursorx-(*w).CameraX, cursory-(*w).CameraY-(*w).Actors[i].State["health"].(int)*8)
	}
}

func playerHotbarSwitch(a *actor.Actor, w *world.World, hotbarname string) {
	switch hotbarname {
	case "Wand":
		playerWandShoot(a, w)
	case "Iron Axe":
		playerAxeUse(a, w)
	}
}

func playerActorLogic(a *actor.Actor, w *world.World, sceneDidMove bool) {
	(*a).VelocityX = (*w).VelocityX
	(*a).VelocityY = (*w).VelocityY
	(*a).X -= int((*a).VelocityX)
	(*a).Y -= int((*a).VelocityY)
	if ebiten.IsKeyPressed(ebiten.Key1) {
		(*a).State["hotbarslot"] = 0
	} else if ebiten.IsKeyPressed(ebiten.Key2) {
		(*a).State["hotbarslot"] = 1
	} else if ebiten.IsKeyPressed(ebiten.Key3) {
		(*a).State["hotbarslot"] = 2
	}
	if (*a).State["tooltimeout"].(int) == 0 {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			(*a).State["tooltimeout"] = 1
			if (*a).State["hotbarslot"].(int) == 0 {
				playerHotbarSwitch(a, w, (*a).State["hotbar0name"].(string))
			} else if (*a).State["hotbarslot"].(int) == 1 {
				playerHotbarSwitch(a, w, (*a).State["hotbar1name"].(string))
			} else if (*a).State["hotbarslot"].(int) == 2 {
				playerHotbarSwitch(a, w, (*a).State["hotbar2name"].(string))
			}
		}
	}
	if (*a).State["tooltimeout"].(int) == 5 {
		(*a).State["tooltimeout"] = 0
	} else if (*a).State["tooltimeout"].(int) != 0 {
		(*a).State["tooltimeout"] = (*a).State["tooltimeout"].(int) + 1
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		(*a).State["health"] = (*a).State["health"].(int) - 1
	}
	if sceneDidMove {
		if a.DetectTileTouch(w, "water") && !(*a).State["Swimming"].(bool) {
			(*a).State["Swimming"] = true
			(*a).Shadow = false
			(*a).Y += 10
		} else if !a.DetectTileTouch(w, "water") && (*a).State["Swimming"].(bool) {
			(*a).State["Swimming"] = false
			(*a).Shadow = true
			(*a).Y -= 10
		}
		fsx, fsy := w.GetActorShift()
		sx, sy := int(fsx), int(fsy)
		if sx > 0 {
			if (*a).State["StrafeLeft"].(int) == 0 {
				(*a).State["StrafeLeft"] = 10
				(*a).X -= (*a).State["StrafeLeft"].(int)
			}
		} else if sx < 0 {
			if (*a).State["StrafeRight"].(int) == 0 {
				(*a).State["StrafeRight"] = 10
				(*a).X += (*a).State["StrafeRight"].(int)
			}
		} else {
			(*a).X += (*a).State["StrafeLeft"].(int)
			(*a).X -= (*a).State["StrafeRight"].(int)
			(*a).State["StrafeLeft"] = 0
			(*a).State["StrafeRight"] = 0
		}
		if sy > 0 {
			if (*a).State["StrafeUp"].(int) == 0 {
				(*a).State["StrafeUp"] = 10
				(*a).Y -= (*a).State["StrafeUp"].(int)
			}
		} else if sy < 0 {
			if (*a).State["StrafeDown"].(int) == 0 {
				(*a).State["StrafeDown"] = 10
				(*a).Y += (*a).State["StrafeDown"].(int)
			}
		} else {
			(*a).Y += (*a).State["StrafeUp"].(int)
			(*a).Y -= (*a).State["StrafeDown"].(int)
			(*a).State["StrafeUp"] = 0
			(*a).State["StrafeDown"] = 0
		}
	} else {
		resetAllStrafe(a)
	}
}
