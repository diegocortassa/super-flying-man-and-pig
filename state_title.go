package main

import (
	_ "embed"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const freePlayBlinkTime = time.Millisecond * 500 // "free play" message blink time

func (g *Game) UpdateTileState() {
	// TBD
}

func (g *Game) DrawTitleState(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(25, 20)
	// screen.DrawImage(titleImage, op)
	DrawImageByCenter(screen, titleImage, screenWidth/2, screenHeight/3, op)

	op = &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(30, screenHeight/3)
	// screen.DrawImage(titleTextImage, op)
	DrawImageByCenter(screen, titleTextImage, screenWidth/2, screenHeight/3+60, op)

	if time.Now().Second()%2 == 0 {
		DrawTextByCenter(screen, "FREE PLAY", arcadeFont, screenWidth/2, screenHeight/3*2, color.White)
	} else {
		DrawTextByCenter(screen, "FREE PLAY", arcadeFont, screenWidth/2, screenHeight/3*2, ColorRed)
	}

	DrawTextByCenter(screen, "PRESS FIRE TO PLAY", arcadeFont, screenWidth/2, screenHeight/3*2+20, color.White)

	DrawTextByCenter(screen, "© 1985   DIEGO CORTASSA", arcadeFont, screenWidth/2, screenHeight/8*7, ColorYellow)

}
