package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *GameState) CameraUpdate() {

	if rl.IsMouseButtonDown(rl.MouseButtonMiddle) {
		delta := rl.GetMouseDelta()

		delta = rl.Vector2Scale(delta, float32(-1.0)/g.Camera.Zoom)
		g.Camera.Target = rl.Vector2Add(g.Camera.Target, delta)
		fmt.Println(g.Camera.Target)
	}

	wheel := rl.GetMouseWheelMove()

	if wheel != float32(0.0) {

		fmt.Println("mouse wheel")
		mouse_world_pos := rl.GetScreenToWorld2D(g.MousePos, g.Camera)
		g.Camera.Offset = g.MousePos
		g.Camera.Target = mouse_world_pos
		scale := float32(0.2) * wheel
		g.Camera.Zoom = rl.Clamp(float32(math.Exp(math.Log(float64(g.Camera.Zoom))+float64(scale))), 1.0, 64.0)
	}

	fmt.Println(g.Camera.Zoom)

	if g.Camera.Zoom == float32(1.0) { //# Put camera back to centre, otherwise can go out of bounds

		g.Camera.Offset = rl.Vector2{X: float32(g.Settings.ScreenWidth / 2), Y: float32(g.Settings.ScreenHeight / 2)}
		g.Camera.Target = rl.Vector2{X: float32(g.Settings.ScreenWidth / 2), Y: float32(g.Settings.ScreenHeight / 2)}
	}

	screen_w := rl.GetScreenWidth()
	screen_h := rl.GetScreenHeight()
	view_w := float32(screen_w) / g.Camera.Zoom
	view_h := float32(screen_h) / g.Camera.Zoom

	// top-left of view in world coords = target - offset / zoom
	// so allowed target.x range is:
	min_tx := g.Camera.Offset.X / g.Camera.Zoom
	max_tx := float32(g.Settings.ScreenWidth) - view_w + (g.Camera.Offset.X / g.Camera.Zoom)
	fmt.Println(max_tx)

	min_ty := float32(g.Camera.Offset.Y) / float32(g.Camera.Zoom)
	max_ty := float32(g.Settings.ScreenHeight) - view_h + (g.Camera.Offset.Y / g.Camera.Zoom)
	fmt.Println("min: ", min_ty, "max: ", max_ty)

	// handle case where world is smaller than view: center camera on world
	if max_tx <= min_tx {
		g.Camera.Target.X = float32(g.Settings.ScreenWidth) * 0.5
	} else {
		g.Camera.Target.X = rl.Clamp(g.Camera.Target.X, min_tx, max_tx)
	}

	if max_ty <= min_ty {
		g.Camera.Target.Y = float32(g.Settings.ScreenHeight) * 0.5
	} else {
		g.Camera.Target.Y = rl.Clamp(g.Camera.Target.Y, min_ty, max_ty)
	}
}
