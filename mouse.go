package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) updateMouse() {
	// I always want to get the mouse position
	g.MousePos = rl.GetScreenToWorld2D(rl.GetMousePosition(), g.Camera)
	g.MouseCell = get_cell(g.MousePos)

	if g.Debug == true {

		fmt.Println("Mouse pos:", g.MousePos)

	}

}

func (g *Game) drawMouse() {
	// rl.DrawTexture(g.Textures["Mouse"], int32(g.MousePos.X)-(CELL_SIZE/2), int32(g.MousePos.Y)-(CELL_SIZE/2), rl.White)

	rl.DrawTexture(g.Textures["Mouse"], int32(g.MouseCell.X), int32(g.MouseCell.Y), rl.White)

}
