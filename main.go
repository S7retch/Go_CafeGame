package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameScene struct {
}

type GameSettings struct {
	ScreenWidth  int32
	ScreenHeight int32
}

type GameState struct {
	Settings GameSettings
	Paused   bool
	Cols     int32
	Rows     int32
	Scene    GameScene

	Camera   rl.Camera2D
	MousePos rl.Vector2
	// MouseCell rl.Vector2
}

const (
	CELL_SIZE int32 = 16
)

// This runs before the main() loop
func (g *GameState) Init() {

	g.Settings.ScreenWidth = 800
	g.Settings.ScreenHeight = 640
	g.Cols = g.Settings.ScreenWidth / CELL_SIZE
	g.Rows = g.Settings.ScreenHeight / CELL_SIZE
	g.Paused = false

	g.Camera.Zoom = float32(1.0)
	g.Camera.Offset = rl.Vector2{X: float32(g.Settings.ScreenWidth / 2), Y: float32(g.Settings.ScreenHeight / 2)}
	g.Camera.Target = rl.Vector2{X: float32(g.Settings.ScreenWidth / 2), Y: float32(g.Settings.ScreenHeight / 2)}
}

func (g *GameState) Update() {
	g.MousePos = rl.GetMousePosition()
	// fmt.Println(get_cell(g.MousePos))

	// Camera update
	g.CameraUpdate()
}

func (g *GameState) Draw() {

	rl.BeginDrawing()

	rl.BeginMode2D(g.Camera)

	rl.DrawRectangle(1, 1, 10, 10, rl.Green)

	rl.ClearBackground(rl.RayWhite)
	// rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
	rl.DrawText(fmt.Sprintf("%f", rl.GetFrameTime()), 10, 10, 20, rl.LightGray)

	for i := int32(0); i <= g.Settings.ScreenWidth; i += CELL_SIZE {
		rl.DrawLine(i, 0, i, g.Settings.ScreenHeight, rl.LightGray)
	}
	for i := int32(0); i <= g.Settings.ScreenHeight; i += CELL_SIZE {
		rl.DrawLine(0, i, g.Settings.ScreenWidth, i, rl.LightGray)
	}

	rl.EndMode2D()

	rl.EndDrawing()
}

func main() {

	game := GameState{}
	game.Init()

	rl.InitWindow(game.Settings.ScreenWidth, game.Settings.ScreenHeight, "Cafe game")
	rl.SetTargetFPS(60)

	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {

		game.Update()

		game.Draw()

	}
}
