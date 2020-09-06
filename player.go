package main

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
)

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
}

func playerSplashActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	i := (*world).TagTable["Player"]
	(*actor).X = (*world).Actors[i].X
	(*actor).Y = (*world).Actors[i].Y + 3
	if (*world).Actors[i].State["Swimming"].(bool) {
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
		wood.State["item"] = Item{
			Name:      "Wood",
			ImageName: "wooditem",
			Quantity:  1,
		}
		(*world).Actors[i].State["health"] = (*world).Actors[i].State["health"].(int) - 1
		world.spawnActor(wood, cursorx-(*world).CameraX, cursory-(*world).CameraY-(*world).Actors[i].State["health"].(int)*8)
	}
	i, collided = world.detectCollisionPointTag(cursorx-(*world).CameraX, cursory-(*world).CameraY, "rock")
	if collided {
		ironpowder := Actor{
			Image:      (*world).getImage("ironpowderitem"),
			ActorLogic: droppedItemActorLogic,
			Shadow:     true,
			Z:          0,
			State:      make(map[string]interface{}),
		}
		ironpowder.State["targetx"] = cursorx - (*world).CameraX + 64 + (4 * (*world).Actors[i].State["health"].(int))
		ironpowder.State["targety"] = cursory - (*world).CameraY
		ironpowder.State["Interval"] = 0
		ironpowder.State["item"] = Item{
			Name:      "Iron Powder",
			ImageName: "ironpowderitem",
			Quantity:  1,
		}
		(*world).Actors[i].State["health"] = (*world).Actors[i].State["health"].(int) - 1
		world.spawnActor(ironpowder, cursorx-(*world).CameraX, cursory-(*world).CameraY-(*world).Actors[i].State["health"].(int)*8)
	}
	i, collided = world.detectCollisionPointTag(cursorx-(*world).CameraX, cursory-(*world).CameraY, "manacrystal")
	if collided {
		manaCrystal := Actor{
			Image:                   (*world).getImage("manacrystal"),
			Tag:                     "manacrystaldropped",
			ActorLogic:              droppedItemActorLogic,
			Shadow:                  true,
			Z:                       0,
			State:                   make(map[string]interface{}),
			RenderDestination:       (*world).getImage("offscreen"),
			CustomRenderDestination: true,
		}
		manaCrystal.State["targetx"] = cursorx - (*world).CameraX + 64 + (4 * (*world).Actors[i].State["health"].(int))
		manaCrystal.State["targety"] = cursory - (*world).CameraY
		manaCrystal.State["Interval"] = 0
		manaCrystal.State["item"] = Item{
			Name:      "Mana Crystal",
			ImageName: "manacrystal",
			Quantity:  1,
		}
		(*world).Actors[i].State["health"] = (*world).Actors[i].State["health"].(int) - 1
		world.spawnActor(manaCrystal, cursorx-(*world).CameraX, cursory-(*world).CameraY-(*world).Actors[i].State["health"].(int)*8)
	}
}

func knockBack(px, py, ex, ey int) (float64, float64) {
	theta := math.Atan2(float64(py-ey), float64(px-ex))
	knockbackDistance := 10.0
	x := knockbackDistance * math.Cos(theta)
	y := knockbackDistance * math.Sin(theta)
	return x, y
}

func playerSwordUse(actor *Actor, world *World) {
	cursorx, cursory := ebiten.CursorPosition()
	i, collided := world.detectCollisionPointTag(cursorx-(*world).CameraX, cursory-(*world).CameraY, "enemy")
	if collided {
		vx, vy := knockBack(actor.X, actor.Y, (*world).Actors[i].X, (*world).Actors[i].Y)
		(*world).Actors[i].VelocityX += vx
		(*world).Actors[i].VelocityY += vy
		profile := (*world).Actors[i].State["profile"].(Enemy)
		profile.Health = profile.Health - 1
		(*world).Actors[i].State["profile"] = profile
		minusone := Actor{
			Image:      (*world).getImage("minusone"),
			ActorLogic: floaterActorLogic,
			State:      make(map[string]interface{}),
		}
		minusone.State["Interval"] = 0
		minusone.State["AnimCount"] = 0
		world.spawnActor(minusone, (*world).Actors[i].X, (*world).Actors[i].Y-16)
		soundnames := []string{"hit1", "hit2", "hit3", "hit4"}
		soundindex := rand.Intn(len(soundnames))
		soundname := soundnames[soundindex]
		sePlayer, _ := audio.NewPlayerFromBytes((*world).AudioContext, (*world.Sounds[soundname]))
		sePlayer.Play()
	} else {
		soundnames := []string{"sword1", "sword2", "sword3", "sword4"}
		soundindex := rand.Intn(len(soundnames))
		soundname := soundnames[soundindex]
		sePlayer, _ := audio.NewPlayerFromBytes((*world).AudioContext, (*world.Sounds[soundname]))
		sePlayer.Play()
	}
}

func playerManaPotionUse(actor *Actor, world *World) {
	hotbar := (*actor).State["hotbar"].(Hotbar)
	hotbar.Slots[(*actor).State["hotbar"].(Hotbar).Slot] = Item{}
	(*actor).State["hotbar"] = hotbar
	addmana := (*actor).State["mana"].(int) + (*actor).State["manamax"].(int)/4
	if addmana > (*actor).State["manamax"].(int) {
		(*actor).State["mana"] = (*actor).State["manamax"].(int)
	} else {
		(*actor).State["mana"] = addmana
	}
}

func playerHotbarSwitch(actor *Actor, world *World, hotbarname string) {
	switch hotbarname {
	case "Wand":
		playerWandShoot(actor, world)
	case "Iron Axe":
		playerAxeUse(actor, world)
	case "Wooden Axe":
		playerAxeUse(actor, world)
	case "Wooden Sword":
		playerSwordUse(actor, world)
	case "Iron Sword":
		playerSwordUse(actor, world)
	case "Mana Potion":
		playerManaPotionUse(actor, world)
	}
}

func playerActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	(*actor).VelocityX = (*world).VelocityX
	(*actor).VelocityY = (*world).VelocityY
	(*actor).X -= int((*actor).VelocityX)
	(*actor).Y -= int((*actor).VelocityY)
	if ebiten.IsKeyPressed(ebiten.Key1) {
		hotbar := (*actor).State["hotbar"].(Hotbar)
		hotbar.Slot = 0
		(*actor).State["hotbar"] = hotbar
	} else if ebiten.IsKeyPressed(ebiten.Key2) {
		hotbar := (*actor).State["hotbar"].(Hotbar)
		hotbar.Slot = 1
		(*actor).State["hotbar"] = hotbar
	} else if ebiten.IsKeyPressed(ebiten.Key3) {
		hotbar := (*actor).State["hotbar"].(Hotbar)
		hotbar.Slot = 2
		(*actor).State["hotbar"] = hotbar
	}
	if (*actor).State["tooltimeout"].(int) == 0 {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			(*actor).State["tooltimeout"] = 1
			playerHotbarSwitch(actor, world, (*actor).State["hotbar"].(Hotbar).Slots[(*actor).State["hotbar"].(Hotbar).Slot].Name)
		}
	}
	if (*actor).State["tooltimeout"].(int) == 10 {
		(*actor).State["tooltimeout"] = 0
	} else if (*actor).State["tooltimeout"].(int) != 0 {
		(*actor).State["tooltimeout"] = (*actor).State["tooltimeout"].(int) + 1
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		(*actor).State["health"] = (*actor).State["health"].(int) - 1
	}
}
