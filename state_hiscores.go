package main

import (
	_ "embed"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) UpdateHiscoresState() {
	// TBD
}

func (g *Game) DrawHiscoreState(screen *ebiten.Image) {

	DrawTextByCenter(screen, "BEST 5", arcadeFont, screenWidth/2, screenHeight/4, ColorRed)

	DrawTextByCenter(screen, "RANK", arcadeFont, screenWidth/6, screenHeight/4+20, color.White)
	DrawTextByCenter(screen, "SCORE", arcadeFont, screenWidth/6*3, screenHeight/4+20, color.White)
	DrawTextByCenter(screen, "NAME", arcadeFont, screenWidth/6*5, screenHeight/4+20, color.White)

	var hiScores = [][]string{{"1", "50000", "DIE"}, {"2", "40000", "LIV"}, {"3", "30000", "AND"}, {"4", "20000", "NOR"}, {"5", "10000", "MRJ"}}
	line := 40

	for _, score := range hiScores {
		DrawTextByCenter(screen, score[0], arcadeFont, screenWidth/6, screenHeight/4+line, ColorYellow)
		DrawTextByCenter(screen, score[1], arcadeFont, screenWidth/6*3, screenHeight/4+line, ColorYellow)
		DrawTextByCenter(screen, score[2], arcadeFont, screenWidth/6*5, screenHeight/4+line, color.White)
		line += 20
	}

	DrawTextByCenter(screen, "PRESS FIRE TO PLAY", arcadeFont, screenWidth/2, screenHeight/3*2+20, color.White)
}
