package main

import (
	_ "embed"
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) UpdateGameOverMode() {
}

func (g *Game) DrawGameOverMode(screen *ebiten.Image) {
	var msg string
	textColor := color.RGBA{0xf6, 0x00, 0x00, 0xff}
	msg = fmt.Sprintf("GAME OVER")
	DrawTextByCenter(screen, msg, arcadeFont, screenWidth/2, screenHeight/2, textColor)

	// if time.Now().UnixMilli()%500 == 0 {
	if time.Now().Second()%2 == 0 {
		msg = fmt.Sprintf("PRESS R TO RESTART")
		DrawTextByCenter(screen, msg, arcadeFont, screenWidth/2, screenHeight/2+30, color.White)
	}
}
