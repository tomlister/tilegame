package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
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
							if actor.Shadow {
								opts := &ebiten.DrawImageOptions{}
								opts.ColorM.Scale(0, 0, 0, 0.5)
								r := float64(0x00)
								g := float64(0x00)
								b := float64(0x00)
								opts.ColorM.Translate(r, g, b, 0)
								opts.GeoM.Rotate(actor.Direction)
								//opts.GeoM.Skew(0.6, 0)
								opts.GeoM.Translate(float64(offsetX+imgwidth/8), float64(offsetY-imgheight/8))
								screen.DrawImage(actor.Image, opts)
							}
							opts := &ebiten.DrawImageOptions{}
							opts.GeoM.Rotate(actor.Direction)
							opts.GeoM.Translate(float64(offsetX), float64(offsetY))
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
	//render text
	for i := 0; i < len((*pipelinewrapper.World).Text); i++ {
		offsetX, offsetY := (*pipelinewrapper.World).Text[i].X+(*pipelinewrapper.World).CameraX, (*pipelinewrapper.World).Text[i].Y+(*pipelinewrapper.World).CameraY
		if (offsetX+(*pipelinewrapper.World).Text[i].Width > 0 && offsetX < viewportwidth) && (offsetY < viewportheight && offsetY+20 > 0) {
			if (*pipelinewrapper.World).Text[i].Background {
				background, _ := ebiten.NewImage((*pipelinewrapper.World).Text[i].Width, 20, ebiten.FilterDefault)
				opts := &ebiten.DrawImageOptions{}
				opts.ColorM.Scale(0, 0, 0, 0.1)
				r := float64(0xff)
				g := float64(0xff)
				b := float64(0xff)
				opts.ColorM.Translate(r, g, b, 1)
				opts.GeoM.Translate(float64(offsetX), float64(offsetY-16))
				screen.DrawImage(background, opts)
				background.Dispose()
			}
			text.Draw(screen, (*pipelinewrapper.World).Text[i].Text, (*pipelinewrapper.World.Font), offsetX, offsetY, color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff})
		}
	}
	(*pipelinewrapper.World).Text = []Text{}

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
