package main

import (
	"encoding/gob"
	"log"
	"os"
)

type GameState struct {
	Attributes []Attribute
	Hotbar
	Inventory [9]Item
	XP        int
	Seed      int
	PlayerX   int
	PlayerY   int
}

func (w *World) saveGame() {
	player := w.Actors[w.TagTable["Player"]]
	gs := GameState{}
	gs.Attributes = player.State["attributes"].([]Attribute)
	gs.Hotbar = player.State["hotbar"].(Hotbar)
	gs.Inventory = player.State["inventory"].([9]Item)
	gs.XP = player.State["xp"].(int)
	gs.Seed = w.Seed
	gs.PlayerX = player.X
	gs.PlayerY = player.Y
	file, err := os.Create("save.state")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	enc := gob.NewEncoder(file)
	err = enc.Encode(gs)
	if err != nil {
		log.Fatal(err)
	}
}

func (w *World) loadGame() GameState {
	file, err := os.Open("save.state")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	enc := gob.NewDecoder(file)
	gs := GameState{}
	err = enc.Decode(&gs)
	if err != nil {
		log.Fatal(err)
	}
	return gs
}

func checkForState() bool {
	if _, err := os.Stat("save.state"); os.IsNotExist(err) {
		return false
	}
	return true
}
