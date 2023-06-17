package main

import (
	_ "embed"
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const freePlayBlinkTime = time.Millisecond * 500 // "free play" message blink time

var (
	freePlayOn bool // used to blink "free play" message
)

func (g *Game) UpdateTileState() {
	// TBD
}

func (g *Game) DrawTitleState(screen *ebiten.Image) {

	var msg string

	op := &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(25, 20)
	// screen.DrawImage(titleImage, op)
	DrawImageByCenter(screen, titleImage, screenWidth/2, screenHeight/3, op)

	op = &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(30, screenHeight/3)
	// screen.DrawImage(titleTextImage, op)
	DrawImageByCenter(screen, titleTextImage, screenWidth/2, screenHeight/3+60, op)

	if time.Since(g.lastEvent) > freePlayBlinkTime && freePlayOn {
		freePlayOn = !freePlayOn
		g.lastEvent = time.Now()
	}
	if time.Since(g.lastEvent) > freePlayBlinkTime && !freePlayOn {
		freePlayOn = !freePlayOn
		g.lastEvent = time.Now()
	}
	if freePlayOn {
		msg = fmt.Sprintf("FREE PLAY")
		DrawTextByCenter(screen, msg, arcadeFont, screenWidth/2, screenHeight/3*2-10, color.White)
	}

	msg = fmt.Sprintf("PRESS FIRE TO PLAY")
	DrawTextByCenter(screen, msg, arcadeFont, screenWidth/2, screenHeight/3*2+10, color.White)

}
