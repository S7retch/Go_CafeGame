package main

import (
	"fmt"
	"strings"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) updateBuild() {
	if g.Debug == true {
		fmt.Println("Build mode:", g.BuildMode)
	}
	if rl.IsKeyPressed(rl.KeyB) {
		g.BuildMode = !g.BuildMode
	}

	if g.BuildMode {
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			var item TileID = 3
			g.placeTile(g.Map.TileSet[item], g.MouseCell, g.Map.CurrentLayer)
		}
	}

}

func (g *Game) drawBuild() {
	// for index, value := range g.Map.LayerNames {
	// 	fmt.Printf("%v: %s\n", index, value)
	// 	// fmt.Println("ID:", g.Map.Layers[layer].Name)
	// }

	var width float32 = 60
	var height float32 = 60
	var num_layers int = len(g.Map.LayerNames)

	gui.ToggleGroup(
		rl.NewRectangle(float32(g.Settings.ScreenWidth/2)-(width/2)*float32(num_layers), float32(g.Settings.ScreenHeight)-160, width, height),
		strings.Join(g.Map.LayerNames, ";"), // Join the names of the layers automatically
		(*int32)(&g.Map.CurrentLayer))       // Needed to convert the enum to the build layer

	if g.Debug == true {
		fmt.Println("Current layer:", g.Map.CurrentLayer)
	}

}
