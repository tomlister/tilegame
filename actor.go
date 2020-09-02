package main

import (
	"github.com/hajimehoshi/ebiten"
)

//Actor Stores image, position and logic for game components
type Actor struct {
	Tag                     string
	Image                   *ebiten.Image
	AltImages               []*ebiten.Image
	RenderDestination       *ebiten.Image
	CustomRenderDestination bool
	X                       int
	Y                       int
	Z                       int
	ActorLogic              func(actor *Actor, world *World, sceneDidMove bool)
	Static                  bool
	Shadow                  bool
	Collidable              bool
	VelocityX               float64
	VelocityY               float64
	Direction               float64
	State                   map[string]interface{}
	Disabled                bool
	Renderhook              bool
	Rendercode              func(actor *Actor, pipelinewrapper PipelineWrapper, screen *ebiten.Image)
	Unpausable              bool
	Kill                    bool
}

func (actor *Actor) runActorLogic(world *World, sceneDidMove bool) {
	if (*world).State["pause"].(bool) == false {
		(*actor).ActorLogic(actor, world, sceneDidMove)
	} else {
		if (*actor).Unpausable == true {
			(*actor).ActorLogic(actor, world, sceneDidMove)
		}
	}
}
