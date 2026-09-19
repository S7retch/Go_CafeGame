package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func get_cell(vector rl.Vector2) rl.Vector2 {

	return rl.Vector2{X: float32(math.Floor(float64(vector.X / float32(CELL_SIZE)))), Y: float32(math.Floor(float64(vector.Y / float32(CELL_SIZE))))}

}
