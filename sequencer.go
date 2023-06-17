package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) UpdateSequencer() {

	// g.mode = ModeTitle
	// g.mode = ModeGameOver

	// soundThemeSource := audio.NewInfiniteLoop(soundTheme, soundTheme.Length())
	// audioThemePlayer, err := audio.NewPlayer(audioContext, soundThemeSource)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// g.audioPlayer.SetVolume(0.02)
	// g.audioPlayer.Play()

	// sePlayer := audio.NewPlayerFromBytes(audioContext, bs)
	// // sePlayer is never GCed as long as it plays.
	// sePlayer.Play()
	// Load and play game theme

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	DebugPrintf(fmt.Sprintf("\tAlloc = %v MiB", m.Alloc/1024/1024))
	DebugPrintf(fmt.Sprintf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024))
	DebugPrintf(fmt.Sprintf("\tSys = %v MiB", m.Sys/1024/1024))
	DebugPrintf(fmt.Sprintf("\tNumGC = %v\n", m.NumGC))

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		log.Println("Bye")
		os.Exit(0)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		if ebiten.IsFullscreen() {
			ebiten.SetFullscreen(false)
		} else {
			ebiten.SetFullscreen(true)
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.reset()
	}

	if g.mode == ModeTitle {
		if inpututil.IsKeyJustPressed(ebiten.KeyControlLeft) || inpututil.IsKeyJustPressed(ebiten.KeyQ) {
			g.mode = ModeGame
		}
	}

	if g.mode == ModeGame {
		if !g.players[0].active && !g.players[1].active {

			// soundThemeSource := audio.NewInfiniteLoop(audioStageSelectTheme, audioStageSelectTheme.Length())
			// audioPlayer, err := audio.NewPlayer(audioContext, soundThemeSource)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// audioPlayer.SetVolume(0.04)
			// audioPlayer.Play()

			g.reset()
			g.mode = ModeGameOver
		}
	}

}
