package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ScreenState int

const (
	StatePlay ScreenState = iota
	StatePaused
)

type GameSettings struct {
	ScreenWidth  int32
	ScreenHeight int32
}

type Game struct {
	Settings GameSettings
	TextureDictionary
	State ScreenState
	// Canvas   rl.RenderTexture2D
	Cols int32
	Rows int32
	// Scene    GameScene
	Debug bool

	Camera    rl.Camera2D
	MousePos  rl.Vector2
	MouseCell rl.Vector2
	// MouseCell rl.Vector2

	Map       TileMap
	BuildMode bool
}

const (
	CELL_SIZE int32 = 16
)

// This runs before the main() loop
func (g *Game) Init() {

	g.Settings.ScreenWidth = 800
	g.Settings.ScreenHeight = 640
	g.State = StatePlay
	g.Textures = make(map[string]rl.Texture2D)

	// g.Canvas = rl.LoadRenderTexture(g.Settings.ScreenWidth, g.Settings.ScreenHeight)

	g.Cols = g.Settings.ScreenWidth / CELL_SIZE
	g.Rows = g.Settings.ScreenHeight / CELL_SIZE
	// g.Paused = false
	g.Debug = false

	g.Camera.Zoom = float32(1.0)
	g.Camera.Offset = rl.Vector2{X: float32(g.Settings.ScreenWidth / 2), Y: float32(g.Settings.ScreenHeight / 2)}
	g.Camera.Target = rl.Vector2{X: float32(g.Settings.ScreenWidth / 2), Y: float32(g.Settings.ScreenHeight / 2)}

}

func (g *Game) Update() {

	if g.Debug == true {
		fmt.Println("Game state:", g.State)
	}

	// Check which state the game is in
	switch g.State {
	case StatePaused:
		g.updatePaused()
	case StatePlay:
		g.updatePlay()
	}

}

func (g *Game) updatePlay() {
	if rl.IsKeyPressed(rl.KeyP) {
		g.State = StatePaused
	}
	g.updateBuild() // Check if build mode is toggled
	// Camera update
	g.updateCamera()
	g.updateMouse()
}

func (g *Game) Draw() {

	rl.BeginDrawing()

	switch g.State {
	case StatePaused:
		g.drawPaused()
	case StatePlay:
		g.drawPlay()
	}

	rl.EndDrawing()
}

func (g *Game) drawPlay() {

	rl.BeginMode2D(g.Camera)

	rl.DrawRectangle(1, 1, 10, 10, rl.Green)

	rl.ClearBackground(rl.RayWhite)

	for i := int32(0); i <= g.Settings.ScreenWidth; i += CELL_SIZE {
		rl.DrawLine(i, 0, i, g.Settings.ScreenHeight, rl.LightGray)
	}
	for i := int32(0); i <= g.Settings.ScreenHeight; i += CELL_SIZE {
		rl.DrawLine(0, i, g.Settings.ScreenWidth, i, rl.LightGray)
	}

	if g.BuildMode {

		rl.DrawText("Building", 100, 10, 20, rl.LightGray)
		g.drawBuild()
	}

	g.drawTiles()

	g.drawMouse()

	rl.EndMode2D()

	// Draw after camera, so it's fixed
	rl.DrawText(fmt.Sprintf("%f", rl.GetFrameTime()), 10, 10, 20, rl.LightGray)

}

func main() {

	game := Game{}
	game.Init()

	rl.InitWindow(game.Settings.ScreenWidth, game.Settings.ScreenHeight, "Cafe game")
	rl.SetTargetFPS(60)

	defer rl.CloseWindow()

	game.createMap()

	// Load textures
	game.LoadT("Mouse", "assets/Mouse/MouseSquare.png") // Different because its not a tile
	game.LoadTile("BrickWall", "assets/Walls/BrickWall.png", true)
	game.LoadTile("Burger", "assets/Food/Burger.png", false)
	game.LoadTile("Drink", "assets/Food/Drink.png", false)
	game.LoadTile("Table", "assets/Furniture/Table.png", true)
	game.LoadTile("Chair", "assets/Furniture/Chair.png", true)
	game.LoadTile("FloorTile", "assets/Floor/Floor.png", false)

	defer game.UnloadTextures()

	for !rl.WindowShouldClose() {

		game.Update()

		game.Draw()

	}
}
