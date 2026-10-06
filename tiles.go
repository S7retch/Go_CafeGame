package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TileID int

type BuildLayer int32

const (
	FloorLayer BuildLayer = iota
	WallsLayer
	FurnitureLayer
	FoodLayer
)

// This is each individual tile
type Tile struct {
	ID      string
	Texture rl.Texture2D
	Solid   bool
	Data    map[string]interface{}
}

func (g *Game) LoadTile(key string, filepath string, solid bool) error {

	// Using the LoadT from the textures page is not great but ehh
	id, err := g.LoadT(key, filepath)
	if err != nil {
		panic(err)
	}

	g.Map.TileSet[id] = Tile{ID: key, Texture: g.GetTexture(key), Solid: solid}

	return nil
}

// This is each layer
type TileLayer struct {
	Name    string
	Grid    [][]TileID // The int is the TIle ID
	Visible bool
	// Opacity float32
}

type TileMap struct {
	Width        int32
	Height       int32
	TileSize     int32
	Layers       []TileLayer
	LayerNames   []string
	TileSet      map[TileID]Tile // Loads all possible data
	SelectedTile TileID
	CurrentLayer BuildLayer
}

func (g *Game) createMap() {
	g.Map.Width = g.Cols
	g.Map.Height = g.Rows
	g.Map.TileSize = CELL_SIZE
	// fmt.Println(len(BuildLayer))
	g.Map.LayerNames = []string{"Floor", "Walls", "Furniture", "Food"}
	g.Map.Layers = make([]TileLayer, len(g.Map.LayerNames))
	g.Map.TileSet = make(map[TileID]Tile)

	// Loop through each layer and initialise it
	for index, value := range g.Map.LayerNames {
		g.Map.Layers[index].Name = value
		// Initialise the number of rows needed
		g.Map.Layers[index].Grid = make([][]TileID, g.Map.Height)

		// Initialise the number of cols within each row
		for i := range g.Map.Height {
			g.Map.Layers[index].Grid[i] = make([]TileID, g.Map.Width)
		}
		// fmt.Println(value, g.Map.Layers[index].Grid)

	}

	// fmt.Println()
}

func (g *Game) placeTile(tile Tile, position rl.Vector2, layer BuildLayer) error {
	// fmt.Println(g.Map.Layers[layer].Grid[int(position.Y)][int(position.X)])
	g.Map.Layers[layer].Grid[int(position.Y)][int(position.X)] = TileID(tile.Texture.ID)

	return nil
}

func (g *Game) drawTiles() {
	for layer := range g.Map.Layers {
		for rows := range g.Rows {
			for cols := range g.Cols {
				cell := g.Map.Layers[layer].Grid[rows][cols]
				rl.DrawTexture(g.Map.TileSet[cell].Texture, cols*CELL_SIZE, rows*CELL_SIZE, rl.White)
			}
		}
	}
}
