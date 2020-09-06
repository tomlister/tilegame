package main

func floaterActorLogic(actor *Actor, world *World, sceneDidMove bool) {
	if (*actor).State["Interval"].(int) == 3 {
		(*actor).State["Interval"] = 0
		if (*actor).State["AnimCount"].(int) == 20 {
			(*actor).State["AnimCount"] = 0
			(*actor).Kill = true
		} else {
			(*actor).State["AnimCount"] = (*actor).State["AnimCount"].(int) + 1
		}
		(*actor).Y--
	} else {
		(*actor).State["Interval"] = (*actor).State["Interval"].(int) + 1
	}
}
