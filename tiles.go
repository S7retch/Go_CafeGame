package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TileID string

type LayerID string

type BuildLayer int32

// Layers := make(map[BuildLayer]LayerID)

// const (
// WalkLayer BuildLayer = iota
// FloorLayer
// WallsLayer
// FurnitureLayer
// FoodLayer
// )

// This is each individual tile
type Tile struct {
	ID           TileID
	Texture      rl.Texture2D
	TexRectangle rl.Rectangle
	Solid        bool
	Layer        LayerID
}

// This is each layer
type TileLayer struct {
	Name    LayerID
	Grid    [][]TileID
	Visible bool
}

// This is all layers combined
type TileMap struct {
	Tileset    map[TileID]Tile
	TileLayer  map[LayerID]TileLayer
	LayerNames map[BuildLayer]LayerID
}

func (m *TileMap) newTile(ID TileID, Texture rl.Texture2D, TexRectangle rl.Rectangle, Solid bool, Layer LayerID) *Tile {
	tile := &Tile{ID: ID, Texture: Texture, TexRectangle: TexRectangle, Solid: Solid, Layer: Layer}

	// &TileSet[ID] = tile
	m.Tileset[ID] = *tile

	return tile
}

func (m *TileMap) loadTiles(t TextureDictionary) {
	m.newTile("Stone Floor", t.GetTexture("Floor"), rl.Rectangle{X: 0, Y: 0, Width: 16, Height: 16}, false, "Floor")
	m.newTile("Burger", t.GetTexture("Food"), rl.Rectangle{X: 0, Y: 0, Width: 16, Height: 16}, false, "Food")
	m.newTile("Drink", t.GetTexture("Food"), rl.Rectangle{X: 16, Y: 0, Width: 16, Height: 16}, false, "Food")
	m.newTile("Chair", t.GetTexture("Furniture"), rl.Rectangle{X: 0, Y: 0, Width: 16, Height: 16}, false, "Furniture")
	m.newTile("Table", t.GetTexture("Furniture"), rl.Rectangle{X: 0, Y: 16, Width: 16, Height: 16}, false, "Furniture")
	// t.GetTexture()
	// fmt.Println(t)
}

func newTileLayer(ID LayerID, Rows int, Cols int) *TileLayer {
	grid := make([][]TileID, Rows)

	layer := &TileLayer{Name: ID, Grid: grid}

	return layer
}

func (m *TileMap) loadLayers() {

}

func (m *TileMap) getLayerNames() []string {
	names := make([]string, 0)
	for key, _ := range m.TileLayer {
		names = append(names, string(key))
	}
	return names
}

var LayerIDMap = map[BuildLayer]LayerID{
	0: "Walk",
	1: "Floor",
	2: "Walls",
	3: "Furniture",
	4: "Food",
}

func newMap() *TileMap {

	// layid := map[BuildLayer]LayerID{
	// 	1: "Walk",
	// 	2: "Floor",
	// 	3: "Walls",
	// 	4: "Furniture",
	// 	5: "Food",
	// }

	tset := make(map[TileID]Tile)
	tlayer := make(map[LayerID]TileLayer)
	// layid := []LayerID{"Walkable", "Floor", "Walls", "Furniture", "Food"}

	return &TileMap{tset, tlayer, LayerIDMap}

}

// This is each layer
// type TileLayer struct {
// 	Name    string
// 	Grid    [][]TileID // The int is the TIle ID
// 	Visible bool
// 	// Opacity float32
// }

// type TileMap struct {
// 	Width        int32
// 	Height       int32
// 	TileSize     int32
// 	Layers       []TileLayer
// 	LayerNames   []string
// 	TileSet      map[TileID]Tile // Loads all possible data
// 	SelectedTile TileID
// 	CurrentLayer BuildLayer
// }

// func (g *Game) createMap() {

// g.Map.Width = g.Cols
// g.Map.Height = g.Rows
// g.Map.TileSize = CELL_SIZE
// // fmt.Println(len(BuildLayer))
// g.Map.LayerNames = []string{"Floor", "Walls", "Furniture", "Food"}
// g.Map.Layers = make([]TileLayer, len(g.Map.LayerNames))
// g.Map.TileSet = make(map[TileID]Tile)

// // Loop through each layer and initialise it
// for index, value := range g.Map.LayerNames {
// 	g.Map.Layers[index].Name = value
// 	// Initialise the number of rows needed
// 	g.Map.Layers[index].Grid = make([][]TileID, g.Map.Height)

// 	// Initialise the number of cols within each row
// 	for i := range g.Map.Height {
// 		g.Map.Layers[index].Grid[i] = make([]TileID, g.Map.Width)
// 	}
// 	// fmt.Println(value, g.Map.Layers[index].Grid)

// }

// fmt.Println()
// }

// func (g *Game) placeTile(tile Tile, position rl.Vector2, layer BuildLayer) error {
// 	// fmt.Println(g.Map.Layers[layer].Grid[int(position.Y)][int(position.X)])
// 	g.Map.Layers[layer].Grid[int(position.Y)][int(position.X)] = TileID(tile.Texture.ID)

// 	return nil
// }

// func (g *Game) drawTiles() {
// 	for layer := range g.Map.Layers {
// 		for rows := range g.Rows {
// 			for cols := range g.Cols {
// 				cell := g.Map.Layers[layer].Grid[rows][cols]
// 				rl.DrawTexture(g.Map.TileSet[cell].Texture, cols*CELL_SIZE, rows*CELL_SIZE, rl.White)
// 			}
// 		}
// 	}
// }
