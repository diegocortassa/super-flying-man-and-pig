package main

import (
	_ "embed"
	"time"

	"github.com/diegocortassa/super-flying-man-and-pig/assets"
	"github.com/diegocortassa/super-flying-man-and-pig/globals"
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) UpdateTileState() {
	// TBD
}

func (g *Game) DrawTitleState(screen *ebiten.Image) {

	// Draw title image
	op := &ebiten.DrawImageOptions{}
	DrawImageByCenter(screen, assets.TitleImage, globals.ScreenWidth/2, globals.ScreenHeight/3, op)

	// Draw title text
	op = &ebiten.DrawImageOptions{}
	DrawImageByCenter(screen, assets.TitleTextImage, globals.ScreenWidth/2, globals.ScreenHeight/3+60, op)

	if time.Now().Second()%2 == 0 {
		DrawTextByCenter(screen, "FREE PLAY", assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/3*2, assets.ColorWhite)
	} else {
		DrawTextByCenter(screen, "FREE PLAY", assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/3*2, assets.ColorRed)
	}

	DrawTextByCenter(screen, "PRESS FIRE TO PLAY", assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/3*2+20, assets.ColorWhite)

	DrawTextByCenter(screen, "© 1985   DIEGO CORTASSA", assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/8*7, assets.ColorYellow)

}
