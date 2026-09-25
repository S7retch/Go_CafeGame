package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) updatePaused() {
	if rl.IsKeyPressed(rl.KeyP) {
		g.State = StatePlay
	}
}

func (g *Game) drawPaused() {
	// Dark transparent overlay.
	rl.DrawRectangle(
		0,
		0,
		g.Settings.ScreenWidth,
		g.Settings.ScreenHeight,
		rl.NewColor(0, 0, 0, 160),
	)

	title := "PAUSED"
	titleSize := int32(40)

	titleWidth := rl.MeasureText(title, titleSize)
	rl.DrawText(
		title,
		(g.Settings.ScreenWidth-titleWidth)/2,
		150,
		titleSize,
		rl.RayWhite,
	)
}
