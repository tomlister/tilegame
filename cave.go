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
