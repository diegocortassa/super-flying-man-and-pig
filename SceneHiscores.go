package main

import (
	_ "embed"

	"github.com/dcortassa/superflyingmanandpig/assets"
	"github.com/dcortassa/superflyingmanandpig/globals"
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) UpdateHiscoresState() {
	// TBD
}

func (g *Game) DrawHiscoreState(screen *ebiten.Image) {

	DrawTextByCenter(screen, "BEST 5", assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/4, assets.ColorRed)

	DrawTextByCenter(screen, "RANK", assets.ArcadeFont, globals.ScreenWidth/6, globals.ScreenHeight/4+20, assets.ColorWhite)
	DrawTextByCenter(screen, "SCORE", assets.ArcadeFont, globals.ScreenWidth/6*3, globals.ScreenHeight/4+20, assets.ColorWhite)
	DrawTextByCenter(screen, "NAME", assets.ArcadeFont, globals.ScreenWidth/6*5, globals.ScreenHeight/4+20, assets.ColorWhite)

	var HiScores = [][]string{{"1", "50000", "DIE"}, {"2", "40000", "LIV"}, {"3", "30000", "AND"}, {"4", "20000", "NOR"}, {"5", "10000", "MRJ"}}
	line := 40

	for _, score := range HiScores {
		DrawTextByCenter(screen, score[0], assets.ArcadeFont, globals.ScreenWidth/6, globals.ScreenHeight/4+line, assets.ColorYellow)
		DrawTextByCenter(screen, score[1], assets.ArcadeFont, globals.ScreenWidth/6*3, globals.ScreenHeight/4+line, assets.ColorYellow)
		DrawTextByCenter(screen, score[2], assets.ArcadeFont, globals.ScreenWidth/6*5, globals.ScreenHeight/4+line, assets.ColorWhite)
		line += 20
	}

	DrawTextByCenter(screen, "PRESS FIRE TO PLAY", assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/3*2+20, assets.ColorWhite)
}
