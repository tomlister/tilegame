package main

import (
	"math"
	"math/rand"

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
			(*world).Actors[i].CustomRenderDestination = false
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
	i := (*world).TagTable["Player"]
	rect := Rect{actor.X, actor.Y, 32, 32}
	if detectPointRect((*world).Actors[i].X+16, (*world).Actors[i].Y+16, rect) {
		c := (*world).TagTable["CaveEntryPoint"]
		caveLighting := Actor{
			Renderhook: true,
			Rendercode: caveLightingRenderCode,
			ActorLogic: backgroundActorLogic,
			Static:     true,
			Z:          2,
			Tag:        "cavemask",
			State:      make(map[string]interface{}),
		}
		caveLighting.State["time"] = 0.0
		world.spawnActor(caveLighting, 0, 0)
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
		(*world).Actors[i].RenderDestination = (*world).getImage("offscreen")
		(*world).Actors[i].CustomRenderDestination = true
	}
}

func caveLightingRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	(*actor).State["time"] = (*actor).State["time"].(float64) + 1
	s := (*pipelinewrapper.World).Shaders["lighting"]
	op := &ebiten.DrawRectShaderOptions{}
	lightcolors := []float32{}      // 3 floats per light
	lightpositions := []float32{}   // 2 floats per light
	lightintensities := []float32{} // 40 max
	lightamount := float32(0.0)     // 40 max
	viewportwidth, viewportheight := pipelinewrapper.WindowSettings.Width, pipelinewrapper.WindowSettings.Height
	for _, a := range (*pipelinewrapper.World).Actors {
		if a.Tag == "manacrystal" || a.Tag == "manacrystaldropped" {
			offsetX, offsetY := a.X+(*pipelinewrapper.World).CameraX, a.Y+(*pipelinewrapper.World).CameraY
			imgwidth, imgheight := a.Image.Size()
			if (offsetX+imgwidth > 0 && offsetX < viewportwidth) && (offsetY < viewportheight && offsetY+imgheight > 0) {
				lightamount++
				color := []float32{0xDA, 0x96, 0xF7}
				lightcolors = append(lightcolors, color...)
				position := []float32{float32(offsetX) + 16, float32(offsetY) + 16}
				lightpositions = append(lightpositions, position...)
				lightintensities = append(lightintensities, float32(math.Sin((2*rand.Float64()+(*actor).State["time"].(float64))/100)*5000000))
			}
		}
		if lightamount == 40 {
			break
		}
	}
	if (len(lightcolors)/3)-int(40) != 0 {
		lightcolors = append(lightcolors, make([]float32, int(120-(lightamount*3)))...)
	}
	if (len(lightpositions)/2)-int(40) != 0 {
		lightpositions = append(lightpositions, make([]float32, int(80-(lightamount*2)))...)
	}
	if len(lightintensities)-int(40) != 0 {
		lightintensities = append(lightintensities, make([]float32, int(40-(lightamount)))...)
	}
	op.Uniforms = []interface{}{
		lightcolors,            // LightColors
		lightintensities,       // LightIntensities
		lightpositions,         // LightPositions
		[]float32{lightamount}, // LightAmount
	}
	offscreen := (*pipelinewrapper.World).getImage("offscreen")
	osw, osh := offscreen.Size()
	op.Images[0] = offscreen
	screen.DrawRectShader(osw, osh, s, op)
}
