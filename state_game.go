package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) UpdateGameState() {

	DebugPrintf("--- Update Loop ---")
	DebugPrintf("position", g.position)

	DebugPrintf("player", g.playerOne.name, g.playerOne.active)
	g.playerOne.Update(g)

	DebugPrintf("player", g.playerTwo.name, g.playerTwo.active)
	g.playerTwo.Update(g)

	for _, playerOneBullet := range g.playerOneBullettPool {
		DebugPrintf("playerOneBullet", playerOneBullet.name, playerOneBullet.active)
		playerOneBullet.Update(g)
	}

	for _, playerTwoBullet := range g.playerTwoBullettPool {
		DebugPrintf("playerTwoBullet", playerTwoBullet.name, playerTwoBullet.active)
		playerTwoBullet.Update(g)
	}

	for _, enemy := range g.enemies {
		DebugPrintf("enemy", enemy.name, enemy.active)
		enemy.Update(g)
	}

	for _, enemyBullet := range g.enemiesBullettPool {
		DebugPrintf("enemyBullet", enemyBullet.name, enemyBullet.active)
		enemyBullet.Update(g)
	}
	DebugPrintf("--- END Update Loop ---")

	CheckCollisions(g)

	if time.Since(g.lastEvent) > time.Millisecond*scrollSpeed {
		spawnEnemies(g)
		g.lastEvent = time.Now()
		g.position += 1 // pixel lines per scroll tick
		// tiles in a screen tilesScreenWidth*tilesScreenHeight
		// as g.position is the low line pixel index
		// we reset to 0 when we have only one screen left of tiles
		if (g.position/tileSize)*tilesScreenWidth >= len(gameMap)-(tilesScreenWidth*tilesScreenHeight) {
			g.position = 0
		}
	}

	// Remove inactive enemy entities every 2 seconds
	if time.Now().Second()%2 == 0 {
		cleanEnemyList(g)
	}
}

func (g *Game) DrawGameState(screen *ebiten.Image) {
	activeBuls := 0
	activeEntities := 0
	totalEntities := 0
	tileIndex := 0
	// Draw world window
	rowPosition := (g.position / tileSize) * tilesScreenWidth
	screenPosition := (g.position % tileSize)
	// for rowIndex := tilesScreenHeight - 1; rowIndex >= 0; rowIndex-- { // use this to show scrolling trick
	for rowIndex := tilesScreenHeight; rowIndex >= 0; rowIndex-- {
		// fmt.Println("WPos: ", g.position, "SPos:", screenPosition, "RPos:", rowPosition, "Row:", rowIndex, "-")
		for columnIndex := 0; columnIndex < tilesScreenWidth; columnIndex++ {
			// fmt.Print("col:", columnIndex, " ")

			op := &ebiten.DrawImageOptions{}
			// op.GeoM.Translate(float64(tileSize*columnIndex), float64(tileSize*rowIndex+screenPosition)) // use this to show scrolling trick
			op.GeoM.Translate(float64(tileSize*columnIndex), float64(tileSize*rowIndex+screenPosition-tileSize))
			tileIndex = gameMap[rowPosition]
			rowPosition++
			screen.DrawImage(getTile(tileIndex).(*ebiten.Image), op)
		}
	}

	g.playerOne.Draw(screen)
	g.playerTwo.Draw(screen)

	for _, playerOneBullet := range g.playerOneBullettPool {
		playerOneBullet.Draw(screen)
		if playerOneBullet.active {
			activeBuls += 1
		}
	}

	for _, playerTwoBullet := range g.playerTwoBullettPool {
		playerTwoBullet.Draw(screen)
		if playerTwoBullet.active {
			activeBuls += 1
		}
	}

	for _, entity := range g.enemies {
		entity.Draw(screen)
		if entity.active {
			activeEntities += 1
		}
		totalEntities += 1
	}

	for _, enemyBullet := range g.enemiesBullettPool {
		enemyBullet.Draw(screen)
		if enemyBullet.active {
			activeBuls += 1
		}
	}

	// Draw Score/Lives
	var msg string

	if g.playerOne.active {
		msg = fmt.Sprintf("1UP %d\n%05d", g.playerOne.lives, g.playerOne.scores)
	} else {
		if time.Now().Second()%2 == 0 {
			msg = fmt.Sprintf("1UP\nPRESS FIRE")
		} else {
			msg = fmt.Sprintf("1UP\n%05d", g.playerOne.scores)
		}
	}
	text.Draw(screen, msg, arcadeFont, 6, 21, color.Black)
	text.Draw(screen, msg, arcadeFont, 5, 20, ColorYellow)

	msg = fmt.Sprintf("HI-SCORE\n  %05d", g.hiScores)
	text.Draw(screen, msg, arcadeFont, 91, 21, color.Black)
	text.Draw(screen, msg, arcadeFont, 90, 20, ColorYellow)

	if g.playerTwo.active {
		msg = fmt.Sprintf("2UP %d\n%05d", g.playerTwo.lives, g.playerTwo.scores)
	} else {
		if time.Now().Second()%2 == 0 {
			msg = fmt.Sprintf("2UP\nPRESS FIRE")
		} else {
			msg = fmt.Sprintf("2UP\n%05d", g.playerTwo.scores)
		}
	}
	text.Draw(screen, msg, arcadeFont, 171, 21, color.Black)
	text.Draw(screen, msg, arcadeFont, 170, 20, ColorYellow)

	// Draw debug data
	if debug {
		ebitenutil.DebugPrintAt(screen,
			fmt.Sprintf("TPS: %0.2f FPS: %0.2f \nPos:%v TEnt: %v AEnt: %v Blt: %v",
				ebiten.ActualTPS(), ebiten.ActualFPS(), g.position, totalEntities, activeEntities, activeBuls),
			2, screenHeight-30)

	}
}
