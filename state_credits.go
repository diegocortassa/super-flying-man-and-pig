package main

import (
	_ "embed"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) UpdateCreditsState() {
	// Empty
}

func (g *Game) DrawCreditsState(screen *ebiten.Image) {
	line := 0
	DrawTextByCenter(screen, "CODE", arcadeFont, screenWidth/2, screenHeight/6, color.White)
	line += 15
	DrawTextByCenter(screen, "DIEGO CORTASSA", arcadeFont, screenWidth/2, screenHeight/6+line, ColorYellow)
	line += 30
	DrawTextByCenter(screen, "GRAPHIC", arcadeFont, screenWidth/2, screenHeight/6+line, color.White)
	line += 15
	DrawTextByCenter(screen, "DIEGO CORTASSA", arcadeFont, screenWidth/2, screenHeight/6+line, ColorYellow)
	line += 15
	DrawTextByCenter(screen, "LIVIO CORTASSA", arcadeFont, screenWidth/2, screenHeight/6+line, ColorYellow)
	line += 30
	DrawTextByCenter(screen, "MUSIC & SOUND FX", arcadeFont, screenWidth/2, screenHeight/6+line, color.White)
	line += 15
	DrawTextByCenter(screen, "JUHANI JUNKALA", arcadeFont, screenWidth/2, screenHeight/6+line, ColorYellow)
	line += 30
	DrawTextByCenter(screen, "TITLE IMAGE", arcadeFont, screenWidth/2, screenHeight/6+line, color.White)
	line += 15
	DrawTextByCenter(screen, "ANDREA PENNAZIO", arcadeFont, screenWidth/2, screenHeight/6+line, ColorYellow)
	//
	DrawTextByCenter(screen, "PRESS FIRE TO PLAY", arcadeFont, screenWidth/2, screenHeight/3*2+20, color.White)
}
