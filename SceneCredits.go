package main

import (
	_ "embed"
	"image/color"

	"github.com/dcortassa/super-flying-man-and-pig/assets"
	"github.com/dcortassa/super-flying-man-and-pig/globals"
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) UpdateCreditsState() {
	// Empty
}

func (g *Game) DrawCreditsState(screen *ebiten.Image) {
	line := 0
	DrawTextByCenter(screen, "CODE", assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/6, assets.ColorWhite)
	line += 15
	DrawTextByCenter(screen, "DIEGO CORTASSA", assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/6+line, assets.ColorYellow)
	line += 30
	DrawTextByCenter(screen, "GRAPHIC", assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/6+line, assets.ColorWhite)
	line += 15
	DrawTextByCenter(screen, "DIEGO CORTASSA", assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/6+line, assets.ColorYellow)
	line += 15
	DrawTextByCenter(screen, "LIVIO CORTASSA", assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/6+line, assets.ColorYellow)
	line += 30
	DrawTextByCenter(screen, "MUSIC & SOUND FX", assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/6+line, assets.ColorWhite)
	line += 15
	DrawTextByCenter(screen, "JUHANI JUNKALA", assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/6+line, assets.ColorYellow)
	line += 30
	DrawTextByCenter(screen, "TITLE IMAGE", assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/6+line, color.White)
	line += 15
	DrawTextByCenter(screen, "ANDREA PENNAZIO", assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/6+line, assets.ColorYellow)
	//
	DrawTextByCenter(screen, "PRESS FIRE TO PLAY", assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/3*2+20, assets.ColorWhite)
}
