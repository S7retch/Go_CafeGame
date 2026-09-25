package main

import (
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextureDictionary struct {
	Textures map[string]rl.Texture2D
}

func (g *Game) LoadT(key string, filepath string) error {

	if _, exists := g.Textures[key]; exists {
		return fmt.Errorf("texture with key '%s' already exists", key)
		// fmt.Println("Invalid key of: %s, loading default texture", key)
	}

	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		fmt.Printf("File %s does not exist\n", filepath)
		g.Textures[key] = rl.LoadTexture("assets/Default/InvalidTexture.png")
	} else if err == nil {
		g.Textures[key] = rl.LoadTexture(filepath)
	} else {
		fmt.Println("Error:", err)
	}

	return err

}

func (g *Game) UnloadTextures() {
	for _, value := range g.Textures {
		rl.UnloadTexture(value)
	}
	g.Textures = make(map[string]rl.Texture2D)
}

func (g *Game) GetTexture(key string) rl.Texture2D {
	if _, exists := g.Textures[key]; exists {
		return g.Textures[key]
	} else {
		fmt.Printf("Texture %s does not exist, loading default texture", key)
		return rl.LoadTexture("assets/Default/InvalidTexture.png")
	}
}
