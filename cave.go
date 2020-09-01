package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func caveEntryPointActorLogic(actor *Actor, world *World, sceneDidMove bool) {

}

func caveEntryPointRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
}

func caveReturnButtonActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	rect := Rect{
		actor.X, actor.Y, 64, 64,
	}
	cx, cy := ebiten.CursorPosition()
	if detectPointRect(cx, cy, rect) {
		if !(*actor).State["hovering"].(bool) {
			ebiten.SetCursorVisibility(true)
			(*actor).State["hovering"] = true
		}
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			cx, cy := (*actor).State["cx"].(int), (*actor).State["cy"].(int)
			i := (*world).TagTable["Player"]
			(*world).Actors[i].VelocityX = 0
			(*world).Actors[i].VelocityY = 0
			(*world).Actors[i].X = cx + 64
			(*world).Actors[i].Y = cy + 64
			(*world).VelocityX = 0
			(*world).VelocityY = 0
			sx, sy := (*world).Actors[i].Image.Size()
			(*world).CameraX = (-(cx + 64)) + (Width / 2) - (sx / 2)
			(*world).CameraY = (-(cy + 64)) + (Height / 2) - (sy / 2)
			for i := 0; i < len((*world).Actors); i++ {
				switch (*world).Actors[i].Tag {
				case "cavemask":
					(*world).Actors[i].Kill = true
				}
			}
			ebiten.SetCursorVisibility(false)
			(*actor).Kill = true
		}
	} else {
		if (*actor).State["hovering"].(bool) {
			ebiten.SetCursorVisibility(false)
			(*actor).State["hovering"] = false
		}
	}
}

func caveReturnButtonRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Scale(2, 2)
	ops.GeoM.Translate(float64((*actor).X), float64((*actor).Y))
	if (*actor).State["hovering"].(bool) {
		screen.DrawImage((*pipelinewrapper.World).getImage("buttonhover"), ops)
	} else {
		screen.DrawImage((*pipelinewrapper.World).getImage("button"), ops)
	}
	screen.DrawImage((*pipelinewrapper.World).getImage("overworld"), ops)
}

func caveHoleActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	col, _, _ := actor.DetectPlayerCollision(world)
	if col {
		c := (*world).TagTable["CaveEntryPoint"]
		i := (*world).TagTable["Player"]
		caveMask := Actor{
			Image:      (*world).getImage("cavemask"),
			ActorLogic: backgroundActorLogic,
			Static:     true,
			Z:          3,
			Tag:        "cavemask",
		}
		world.spawnActor(caveMask, 0, 0)
		overWorldButton := Actor{
			Renderhook: true,
			Rendercode: caveReturnButtonRenderCode,
			ActorLogic: caveReturnButtonActorLogic,
			Static:     true,
			State:      make(map[string]interface{}),
			Z:          3,
		}
		overWorldButton.State["hovering"] = false
		overWorldButton.State["cx"] = (*world).Actors[i].X
		overWorldButton.State["cy"] = (*world).Actors[i].Y
		world.spawnActor(overWorldButton, Width-84, 20)
		(*world).Actors[i].VelocityX = 0
		(*world).Actors[i].VelocityY = 0
		(*world).Actors[i].X = (*world).Actors[c].X
		(*world).Actors[i].Y = (*world).Actors[c].Y
		(*world).VelocityX = 0
		(*world).VelocityY = 0
		sx, sy := (*world).Actors[i].Image.Size()
		(*world).CameraX = (-((*world).Actors[c].X)) + (Width / 2) - (sx / 2)
		(*world).CameraY = (-((*world).Actors[c].Y)) + (Height / 2) - (sy / 2)
	}
}
