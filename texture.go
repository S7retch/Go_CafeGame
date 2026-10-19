package main

import (
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextureDictionary struct {
	Textures map[string]rl.Texture2D
}

func newTextureDictionary() *TextureDictionary {
	dict := make(map[string]rl.Texture2D)
	return &TextureDictionary{Textures: dict}
}

func (t *TextureDictionary) LoadT(key string, filepath string) error {

	if _, exists := t.Textures[key]; exists {
		return fmt.Errorf("texture with key '%s' already exists", key)
	}

	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		panic(err)
	}

	t.Textures[key] = rl.LoadTexture(filepath)

	return nil

}

func (t *TextureDictionary) UnloadTextures() {
	for _, value := range t.Textures {
		rl.UnloadTexture(value)
	}
	t.Textures = make(map[string]rl.Texture2D)
}

func (t *TextureDictionary) GetTexture(key string) rl.Texture2D {
	if _, exists := t.Textures[key]; exists {
		return t.Textures[key]
	} else {
		fmt.Printf("Texture %s does not exist, loading default texture", key)
		return rl.LoadTexture("assets/Default/InvalidTexture.png")
	}
}
