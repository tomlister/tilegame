package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/text"
)

type Craftable struct {
	Item
	Needs    []Item
	Quantity int
}

//lint:ignore U1000 Stubs
func craftingActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	_, yoff := ebiten.Wheel()
	(*actor).State["scrolloffset"] = (*actor).State["scrolloffset"].(float64) + yoff
	inventory := (*world).State["craftable"].([]Craftable)
	nohovers := 0
	for j := 0; j < len(inventory); j++ {
		if inventory[j].craftingListItemLogic(actor, world, 32, int((*actor).State["scrolloffset"].(float64))+100+(j*64), j) {
			nohovers++
		}
	}
	if nohovers == len(inventory) {
		(*actor).State["hoveroffset"] = -1
	}
}

func (i *Craftable) craftingListItemLogic(actor *Actor, world *World, x, y, pos int) bool {
	p := (*world).Actors[(*world).TagTable["Player"]]
	mx, my := ebiten.CursorPosition()
	rect := Rect{x, y, 330, 64}
	if detectPointRect(mx, my, rect) {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			if !(*actor).State["buttondown"].(bool) {
				(*actor).State["buttondown"] = true
				if i.canCraft(p) {
					i.craftItem(p)
				}
			}
		} else {
			(*actor).State["buttondown"] = false
		}
		if (*actor).State["hoveroffset"] != pos {
			sePlayer, _ := audio.NewPlayerFromBytes((*world).AudioContext, (*world.Sounds["hover"]))
			sePlayer.SetVolume(0.2)
			sePlayer.Play()
			(*actor).State["hoveroffset"] = pos
		}
		return false
	} else {
		return true
	}
}

func (i *Craftable) craftItem(player Actor) {
	inventory := player.State["inventory"].([9]Item)
	for _, need := range (*i).Needs {
		for i, item := range inventory {
			if item.Name == need.Name {
				if item.Quantity >= need.Quantity {
					if player.State["inventory"].([9]Item)[i].Quantity-need.Quantity == 0 {
						inv := player.State["inventory"].([9]Item)
						inv[i] = Item{}
						player.State["inventory"] = inv
					} else {
						inv := player.State["inventory"].([9]Item)
						inv[i].Quantity = player.State["inventory"].([9]Item)[i].Quantity - need.Quantity
						player.State["inventory"] = inv
					}
				}
			}
		}
	}
	inv := player.State["inventory"].([9]Item)
	for j := 0; j < len(inv); j++ {
		if inv[j].ImageName == "" {
			inv[j] = (*i).Item
			break
		}
	}
	player.State["inventory"] = inv
}

func craftingRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	/*
		Create offscreen image
	*/
	//offscreen, _ := ebiten.NewImage(Width, Height, ebiten.FilterDefault)
	/*
		Draw background
	*/
	sx, sy := screen.Size()
	blackbg, _ := ebiten.NewImage(sx, sy, ebiten.FilterDefault)
	blackbg.Fill(color.RGBA{25, 25, 25, 0xff})
	screen.DrawImage(blackbg, &ebiten.DrawImageOptions{})
	/*
		Draw the crafting list
	*/
	craftingListRenderCode(actor, pipelinewrapper, screen)
	/*
		Draw scrollview occluder
	*/
	/*s := (*pipelinewrapper.World).Shaders["blur"]
	op := &ebiten.DrawRectShaderOptions{}
	op.Uniforms = []interface{}{
		[]float32{float32(sx), float32(80), float32(2)}, // Blur w,h,size
	}*/
	//offscreensub := offscreen.SubImage(image.Rect(0, 0, sx, 80))
	//offscreenfin, _ := ebiten.NewImageFromImage(offscreensub, ebiten.FilterDefault)
	/*d := (*pipelinewrapper.World).getImage("titlenormal")
	dw, dh := d.Size()
	op.Images[0] = d
	screen.DrawRectShader(dw, dh, s, op)*/
	svo, _ := ebiten.NewImage(sx, 80, ebiten.FilterDefault)
	svo.Fill(color.RGBA{25, 25, 25, 0xff})
	screen.DrawImage(svo, &ebiten.DrawImageOptions{})
	/*
		Draw title
	*/
	text.Draw(screen, "Crafting", (*pipelinewrapper.World.Font[2]), 20, 50, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xff})
}

func craftingListRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	inventory := (*pipelinewrapper.World).State["craftable"].([]Craftable)
	for j := 0; j < len(inventory); j++ {
		inventory[j].craftingListItemRenderCode(32, int((*actor).State["scrolloffset"].(float64))+100+(j*64), pipelinewrapper, screen)
	}
}

func (i *Craftable) canCraft(player Actor) bool {
	inventory := player.State["inventory"].([9]Item)
	needs := 0
	fulfilled := 0
	for _, need := range (*i).Needs {
		needs++
		for _, item := range inventory {
			if item.Name == need.Name {
				if item.Quantity >= need.Quantity {
					fulfilled++
				}
			}
		}
	}
	if fulfilled >= needs {
		return true
	}
	return false
}

func (i *Craftable) craftingRequirementsRenderCode(x, y int, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	c := 0
	for _, item := range (*i).Needs {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Scale(2, 2)
		opts.GeoM.Translate(float64(342+x+(64*c)), float64(100))
		screen.DrawImage((*pipelinewrapper.World).getImage(item.ImageName), opts)
		c++
	}
}

func (i *Craftable) craftingListItemRenderCode(x, y int, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	mx, my := ebiten.CursorPosition()
	rect := Rect{x, y, 330, 64}
	if detectPointRect(mx, my, rect) {
		i.craftingRequirementsRenderCode(x, y, pipelinewrapper, screen)
		opts := &ebiten.DrawImageOptions{}
		itembg, _ := ebiten.NewImage(330, 64, ebiten.FilterDefault)
		itembg.Fill(color.RGBA{75, 75, 75, 0xff})
		opts.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(itembg, opts)
	}
	p := (*pipelinewrapper.World).Actors[(*pipelinewrapper.World).TagTable["Player"]]
	if !i.canCraft(p) {
		opts := &ebiten.DrawImageOptions{}
		itembg, _ := ebiten.NewImage(330, 64, ebiten.FilterDefault)
		itembg.Fill(color.RGBA{0xff, 0, 0, 0x50})
		opts.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(itembg, opts)
	}
	text.Draw(screen, i.Item.Name, (*pipelinewrapper.World.Font[0]), x+100, y+32+10, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xff})
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(2, 2)
	opts.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage((*pipelinewrapper.World).getImage((*i).ImageName), opts)
}
