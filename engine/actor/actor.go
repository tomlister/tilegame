package engine

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/tomlister/tilegame/engine/pipeline"
	"github.com/tomlister/tilegame/engine/world"
)

//Actor Stores image, position and logic for game components
type Actor struct {
	Tag        string
	Image      *ebiten.Image
	AltImages  []*ebiten.Image
	X          int
	Y          int
	Z          int
	ActorLogic func(actor *Actor, world *world.World, sceneDidMove bool)
	Static     bool
	Shadow     bool
	Collidable bool
	VelocityX  float64
	VelocityY  float64
	Direction  float64
	State      map[string]interface{}
	Disabled   bool
	Renderhook bool
	Rendercode func(actor *Actor, pipelinewrapper pipeline.PipelineWrapper, screen *ebiten.Image)
	Unpausable bool
	Kill       bool
}

func (a *Actor) RunActorLogic(w *world.World, sceneDidMove bool) {
	if (*w).State["pause"].(bool) == false {
		(*a).ActorLogic(a, w, sceneDidMove)
	} else {
		if (*a).Unpausable == true {
			(*a).ActorLogic(a, w, sceneDidMove)
		}
	}
}
