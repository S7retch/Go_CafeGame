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
	source := rl.Rectangle{X: 0, Y: 16, Width: 16, Height: 16}
	destination := rl.Rectangle{X: g.MouseCell.X * float32(CELL_SIZE), Y: g.MouseCell.Y * float32(CELL_SIZE), Width: 16, Height: 16}
	origin := rl.Vector2{X: 0, Y: 0}
	rotation := float32(0.0)
	rl.DrawTexturePro(g.Textures.GetTexture("UI"), source, destination, origin, rotation, rl.White)

}
