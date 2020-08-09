package main

import (
	"fmt"
	"math/rand"

	"github.com/aquilax/go-perlin"
	"github.com/cheggaaa/pb"
	"github.com/hajimehoshi/ebiten"
)

/*
	generateWorld uses a 2D Perlin noise generator to generate a height map.
	From the height map we derive the type of land the tile at (x,y) should be

	HEIGHT TILETYPE
	> 0.2  Stone
	> 0    Grass
	  0    Beach
	< 0    Water

	We also spawn in extra actors per tiletype:
		- Stone spawns chests
		- Grass spawns trees
*/

func (world *World) generateWorld() {
	fmt.Println("Generating world...")
	bar := pb.StartNew(200 * 200)
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
			if height > 0.2 {
				tile.State["world"] = true
				tile.State["imagename"] = "stone"
				tile.Tag = "stone"
				chest := Actor{
					Image:      (*world).Images["chestclosed"],
					AltImages:  []*ebiten.Image{(*world).Images["chestclosed"], (*world).Images["chestopen"]},
					ActorLogic: chestActorLogic,
					Z:          0,
					State:      make(map[string]interface{}),
				}
				chest.State["Opened"] = false
				world.spawnActorRandom(chest, (x-1)*(32), (y-1)*(32), ((x-1)*(32))+32, ((y-1)*(32))+32, 3)
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
			} else if height == 0.0 {
				tile.State["world"] = true
				tile.State["imagename"] = "beach"
				tile.Tag = "sand"
			} else {
				tile.State["world"] = true
				tile.State["imagename"] = "water"
				tile.Tag = "water"
				tile.Z = -3
			}
			if tile.State["imagename"] != nil {
				tile.Image = world.getImage(tile.State["imagename"].(string))
				world.spawnActorRepeatSizeDefined(tile, (x-1)*(32), (y-1)*(32), 32, 32, 1, 1)
			}
			bar.Increment()
		}
	}
	/*watertile := Actor{
		ActorLogic: backgroundActorLogic,
		Z:          -3,
		Static:     true,
		State:      make(map[string]interface{}),
	}
	watertile.State["world"] = true
	watertile.State["imagename"] = "water"
	watertile.Image = world.getImage(watertile.State["imagename"].(string))
	watertile.Tag = "water"
	world.spawnActorRepeatSizeDefined(watertile, 0, 0, 32, 32, 20, 15)*/
	bar.Finish()
}

/*func (world *World) generateWorldOnTheFly(x, y, w, h int) {
alpha := 3.0
beta := 5.0
n := 5
seed := int64(69)
p := perlin.NewPerlinRandSource(alpha, beta, n, rand.NewSource(seed))
for y := x; y < h; y++ {
	for x := y; x < w; x++ {
		height := p.Noise2D(float64(x)/10, float64(y)/10)*/
