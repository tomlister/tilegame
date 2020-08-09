package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

//PipelineWrapper Wraps logic, world and windowsettings
type PipelineWrapper struct {
	Logic          func(world *World)
	World          *World
	WindowSettings WindowSettings
}

//WindowSettings Stores settings for the window
type WindowSettings struct {
	Name   string
	Width  int
	Height int
}

func (pipelinewrapper PipelineWrapper) update(screen *ebiten.Image) error {
	pipelinewrapper.Logic(pipelinewrapper.World)
	if ebiten.IsDrawingSkipped() {
		//Drop frames
		println("Skipped A Frame")
		return nil
	}
	//Draw World
	viewportwidth, viewportheight := pipelinewrapper.WindowSettings.Width, pipelinewrapper.WindowSettings.Height
	//render actors
	for zpass := -3; zpass < 4; zpass++ {
		for i := 0; i < len((*pipelinewrapper.World).Actors); i++ {
			if (*pipelinewrapper.World).Actors[i].Z == zpass {
				actor := (*pipelinewrapper.World).Actors[i]
				if !actor.Disabled {
					offsetX, offsetY := actor.X+(*pipelinewrapper.World).CameraX, actor.Y+(*pipelinewrapper.World).CameraY
					if actor.Static {
						offsetX = actor.X
						offsetY = actor.Y
					}
					if !actor.Renderhook {
						imgwidth, imgheight := actor.Image.Size()
						//only render actors in viewport
						if (offsetX+imgwidth > 0 && offsetX < viewportwidth) && (offsetY < viewportheight && offsetY+imgheight > 0) {
							sopts := &ebiten.DrawImageOptions{}
							if actor.Shadow {
								sopts.ColorM.Scale(0, 0, 0, 0.5)
								r := float64(0x00)
								g := float64(0x00)
								b := float64(0x00)
								sopts.ColorM.Translate(r, g, b, 0)
								sopts.GeoM.Rotate(actor.Direction)
								//opts.GeoM.Skew(0.6, 0)
								sopts.GeoM.Translate(float64(offsetX+imgwidth/8), float64(offsetY-imgheight/8))
							}
							opts := &ebiten.DrawImageOptions{}
							opts.GeoM.Rotate(actor.Direction)
							opts.GeoM.Translate(float64(offsetX), float64(offsetY))
							screen.DrawImage(actor.Image, sopts)
							screen.DrawImage(actor.Image, opts)
						}
					} else {
						//only render actors in viewport
						actor.Rendercode(&actor, pipelinewrapper, screen)
					}
				}
			}
		}
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%d", int(ebiten.CurrentFPS())))
	if (*pipelinewrapper.World).Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("\nVx: %f, Vy: %f", (*pipelinewrapper.World).VelocityX, (*pipelinewrapper.World).VelocityY))
		ebitenutil.DebugPrint(screen, fmt.Sprintf("\n\nCx: %d, Cy: %d", (*pipelinewrapper.World).CameraX, (*pipelinewrapper.World).CameraY))
	}
	return nil
}

//StartEngine Starts the bruh engine
func StartEngine(logic func(world *World), world *World, windowsettings WindowSettings) {
	(*world).Images["missingtexture"] = importImage("assets/missing.png")
	pw := PipelineWrapper{
		Logic:          logic,
		World:          world,
		WindowSettings: windowsettings,
	}
	ebiten.SetCursorVisibility(false)
	if err := ebiten.Run(pw.update, windowsettings.Width, windowsettings.Height, 1, windowsettings.Name); err != nil {
		log.Fatal(err)
	}
}
