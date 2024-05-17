package main

import (
	"fmt"
	"time"

	"github.com/dcortassa/super-flying-man-and-pig/assets"
	"github.com/dcortassa/super-flying-man-and-pig/debug"
	"github.com/dcortassa/super-flying-man-and-pig/globals"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) UpdateAttractState() {

	for _, enemy := range g.enemies {
		debug.DebugPrintf("enemy", enemy.Name, enemy.Active)
		enemy.Update()
	}

	for _, enemyBullet := range g.enemiesBullettPool {
		debug.DebugPrintf("enemyBullet", enemyBullet.Name, enemyBullet.Active)
		enemyBullet.Update()
	}
	debug.DebugPrintf("--- END Update Loop ---")

	CheckCollisions(g)

	if time.Since(g.lastEvent) > time.Millisecond*globals.ScrollSpeed {
		SpawnEnemies(g)
		g.lastEvent = time.Now()
		g.Position += 1 // pixel lines per scroll tick
		// tiles in a screen TilesScreenWidth*tilesScreenHeight
		// as g.Position is the low line pixel index
		// we reset to 0 when we have only one screen left of tiles
		if (g.Position/assets.TileSize)*globals.TilesScreenWidth >= len(assets.GameMap)-(globals.TilesScreenWidth*globals.TilesScreenHeight) {
			g.Position = 0
		}
	}

	// Remove inactive enemy entities every 2 seconds
	if time.Now().Second()%2 == 0 {
		CleanEnemyList(g)
	}
}

func (g *Game) DrawAttractState(screen *ebiten.Image) {
	activeBuls := 0
	activeEntities := 0
	totalEntities := 0
	tileIndex := 0
	// Draw world window
	rowPosition := (g.Position / assets.TileSize) * globals.TilesScreenWidth
	screenPosition := (g.Position % assets.TileSize)
	// for rowIndex := tilesScreenHeight - 1; rowIndex >= 0; rowIndex-- { // use this to show scrolling trick
	for rowIndex := globals.TilesScreenHeight; rowIndex >= 0; rowIndex-- {
		// fmt.Println("WPos: ", g.Position, "SPos:", screenPosition, "RPos:", rowPosition, "Row:", rowIndex, "-")
		for columnIndex := 0; columnIndex < globals.TilesScreenWidth; columnIndex++ {
			// fmt.Print("col:", columnIndex, " ")

			op := &ebiten.DrawImageOptions{}
			// op.GeoM.Translate(float64(tileSize*columnIndex), float64(tileSize*rowIndex+screenPosition)) // use this to show scrolling trick
			op.GeoM.Translate(float64(assets.TileSize*columnIndex), float64(assets.TileSize*rowIndex+screenPosition-assets.TileSize))
			tileIndex = assets.GameMap[rowPosition]
			rowPosition++
			screen.DrawImage(assets.GetTile(tileIndex).(*ebiten.Image), op)
		}
	}

	g.playerOne.Draw(screen)
	g.playerTwo.Draw(screen)

	for _, playerOneBullet := range g.playerOneBullettPool {
		playerOneBullet.Draw(screen)
		if playerOneBullet.Active {
			activeBuls += 1
		}
	}

	for _, playerTwoBullet := range g.playerTwoBullettPool {
		playerTwoBullet.Draw(screen)
		if playerTwoBullet.Active {
			activeBuls += 1
		}
	}

	for _, entity := range g.enemies {
		entity.Draw(screen)
		if entity.Active {
			activeEntities += 1
		}
		totalEntities += 1
	}

	for _, enemyBullet := range g.enemiesBullettPool {
		enemyBullet.Draw(screen)
		if enemyBullet.Active {
			activeBuls += 1
		}
	}

	// Draw Score/Lives
	var msg string

	if time.Now().Second()%2 == 0 {
		msg = fmt.Sprintf("1UP\nPRESS FIRE")
	} else {
		msg = fmt.Sprintf("1UP")
	}
	text.Draw(screen, msg, assets.ArcadeFont, 6, 21, assets.ColorBlack)
	text.Draw(screen, msg, assets.ArcadeFont, 5, 20, assets.ColorYellow)

	msg = fmt.Sprintf("HI-SCORE\n  %05d", g.HiScores)
	text.Draw(screen, msg, assets.ArcadeFont, 91, 21, assets.ColorBlack)
	text.Draw(screen, msg, assets.ArcadeFont, 90, 20, assets.ColorYellow)

	if time.Now().Second()%2 == 0 {
		msg = fmt.Sprintf("2UP\nPRESS FIRE")
	} else {
		msg = fmt.Sprintf("2UP")
	}
	text.Draw(screen, msg, assets.ArcadeFont, 171, 21, assets.ColorBlack)
	text.Draw(screen, msg, assets.ArcadeFont, 170, 20, assets.ColorYellow)

	// Draw debug data
	if globals.Debug {
		ebitenutil.DebugPrintAt(screen,
			fmt.Sprintf("TPS: %0.2f FPS: %0.2f \nPos:%v TEnt: %v AEnt: %v Blt: %v",
				ebiten.ActualTPS(), ebiten.ActualFPS(), g.Position, totalEntities, activeEntities, activeBuls),
			2, globals.ScreenHeight-30)

	}

	// Draw title
	op := &ebiten.DrawImageOptions{}
	DrawImageByCenter(screen, assets.TitleTextImage, globals.ScreenWidth/2, globals.ScreenHeight/3+20, op)

	// Draw attract text
	if time.Now().Second()%2 == 0 {
		DrawTextWithShadowByCenter(screen, "GAME OVER", assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/2+40, assets.ColorRed)
	} else {
		DrawTextWithShadowByCenter(screen, "GAME OVER", assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/2+40, assets.ColorYellow)
	}
	DrawTextWithShadowByCenter(screen, "PRESS FIRE TO PLAY", assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/2+60, assets.ColorWhite)

}
