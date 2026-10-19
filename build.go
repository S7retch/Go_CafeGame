package main

import (
	"fmt"
	"strings"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BuildMode struct {
	Active    bool
	LayerName LayerID
	LayerInt  int32
}

func (g *Game) updateBuild() {
	if g.Debug == true {
		fmt.Println("Build mode:", g.BuildMode)
	}
	if rl.IsKeyPressed(rl.KeyB) {
		g.BuildMode.Active = !g.BuildMode.Active
	}

	if g.BuildMode.Active {
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			// var item TileID = 3
			// g.placeTile(g.Map.TileSet[item], g.MouseCell, g.Map.CurrentLayer)
		}
	}

}

func (g *Game) drawBuild() {

	var width float32 = 60
	var height float32 = 60
	var num_layers int = len(g.Map.LayerNames)
	// Convert the LayerID to a string
	layerLabels := make([]string, num_layers)
	for i, layer := range g.Map.LayerNames {
		layerLabels[i] = fmt.Sprint(layer)
	}

	gui.ToggleGroup(
		rl.NewRectangle(float32(g.Settings.ScreenWidth/2)-(width/2)*float32(num_layers), float32(g.Settings.ScreenHeight)-160, width, height),
		strings.Join(layerLabels, ";"), // Join the layer labels for the toggle group.
		&g.BuildMode.LayerInt)

	if g.Debug == true {
		fmt.Println("Current layer:", g.BuildMode.LayerInt)
	}

}
