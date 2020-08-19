package main

func backgroundActorLogic(actor *Actor, world *World, sceneDidMove bool) {

}

func backgroundTreeActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	if actor.DetectBehind(world) {
		(*actor).Z = 1
	} else {
		(*actor).Z = 0
	}
	if (*actor).State["Interval"].(int) == 25 {
		(*actor).State["Interval"] = 0
		if (*actor).State["AnimCount"].(int) == len((*actor).AltImages)-1 {
			(*actor).State["AnimCount"] = 0
		} else {
			(*actor).State["AnimCount"] = (*actor).State["AnimCount"].(int) + 1
		}
		(*actor).Image = (*actor).AltImages[(*actor).State["AnimCount"].(int)]
	} else {
		(*actor).State["Interval"] = (*actor).State["Interval"].(int) + 1
	}
	if (*actor).State["health"] == 0 {
		(*actor).Kill = true
	}
}

func backgroundRockActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	if (*actor).State["health"] == 0 {
		(*actor).Kill = true
	}
}
