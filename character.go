package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/text"
)

type Attribute struct {
	Name   string
	Amount int
	Cost   int
}

//lint:ignore U1000 Stubs
func characterActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	_, yoff := ebiten.Wheel()
	(*actor).State["scrolloffset"] = (*actor).State["scrolloffset"].(float64) + yoff
	attributes := (*world).Actors[(*world).TagTable["Player"]].State["attributes"].([]Attribute)
	nohovers := 0
	for j := 0; j < len(attributes); j++ {
		if attributes[j].characterListItemLogic(actor, world, 32, int((*actor).State["scrolloffset"].(float64))+100+(j*64), j) {
			nohovers++
		}
	}
	if nohovers == len(attributes) {
		(*actor).State["hoveroffset"] = -1
	}
}

func (a *Attribute) characterListItemLogic(actor *Actor, world *World, x, y, pos int) bool {
	p := (*world).Actors[(*world).TagTable["Player"]]
	mx, my := ebiten.CursorPosition()
	rect := Rect{x, y, 330, 64}
	if detectPointRect(mx, my, rect) {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			if !(*actor).State["buttondown"].(bool) {
				(*actor).State["buttondown"] = true
				if a.canUpgrade(p) {
					a.upgradeItem(p)
				}
			}
		} else {
			(*actor).State["buttondown"] = false
		}
		if (*actor).State["hoveroffset"] != pos {
			sePlayer, _ := audio.NewPlayerFromBytes((*world).AudioContext, (*world.Sounds["hover"]))
			sePlayer.SetVolume(0.75)
			sePlayer.Play()
			(*actor).State["hoveroffset"] = pos
		}
		return false
	}
	return true
}

func (a *Attribute) upgradeItem(player Actor) {
	player.State["xp"] = player.State["xp"].(int) - a.Cost
	setAttribute(player, a.Name, a.Amount+1)
}

func characterRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	/*
		Draw background
	*/
	sx, sy := screen.Size()
	blackbg, _ := ebiten.NewImage(sx, sy, ebiten.FilterDefault)
	blackbg.Fill(color.RGBA{25, 25, 25, 0xff})
	screen.DrawImage(blackbg, &ebiten.DrawImageOptions{})
	/*
		Draw the upgrades list
	*/
	characterListRenderCode(actor, pipelinewrapper, screen)
	svo, _ := ebiten.NewImage(sx, 80, ebiten.FilterDefault)
	svo.Fill(color.RGBA{25, 25, 25, 0xff})
	screen.DrawImage(svo, &ebiten.DrawImageOptions{})
	/*
		Draw title
	*/
	text.Draw(screen, "Character", (*pipelinewrapper.World.Font[2]), 20, 50, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xff})
}

func characterListRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	attributes := (*pipelinewrapper.World).Actors[(*pipelinewrapper.World).TagTable["Player"]].State["attributes"].([]Attribute)
	for j := 0; j < len(attributes); j++ {
		attributes[j].characterListItemRenderCode(32, int((*actor).State["scrolloffset"].(float64))+100+(j*64), pipelinewrapper, screen)
	}
}

func (a *Attribute) canUpgrade(player Actor) bool {
	if player.State["xp"].(int)-a.Cost < 0 {
		return false
	}
	return true
}

func (a *Attribute) characterListItemStatusRenderCode(x, y int, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprintf("%dxp", a.Cost), (*pipelinewrapper.World.Font[0]), 342+x, y+40, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xff})
}

func (a *Attribute) characterListItemRenderCode(x, y int, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	mx, my := ebiten.CursorPosition()
	rect := Rect{x, y, 330, 64}
	if detectPointRect(mx, my, rect) {
		a.characterListItemStatusRenderCode(x, y, pipelinewrapper, screen)
		opts := &ebiten.DrawImageOptions{}
		itembg, _ := ebiten.NewImage(330, 64, ebiten.FilterDefault)
		itembg.Fill(color.RGBA{75, 75, 75, 0xff})
		opts.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(itembg, opts)
	}
	p := (*pipelinewrapper.World).Actors[(*pipelinewrapper.World).TagTable["Player"]]
	if !a.canUpgrade(p) {
		opts := &ebiten.DrawImageOptions{}
		itembg, _ := ebiten.NewImage(330, 64, ebiten.FilterDefault)
		itembg.Fill(color.RGBA{0xff, 0, 0, 0x50})
		opts.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(itembg, opts)
	}
	text.Draw(screen, fmt.Sprintf("%s - %d", a.Name, a.Amount), (*pipelinewrapper.World.Font[0]), x+10, y+32+10, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xff})
}
