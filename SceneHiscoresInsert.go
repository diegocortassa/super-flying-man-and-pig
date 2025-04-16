package main

import (
	_ "embed"
	"fmt"
	"image/color"
	"time"

	"github.com/dcortassa/super-flying-man-and-pig/assets"
	"github.com/dcortassa/super-flying-man-and-pig/globals"
	"github.com/dcortassa/super-flying-man-and-pig/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) UpdateHiscoresInsertState() {
	// Empty
}

func (g *Game) DrawHiscoreInsertState(screen *ebiten.Image) {

	var msg string
	var blinkColor color.Color
	var scorePosition int

	colors := []color.Color{
		color.RGBA{255, 0, 0, 255}, // Red
		color.RGBA{0, 255, 0, 255}, // Green
		color.RGBA{0, 0, 255, 255}, // Blue
	}

	// Calculate the current color index based on elapsed time.
	// Each color changes every 100 milliseconds, so cycle through 3 colors.
	elapsed := time.Since(g.lastEvent)
	colorIndex := int(elapsed.Milliseconds()/100) % len(colors)
	currentColor := colors[colorIndex]

	// Toggle text visibility every 500 milliseconds.
	// if elapsed.Milliseconds()%500 < 250 {
	if elapsed.Milliseconds()%1000 < 500 {
		blinkColor = assets.ColorYellow
	} else {
		blinkColor = assets.ColorRed
	}

	// if time.Now().Second()%2 == 0 {
	// 	blinkColor = assets.ColorYellow
	// } else {
	// 	blinkColor = assets.ColorWhite
	// }

	i := 0
	for i = 0; i < len(g.HiScoresTable); i++ {
		if g.playerOne.Scores >= g.HiScoresTable[i].Scores {
			scorePosition = i
			// g.HiScoresTable[i].Name = "1UP"
			// g.HiScoresTable[i].Scores = g.playerOne.Scores
			break
		}
	}

	DrawTextByCenter(screen, "BEST 5", assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/4, blinkColor)

	DrawTextByCenter(screen, "RANK", assets.ArcadeFont, globals.ScreenWidth/6, globals.ScreenHeight/4+20, assets.ColorWhite)
	DrawTextByCenter(screen, "SCORE", assets.ArcadeFont, globals.ScreenWidth/6*3, globals.ScreenHeight/4+20, assets.ColorWhite)
	DrawTextByCenter(screen, "NAME", assets.ArcadeFont, globals.ScreenWidth/6*5, globals.ScreenHeight/4+20, assets.ColorWhite)

	line := 40
	for i, HiScoresItem := range g.HiScoresTable {
		if i == scorePosition {
			DrawTextByCenter(screen, fmt.Sprintf("%1d", i+1), assets.ArcadeFont, globals.ScreenWidth/6, globals.ScreenHeight/4+line, assets.ColorRed)
			DrawTextByCenter(screen, fmt.Sprintf("%5d", g.HiScoresTable[i].Scores), assets.ArcadeFont, globals.ScreenWidth/6*3, globals.ScreenHeight/4+line, assets.ColorRed)
			// msg := fmt.Sprintf("%-*s", 3, g.HiScoresTable[i].Name)
			// msg := padRight(g.HiScoresTable[i].Name, 3, '_')
			DrawTextByCenter(screen, g.HiScoresTable[i].Name, assets.ArcadeFont, globals.ScreenWidth/6*5, globals.ScreenHeight/4+line, currentColor)
			x1 := float32((199) + 8*g.HiScoresInsertCursorPosition)
			y1 := float32((globals.ScreenHeight/4 + line) + 7)
			vector.StrokeLine(screen, x1, y1, x1+8, y1, 1, currentColor, false)

		} else {
			DrawTextByCenter(screen, fmt.Sprintf("%1d", i+1), assets.ArcadeFont, globals.ScreenWidth/6, globals.ScreenHeight/4+line, assets.ColorYellow)
			DrawTextByCenter(screen, fmt.Sprintf("%5d", HiScoresItem.Scores), assets.ArcadeFont, globals.ScreenWidth/6*3, globals.ScreenHeight/4+line, assets.ColorYellow)
			DrawTextByCenter(screen, HiScoresItem.Name, assets.ArcadeFont, globals.ScreenWidth/6*5, globals.ScreenHeight/4+line, assets.ColorWhite)
		}
		line += 20
	}

	if g.playerOne.Scores >= g.HiScoresTable[4].Scores {
		msg = "CONGRATULATIONS PLAYER ONE"
		DrawTextByCenter(screen, msg, assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/2+50, color.White)
		msg = "YOU MADE TO THE TOP 5 !"
		DrawTextByCenter(screen, msg, assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/2+70, color.White)
		msg = fmt.Sprintf("%1d", 28-int(time.Since(g.lastStateTransition).Seconds()))
		DrawTextByCenter(screen, msg, assets.ArcadeFont, globals.ScreenWidth/2, globals.ScreenHeight/2+90, assets.ColorRed)

		// Convert the string to a slice of runes
		runes := []rune(g.HiScoresTable[scorePosition].Name)

		if input.IsP1UpJustPressed() {
			firstLetter := runes[g.HiScoresInsertCursorPosition]
			runes[g.HiScoresInsertCursorPosition] = getNextRune(firstLetter)
			g.HiScoresTable[scorePosition].Name = string(runes)
		}
		if input.IsP1DownJustPressed() {
			firstLetter := runes[g.HiScoresInsertCursorPosition]
			runes[g.HiScoresInsertCursorPosition] = getPreviousRune(firstLetter)
			g.HiScoresTable[scorePosition].Name = string(runes)
		}
		if input.IsP1LeftJustPressed() {
			g.HiScoresInsertCursorPosition = (g.HiScoresInsertCursorPosition - 1 + 3) % 3
		}
		if input.IsP1RightJustPressed() {
			g.HiScoresInsertCursorPosition = (g.HiScoresInsertCursorPosition + 1) % 3
		}
	}

	// This is done every draw, must be moved out of DrawGameOverState
	// Re-sorts scores table
	// sort.Slice(g.HiScoresTable, func(i, j int) bool {
	// 	return g.HiScoresTable[i].Scores > g.HiScoresTable[j].Scores
	// })
}

func getNextRune(ch rune) rune {
	if ch == 'Z' {
		return 'A'
	} else {
		return ch + 1
	}
}

func getPreviousRune(ch rune) rune {
	if ch == 'A' {
		return 'Z'
	} else {
		return ch - 1
	}
}

// func padRight(s string, size int, padChar rune) string {
// 	if len(s) >= size {
// 		return s
// 	}
// 	padding := strings.Repeat(string(padChar), size-len(s))
// 	return s + padding
// }
