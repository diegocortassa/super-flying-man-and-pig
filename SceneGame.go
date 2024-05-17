package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/dcortassa/super-flying-man-and-pig/assets"
	"github.com/dcortassa/super-flying-man-and-pig/debug"
	"github.com/dcortassa/super-flying-man-and-pig/globals"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) UpdateGameState() {

	debug.DebugPrintf("--- Update Loop ---")
	debug.DebugPrintf("position", g.Position)

	debug.DebugPrintf("player", g.playerOne.Name, g.playerOne.Active)
	g.playerOne.Update()

	debug.DebugPrintf("player", g.playerTwo.Name, g.playerTwo.Active)
	g.playerTwo.Update()

	for _, playerOneBullet := range g.playerOneBullettPool {
		debug.DebugPrintf("playerOneBullet", playerOneBullet.Name, playerOneBullet.Active)
		playerOneBullet.Update()
	}

	for _, playerTwoBullet := range g.playerTwoBullettPool {
		debug.DebugPrintf("playerTwoBullet", playerTwoBullet.Name, playerTwoBullet.Active)
		playerTwoBullet.Update()
	}

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
		// tiles in a screen tilesScreenWidth*tilesScreenHeight
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

func (g *Game) DrawGameState(screen *ebiten.Image) {
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

	if g.playerOne.Active {
		msg = fmt.Sprintf("1UP\n%05d", g.playerOne.Scores)
		DrawTextWithShadow(screen, strings.Repeat("*", g.playerOne.Lives), assets.ArcadeFont, 10, globals.ScreenHeight-10, assets.ColorYellow)
	} else {
		if time.Now().Second()%2 == 0 {
			msg = fmt.Sprintf("1UP\nPRESS FIRE")
		} else {
			msg = fmt.Sprintf("1UP\n%05d", g.playerOne.Scores)
		}
	}
	DrawTextWithShadow(screen, msg, assets.ArcadeFont, 5, 20, assets.ColorYellow)

	msg = fmt.Sprintf("HI-SCORE\n  %05d", g.HiScores)
	DrawTextWithShadow(screen, msg, assets.ArcadeFont, 90, 20, assets.ColorYellow)

	if g.playerTwo.Active {
		msg = fmt.Sprintf("2UP\n%05d", g.playerTwo.Scores)
		DrawTextWithShadow(screen, strings.Repeat("*", g.playerTwo.Lives), assets.ArcadeFont, globals.ScreenWidth-30, globals.ScreenHeight-10, assets.ColorYellow)
	} else {
		if time.Now().Second()%2 == 0 {
			msg = fmt.Sprintf("2UP\nPRESS FIRE")
		} else {
			msg = fmt.Sprintf("2UP\n%05d", g.playerTwo.Scores)
		}
	}
	DrawTextWithShadow(screen, msg, assets.ArcadeFont, 170, 20, assets.ColorYellow)

	// Draw debug data
	if globals.Debug {
		ebitenutil.DebugPrintAt(screen,
			fmt.Sprintf("TPS: %0.2f FPS: %0.2f \nPos:%v TEnt: %v AEnt: %v Blt: %v",
				ebiten.ActualTPS(), ebiten.ActualFPS(), g.Position, totalEntities, activeEntities, activeBuls),
			2, globals.ScreenHeight-30)

	}
}
