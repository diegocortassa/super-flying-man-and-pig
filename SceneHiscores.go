package main

import (
	_ "embed"
	"fmt"

	"github.com/diegocortassa/super-flying-man-and-pig/assets"
	"github.com/diegocortassa/super-flying-man-and-pig/globals"
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) UpdateHiscoresState() {
	// Empty
}

func (g *Game) DrawHiscoreState(screen *ebiten.Image) {

	DrawTextByCenter(screen, "BEST 5", assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/4, assets.ColorRed)

	DrawTextByCenter(screen, "RANK", assets.ArcadeFont, globals.ScreenWidth/6, globals.ScreenHeight/4+20, assets.ColorWhite)
	DrawTextByCenter(screen, "SCORE", assets.ArcadeFont, globals.ScreenWidth/6*3, globals.ScreenHeight/4+20, assets.ColorWhite)
	DrawTextByCenter(screen, "NAME", assets.ArcadeFont, globals.ScreenWidth/6*5, globals.ScreenHeight/4+20, assets.ColorWhite)

	line := 40
	for i, HiScoresItem := range g.HiScoresTable {
		DrawTextByCenter(screen, fmt.Sprintf("%1d", i+1), assets.ArcadeFont, globals.ScreenWidth/6, globals.ScreenHeight/4+line, assets.ColorYellow)
		DrawTextByCenter(screen, fmt.Sprintf("%5d", HiScoresItem.Scores), assets.ArcadeFont, globals.ScreenWidth/6*3, globals.ScreenHeight/4+line, assets.ColorYellow)
		DrawTextByCenter(screen, HiScoresItem.Name, assets.ArcadeFont, globals.ScreenWidth/6*5, globals.ScreenHeight/4+line, assets.ColorWhite)
		line += 20
	}

	DrawTextByCenter(screen, "PRESS FIRE TO PLAY", assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/3*2+20, assets.ColorWhite)
}
