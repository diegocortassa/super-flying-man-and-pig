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

	for _, player := range g.players {
		DebugPrintf("player", player.name, player.active)
		player.Update(g)
	}

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

	if time.Since(lastUpdate) > time.Millisecond*scrollSpeed {
		spawnEnemies(g)
		lastUpdate = time.Now()
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

	for _, player := range g.players {
		player.Draw(screen)
	}

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
	textColor := color.RGBA{0xf6, 0xf4, 0x0d, 0xff}
	if g.players[0].active {
		msg = fmt.Sprintf("1UP %d\n%05d", g.players[0].lives, g.players[0].scores)
	} else {
		if time.Now().Second()%2 == 0 {
			msg = fmt.Sprintf("1UP\nPRESS FIRE")
		} else {
			msg = fmt.Sprintf("1UP\n%05d", g.players[0].scores)
		}
	}
	text.Draw(screen, msg, arcadeFont, 6, 21, color.Black)
	text.Draw(screen, msg, arcadeFont, 5, 20, textColor)

	msg = fmt.Sprintf("HI-SCORE\n  12200")
	text.Draw(screen, msg, arcadeFont, 91, 21, color.Black)
	text.Draw(screen, msg, arcadeFont, 90, 20, textColor)

	if g.players[1].active {
		msg = fmt.Sprintf("2UP %d\n%05d", g.players[1].lives, g.players[1].scores)
	} else {
		if time.Now().Second()%2 == 0 {
			msg = fmt.Sprintf("2UP\nPRESS FIRE")
		} else {
			msg = fmt.Sprintf("2UP\n%05d", g.players[1].scores)
		}
	}
	text.Draw(screen, msg, arcadeFont, 171, 21, color.Black)
	text.Draw(screen, msg, arcadeFont, 170, 20, textColor)

	// Draw debug data
	if debug {
		newLines := "\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n"
		ebitenutil.DebugPrint(screen,
			fmt.Sprintf("%vTPS: %0.2f FPS: %0.2f \n Pos:%v TEnt: %v AEnt: %v Blt: %v", newLines,
				ebiten.ActualTPS(), ebiten.ActualFPS(), g.position, totalEntities, activeEntities, activeBuls))
	}
}
