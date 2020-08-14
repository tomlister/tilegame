package main

import (
	"fmt"
	"math/rand"

	"github.com/aquilax/go-perlin"
	"github.com/hajimehoshi/ebiten"
	"github.com/tomlister/tilegame/engine/actor"
	"github.com/tomlister/tilegame/engine/world"
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

func generateWorld(w *world.World) {
	fmt.Println("Generating world...")
	alpha := 3.0
	beta := 5.0
	n := 5
	seed := int64(69)
	p := perlin.NewPerlinRandSource(alpha, beta, n, rand.NewSource(seed))
	for y := 1; y < 201; y++ {
		for x := 1; x < 201; x++ {
			height := p.Noise2D(float64(x)/10, float64(y)/10)
			tile := actor.Actor{
				ActorLogic: backgroundActorLogic,
				Z:          -1,
				State:      make(map[string]interface{}),
			}
			if height > 0.2 {
				tile.State["world"] = true
				tile.State["imagename"] = "stone"
				tile.Tag = "stone"
				chest := actor.Actor{
					Image:      (*w).Images["chestclosed"],
					AltImages:  []*ebiten.Image{(*w).Images["chestclosed"], (*w).Images["chestopen"]},
					ActorLogic: chestActorLogic,
					Z:          0,
					State:      make(map[string]interface{}),
				}
				chest.State["Opened"] = false
				w.SpawnActorRandom(chest, (x-1)*(32), (y-1)*(32), ((x-1)*(32))+32, ((y-1)*(32))+32, 3)
			} else if height > 0 {
				tile.State["world"] = true
				tile.State["imagename"] = "grass"
				tile.Tag = "grass"
				tree := actor.Actor{
					Image:      (*w).Images["tree0"],
					AltImages:  []*ebiten.Image{(*w).Images["tree0"], (*w).Images["tree1"], (*w).Images["tree2"], (*w).Images["tree3"]},
					ActorLogic: backgroundTreeActorLogic,
					Z:          0,
					State:      make(map[string]interface{}),
				}
				tree.Tag = "tree"
				tree.State["health"] = 4
				tree.State["AnimCount"] = 0
				tree.State["Interval"] = 0
				w.SpawnActorRandom(tree, (x-1)*(32), (y-1)*(32), ((x-1)*(32))+32, ((y-1)*(32))+32, 1)
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
				tile.Image = w.GetImage(tile.State["imagename"].(string))
				w.SpawnActorRepeatSizeDefined(tile, (x-1)*(32), (y-1)*(32), 32, 32, 1, 1)
			}
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
}
