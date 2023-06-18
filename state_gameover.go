package main

import (
	_ "embed"
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) UpdateGameOverState() {
	// TBD
}

func (g *Game) DrawGameOverState(screen *ebiten.Image) {
	var msg string
	msg = fmt.Sprintf("GAME OVER")
	DrawTextByCenter(screen, msg, arcadeFont, screenWidth/2, screenHeight/2, ColorRed)

	// if time.Now().Second()%2 == 0 {
	// 	msg = fmt.Sprintf("PRESS FIRE")
	// 	DrawTextByCenter(screen, msg, arcadeFont, screenWidth/2, screenHeight/2+30, color.White)
	// }
	if g.playerOne.scores == g.hiScores || g.playerTwo.scores == g.hiScores {
		if time.Now().Second()%2 == 0 {
			msg = fmt.Sprintf("NEW HISCORES %d !", g.hiScores)
			DrawTextByCenter(screen, msg, arcadeFont, screenWidth/2, screenHeight/2+30, color.White)
		}
	}
}
