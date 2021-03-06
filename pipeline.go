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
		return nil
	}
	currentlyRendering := 0
	//Draw World
	viewportwidth, viewportheight := pipelinewrapper.WindowSettings.Width, pipelinewrapper.WindowSettings.Height
	//render actors
	for zpass := -3; zpass < 4; zpass++ {
		for _, actor := range (*pipelinewrapper.World).Actors {
			if actor.Z == zpass {
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
							currentlyRendering++
							renderDst := screen
							if actor.CustomRenderDestination {
								renderDst = actor.RenderDestination
							}
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
							if actor.Shadow {
								renderDst.DrawImage(actor.Image, sopts)
							}
							renderDst.DrawImage(actor.Image, opts)
						}
					} else {
						currentlyRendering++
						renderDst := screen
						if actor.CustomRenderDestination {
							renderDst = actor.RenderDestination
						}
						actor.Rendercode(&actor, pipelinewrapper, renderDst)
					}
				}
			}
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("%d", int(ebiten.CurrentFPS())))
	if (*pipelinewrapper.World).Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("\nVx: %f, Vy: %f", (*pipelinewrapper.World).VelocityX, (*pipelinewrapper.World).VelocityY))
		ebitenutil.DebugPrint(screen, fmt.Sprintf("\n\nCx: %d, Cy: %d", (*pipelinewrapper.World).CameraX, (*pipelinewrapper.World).CameraY))
		ebitenutil.DebugPrint(screen, fmt.Sprintf("\n\n\nActors: %d, Rendering: %d", len((*pipelinewrapper.World).Actors), currentlyRendering))
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
	if err := ebiten.Run(pw.update, windowsettings.Width, windowsettings.Height, 1, windowsettings.Name); err != nil {
		log.Fatal(err)
	}
}
