package main

import (
	"fmt"
	"math/rand"

	"github.com/aquilax/go-perlin"
	"github.com/hajimehoshi/ebiten"
)

/*
	generateWorld uses a 2D Perlin noise generator to generate a height map.
	From the height map we derive the type of land the tile at (x,y) should be

	HEIGHT TILETYPE
	> 0.2  Stone
	> 0    Grass
	< 0    Water

	We also spawn in extra actors per tiletype:
		- Stone spawns chests
		- Grass spawns trees
*/

func (world *World) generateWorld() {
	fmt.Println("Generating world...")
	alpha := 3.0
	beta := 5.0
	n := 5
	seed := int64(69)
	p := perlin.NewPerlinRandSource(alpha, beta, n, rand.NewSource(seed))
	for y := 1; y < 201; y++ {
		for x := 1; x < 201; x++ {
			height := p.Noise2D(float64(x)/10, float64(y)/10)
			tile := Actor{
				ActorLogic: backgroundActorLogic,
				Z:          -1,
				State:      make(map[string]interface{}),
			}
			if height > 0.4 {
				tile.State["world"] = true
				tile.State["imagename"] = "stone"
				tile.Tag = "stone"
				rock := Actor{
					Tag:        "rock",
					Image:      (*world).Images["rock"],
					ActorLogic: backgroundRockActorLogic,
					Z:          0,
					State:      make(map[string]interface{}),
				}
				rock.State["health"] = 2
				world.spawnActorRandom(rock, (x-1)*(32), (y-1)*(32), ((x-1)*(32))+32, ((y-1)*(32))+32, 3)
				chest := Actor{
					Image:      (*world).Images["chestclosed"],
					AltImages:  []*ebiten.Image{(*world).Images["chestclosed"], (*world).Images["chestopen"]},
					ActorLogic: chestActorLogic,
					Z:          0,
					State:      make(map[string]interface{}),
				}
				chest.State["Opened"] = false
				world.spawnActorRandom(chest, (x-1)*(32), (y-1)*(32), ((x-1)*(32))+32, ((y-1)*(32))+32, 1)
				if height > 0.65 {
					trader := Actor{
						Image:      (*world).Images["trader"],
						ActorLogic: traderActorLogic,
						Z:          0,
						State:      make(map[string]interface{}),
					}
					trader.State["inspeech"] = false
					world.spawnActorRandom(trader, (x-1)*(32), (y-1)*(32), ((x-1)*(32))+32, ((y-1)*(32))+32, 1)
				}
			} else if height > 0 {
				tile.State["world"] = true
				tile.State["imagename"] = "grass"
				tile.Tag = "grass"
				tree := Actor{
					Image:      (*world).Images["tree0"],
					AltImages:  []*ebiten.Image{(*world).Images["tree0"], (*world).Images["tree1"], (*world).Images["tree2"], (*world).Images["tree3"]},
					ActorLogic: backgroundTreeActorLogic,
					Z:          0,
					State:      make(map[string]interface{}),
				}
				tree.Tag = "tree"
				tree.State["health"] = 4
				tree.State["AnimCount"] = 0
				tree.State["Interval"] = 0
				world.spawnActorRandom(tree, (x-1)*(32), (y-1)*(32), ((x-1)*(32))+32, ((y-1)*(32))+32, 1)
			} else {
				tile.State["world"] = true
				if p.Noise2D(float64(x)/10, float64(y-1)/10) > 0 {
					tile.State["imagename"] = "wateredgeS"
				} else if p.Noise2D(float64(x)/10, float64(y+1)/10) > 0 {
					tile.State["imagename"] = "wateredgeN"
				} else if p.Noise2D(float64(x-1)/10, float64(y)/10) > 0 {
					tile.State["imagename"] = "wateredgeE"
				} else if p.Noise2D(float64(x+1)/10, float64(y)/10) > 0 {
					tile.State["imagename"] = "wateredgeW"
				} else if p.Noise2D(float64(x-1)/10, float64(y-1)/10) > 0 {
					tile.State["imagename"] = "wateredgeSE"
				} else if p.Noise2D(float64(x-1)/10, float64(y+1)/10) > 0 {
					tile.State["imagename"] = "wateredgeNE"
				} else if p.Noise2D(float64(x+1)/10, float64(y-1)/10) > 0 {
					tile.State["imagename"] = "wateredgeSW"
				} else if p.Noise2D(float64(x+1)/10, float64(y+1)/10) > 0 {
					tile.State["imagename"] = "wateredgeNW"
				} else {
					tile.State["imagename"] = "water"
				}
				tile.Tag = "water"
				tile.Z = -3
			}
			if tile.State["imagename"] != nil {
				tile.Image = world.getImage(tile.State["imagename"].(string))
				world.spawnActorRepeatSizeDefined(tile, (x-1)*(32), (y-1)*(32), 32, 32, 1, 1)
			}
		}
	}
}

/*
	generateDungeonWorld generates the dungeon world
*/

func (world *World) generateDungeonWorld() {
	fmt.Println("Generating dungeon world...")
	alpha := 3.0
	beta := 5.0
	n := 5
	seed := int64(69 * 4)
	p := perlin.NewPerlinRandSource(alpha, beta, n, rand.NewSource(seed))
	for y := 1; y < 51; y++ {
		for x := 1; x < 51; x++ {
			height := p.Noise2D(float64(x)/10, float64(y)/10)
			tile := Actor{
				ActorLogic: backgroundActorLogic,
				Z:          -1,
				State:      make(map[string]interface{}),
			}
			if height >= 0 {
				tile.State["world"] = true
				tile.State["imagename"] = "cavewall"
				tile.Tag = "cavewall"
			} else if height < 0 {
				tile.State["world"] = true
				tile.State["imagename"] = "cavefloor"
				tile.Tag = "cavefloor"
				chest := Actor{
					Image:      (*world).Images["chestclosed"],
					AltImages:  []*ebiten.Image{(*world).Images["chestclosed"], (*world).Images["chestopen"]},
					ActorLogic: chestActorLogic,
					Z:          0,
					State:      make(map[string]interface{}),
				}
				chest.State["Opened"] = false
			}
			if tile.State["imagename"] != nil {
				tile.Image = world.getImage(tile.State["imagename"].(string))
				world.spawnActorRepeatSizeDefined(tile, (-x-1)*(32), (-y-1)*(32), 32, 32, 1, 1)
			}
		}
	}
	tpsearch := true
	for y := 1; y < 51; y++ {
		for x := 1; x < 51; x++ {
			height := p.Noise2D(float64(x)/10, float64(y)/10)
			tile := Actor{
				ActorLogic: backgroundActorLogic,
				Z:          -1,
				State:      make(map[string]interface{}),
			}
			if height < 0 {
				tile.State["world"] = true
				if p.Noise2D(float64(x)/10, float64(y+1)/10) >= 0 {
					tile.State["imagename"] = "cavewallS"
				} else if p.Noise2D(float64(x-1)/10, float64(y)/10) >= 0 {
					tile.State["imagename"] = "cavewallW"
				} else if p.Noise2D(float64(x)/10, float64(y-1)/10) >= 0 {
					tile.State["imagename"] = "cavewallN"
				} else if p.Noise2D(float64(x+1)/10, float64(y)/10) >= 0 {
					tile.State["imagename"] = "cavewallE"
				}
				if p.Noise2D(float64(x+1)/10, float64(y)/10) >= 0 {
					if p.Noise2D(float64(x)/10, float64(y+1)/10) >= 0 {
						if p.Noise2D(float64(x)/10, float64(y-1)/10) >= 0 {
							tile.State["imagename"] = "cavewallEFullCorner"
						}
					}
				}
				tile.Tag = "cavewall"
			}
			if tile.State["imagename"] != nil {
				tile.Image = world.getImage(tile.State["imagename"].(string))
				world.spawnActorRepeatSizeDefined(tile, (-x-1)*(32), (-y-1)*(32), 32, 32, 1, 1)
			}
			if tpsearch == true {
				if p.Noise2D(float64(x)/10, float64(y+1)/10) < 0 {
					if p.Noise2D(float64(x-1)/10, float64(y)/10) < 0 {
						if p.Noise2D(float64(x)/10, float64(y-1)/10) < 0 {
							if p.Noise2D(float64(x+1)/10, float64(y)/10) < 0 {
								caveEntry := Actor{
									Tag:        "CaveEntryPoint",
									ActorLogic: backgroundActorLogic,
									Rendercode: backgroundActorRenderLogic,
									Renderhook: true,
								}
								world.spawnActor(caveEntry, (-x-1)*(32), (-y-1)*(32))
								tpsearch = false
							}
						}
					}
				}
			}
		}
	}
}
