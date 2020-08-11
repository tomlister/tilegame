package main

import (
	"github.com/hajimehoshi/ebiten"
)

func titleActorLogic(actor *Actor, world *World, sceneDidMove bool) {
}

func titleRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	s := (*pipelinewrapper.World).Shaders["title"]
	w, h := screen.Size()
	cx, cy := ebiten.CursorPosition()
	op := &ebiten.DrawRectShaderOptions{}
	op.Uniforms = []interface{}{
		float32(g.time) / 60,                // Time
		[]float32{float32(cx), float32(cy)}, // Cursor
		[]float32{float32(w), float32(h)},   // ScreenSize
	}
	screen.DrawRectShader(w, h, s, op)
}
