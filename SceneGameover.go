package main

import (
	_ "embed"
	"fmt"
	"image/color"
	"time"

	"github.com/dcortassa/super-flying-man-and-pig/assets"
	"github.com/dcortassa/super-flying-man-and-pig/globals"
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) UpdateGameOverState() {
	// TBD
}

func (g *Game) DrawGameOverState(screen *ebiten.Image) {
	var msg string
	msg = "GAME OVER"
	DrawTextByCenter(screen, msg, assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/2, assets.ColorRed)

	// if time.Now().Second()%2 == 0 {
	// 	msg = fmt.Sprintf("PRESS FIRE")
	// 	DrawTextByCenter(screen, msg, arcadeFont, screenWidth/2, screenHeight/2+30, color.White)
	// }
	if g.playerOne.Scores == g.HiScores || g.playerTwo.Scores == g.HiScores {
		if time.Now().Second()%2 == 0 {
			msg = fmt.Sprintf("NEW HISCORES %d !", g.HiScores)
			DrawTextByCenter(screen, msg, assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/2+30, color.White)
		}
	}
}
