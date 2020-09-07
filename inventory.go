package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

type Item struct {
	Name      string
	ImageName string
	Quantity  int
}

type Hotbar struct {
	Slot  int
	Slots [3]Item
}

type ItemMove struct {
	SrcSlot int
	SrcType string
	DstSlot int
	DstType string
}

//lint:ignore U1000 Stubs
func inventoryActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	inventoryHotbarLogicCode(actor, world)
	inventoryGridLogicCode(actor, world)
}

func inventoryRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	/*
		Draw background
	*/
	sx, sy := screen.Size()
	blackbg, _ := ebiten.NewImage(sx, sy, ebiten.FilterDefault)
	blackbg.Fill(color.RGBA{25, 25, 25, 0xff})
	screen.DrawImage(blackbg, &ebiten.DrawImageOptions{})
	/*
		Draw title
	*/
	text.Draw(screen, "Inventory", (*pipelinewrapper.World.Font[2]), 20, 50, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xff})
	/*
		Draw the inventory grid
	*/
	inventoryGridRenderCode(pipelinewrapper, screen, actor)
	/*
		Draw the hotbar
	*/
	inventoryHotbarRenderCode(actor, pipelinewrapper, screen)
}

func inventoryHotbarLogicCode(actor *Actor, world *World) {
	i := (*world).TagTable["Player"]
	slots := (*world).Actors[i].State["hotbar"].(Hotbar).Slots
	x, y := 0, 0
	for j := 0; j < len(slots); j++ {
		slots[j].inventoryGridItemLogicCode(160+(x*128), 80+(y*64), actor, world, true, j)
		if x == 2 {
			x = 0
			y++
		} else {
			x++
		}
	}
}

func inventoryGridLogicCode(actor *Actor, world *World) {
	i := (*world).TagTable["Player"]
	inventory := (*world).Actors[i].State["inventory"].([9]Item)
	x, y := 0, 0
	for j := 0; j < len(inventory); j++ {
		inventory[j].inventoryGridItemLogicCode(160+(x*128), 160+(y*64), actor, world, false, j)
		if x == 2 {
			x = 0
			y++
		} else {
			x++
		}
	}
}

func (i *Item) inventoryGridItemLogicCode(x, y int, actor *Actor, world *World, hotbar bool, slot int) {
	mx, my := ebiten.CursorPosition()
	rect := Rect{x, y, 64, 64}
	if detectPointRect(mx, my, rect) {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			if (*actor).State["mousedown"].(bool) == false {
				(*actor).State["mousedown"] = true
				if (*actor).State["move"] == nil {
					srctype := "hotbar"
					if !hotbar {
						srctype = "inventory"
					}
					(*actor).State["move"] = ItemMove{
						SrcSlot: slot,
						SrcType: srctype,
					}
				} else {
					dsttype := "hotbar"
					if !hotbar {
						dsttype = "inventory"
					}
					i := (*world).TagTable["Player"]
					dstitem := Item{}
					if dsttype == "hotbar" {
						dstitem = (*world).Actors[i].State["hotbar"].(Hotbar).Slots[slot]
					} else if dsttype == "inventory" {
						dstitem = (*world).Actors[i].State["inventory"].([9]Item)[slot]
					}
					srcitem := Item{}
					if (*actor).State["move"].(ItemMove).SrcType == "hotbar" {
						srcitem = (*world).Actors[i].State["hotbar"].(Hotbar).Slots[(*actor).State["move"].(ItemMove).SrcSlot]
					} else if (*actor).State["move"].(ItemMove).SrcType == "inventory" {
						srcitem = (*world).Actors[i].State["inventory"].([9]Item)[(*actor).State["move"].(ItemMove).SrcSlot]
					}
					if dsttype == "inventory" {
						inv := (*world).Actors[i].State["inventory"].([9]Item)
						inv[slot] = srcitem
						(*world).Actors[i].State["inventory"] = inv
					} else if dsttype == "hotbar" {
						hotbar := (*world).Actors[i].State["hotbar"].(Hotbar)
						hotbar.Slots[slot] = srcitem
						(*world).Actors[i].State["hotbar"] = hotbar
					}
					if (*actor).State["move"].(ItemMove).SrcType == "inventory" {
						inv := (*world).Actors[i].State["inventory"].([9]Item)
						inv[(*actor).State["move"].(ItemMove).SrcSlot] = dstitem
						(*world).Actors[i].State["inventory"] = inv
					} else if (*actor).State["move"].(ItemMove).SrcType == "hotbar" {
						hotbar := (*world).Actors[i].State["hotbar"].(Hotbar)
						hotbar.Slots[(*actor).State["move"].(ItemMove).SrcSlot] = dstitem
						(*world).Actors[i].State["hotbar"] = hotbar
					}
					(*actor).State["move"] = nil
				}
			}
		} else {
			(*actor).State["mousedown"] = false
		}
	}
}

func inventoryHotbarRenderCode(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image) {
	i := (*pipelinewrapper.World).TagTable["Player"]
	slots := (*pipelinewrapper.World).Actors[i].State["hotbar"].(Hotbar).Slots
	x, y := 0, 0
	for j := 0; j < len(slots); j++ {
		slots[j].inventoryGridItemRenderCode(160+(x*128), 80+(y*64), pipelinewrapper, screen, actor, true, j)
		if x == 2 {
			x = 0
			y++
		} else {
			x++
		}
	}
}

func inventoryGridRenderCode(pipelinewrapper PipelineWrapper, screen *ebiten.Image, actor *Actor) {
	i := (*pipelinewrapper.World).TagTable["Player"]
	inventory := (*pipelinewrapper.World).Actors[i].State["inventory"].([9]Item)
	x, y := 0, 0
	for j := 0; j < len(inventory); j++ {
		inventory[j].inventoryGridItemRenderCode(160+(x*128), 160+(y*64), pipelinewrapper, screen, actor, false, j)
		if x == 2 {
			x = 0
			y++
		} else {
			x++
		}
	}
}

func (i *Item) inventoryGridItemRenderCode(x, y int, pipelinewrapper PipelineWrapper, screen *ebiten.Image, actor *Actor, hotbar bool, slot int) {
	mx, my := ebiten.CursorPosition()
	rect := Rect{x, y, 64, 64}
	opts := &ebiten.DrawImageOptions{}
	griditembg, _ := ebiten.NewImage(64, 64, ebiten.FilterDefault)
	griditembg.Fill(color.RGBA{50, 50, 50, 0xff})
	opts.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(griditembg, opts)
	if detectPointRect(mx, my, rect) {
		opts = &ebiten.DrawImageOptions{}
		itembg, _ := ebiten.NewImage(64, 64, ebiten.FilterDefault)
		itembg.Fill(color.RGBA{75, 75, 75, 0xff})
		opts.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(itembg, opts)
	}
	if (*actor).State["move"] != nil {
		if hotbar && (*actor).State["move"].(ItemMove).SrcType == "hotbar" {
			if (*actor).State["move"].(ItemMove).SrcSlot == slot {
				opts := &ebiten.DrawImageOptions{}
				itembg, _ := ebiten.NewImage(64, 64, ebiten.FilterDefault)
				itembg.Fill(color.RGBA{00, 75, 00, 0x50})
				opts.GeoM.Translate(float64(x), float64(y))
				screen.DrawImage(itembg, opts)
			}
		} else if !hotbar && (*actor).State["move"].(ItemMove).SrcType == "inventory" {
			if (*actor).State["move"].(ItemMove).SrcSlot == slot {
				opts := &ebiten.DrawImageOptions{}
				itembg, _ := ebiten.NewImage(64, 64, ebiten.FilterDefault)
				itembg.Fill(color.RGBA{00, 75, 00, 0x50})
				opts.GeoM.Translate(float64(x), float64(y))
				screen.DrawImage(itembg, opts)
			}
		}
	}
	opts = &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(2, 2)
	opts.GeoM.Translate(float64(x), float64(y))
	if (*i).ImageName != "" {
		screen.DrawImage((*pipelinewrapper.World).getImage((*i).ImageName), opts)
		text.Draw(screen, fmt.Sprintf("%d", (*i).Quantity), (*pipelinewrapper.World.Font[0]), x+32, y+60, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xff})
	}
}
