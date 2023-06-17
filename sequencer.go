package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const STATECOOLDOWN = time.Millisecond * 1000

func (g *Game) UpdateSequencer() {

	// g.mode = ModeTitle
	// g.mode = ModeGameOver

	var err error

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	DebugPrintf(fmt.Sprintf("\tAlloc = %v MiB", m.Alloc/1024/1024))
	DebugPrintf(fmt.Sprintf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024))
	DebugPrintf(fmt.Sprintf("\tSys = %v MiB", m.Sys/1024/1024))
	DebugPrintf(fmt.Sprintf("\tNumGC = %v\n", m.NumGC))

	// Always trasition
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
		g.ChangeState(StateTitle)
	}

	if g.state == StateInit {
		changed := g.ChangeState(StateTitle)
		if changed {
			soundThemeSource := audio.NewInfiniteLoop(audio1StageTheme, audio1StageTheme.Length()+1)
			audioPlayer, err = audio.NewPlayer(audioContext, soundThemeSource)
			if err != nil {
				log.Fatal(err)
			}
			audioPlayer.SetVolume(0.05)
			audioPlayer.Play()
		}
	}

	// State trasitions
	if g.state == StateTitle {
		if inpututil.IsKeyJustPressed(ebiten.KeyControlLeft) {
			g.reset()
			g.resetPlayerOne()
			g.resetPlayerTwo()
			g.playerOne.active = true
			changed := g.ChangeState(StateGame)
			if changed {
				audioPlayer.Close()
				soundThemeSource := audio.NewInfiniteLoop(audio2StageTheme, audio2StageTheme.Length()+1)
				audioPlayer, err = audio.NewPlayer(audioContext, soundThemeSource)
				if err != nil {
					log.Fatal(err)
				}
				audioPlayer.SetVolume(0.05)
				audioPlayer.Play()
			}
		}
	}

	if g.state == StateTitle {
		if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
			g.reset()
			g.resetPlayerOne()
			g.resetPlayerTwo()
			g.playerTwo.active = true
			changed := g.ChangeState(StateGame)
			if changed {
				fmt.Println("going to game")
				audioPlayer.Close()
				soundThemeSource := audio.NewInfiniteLoop(audio2StageTheme, audio2StageTheme.Length()+1)
				audioPlayer, err = audio.NewPlayer(audioContext, soundThemeSource)
				if err != nil {
					log.Fatal(err)
				}
				audioPlayer.SetVolume(0.05)
				audioPlayer.Play()
			}
		}
	}

	if g.state == StateGame {
		if !g.playerOne.active && !g.playerTwo.active {
			g.reset()
			changed := g.ChangeState(StateGameOver)
			if changed {
				audioPlayer.Close()
				soundThemeSource := audio.NewInfiniteLoop(audioStageSelectTheme, audioStageSelectTheme.Length()+1)
				audioPlayer, err = audio.NewPlayer(audioContext, soundThemeSource)
				if err != nil {
					log.Fatal(err)
				}
				audioPlayer.SetVolume(0.05)
				audioPlayer.Play()
			}
		}
		if !g.playerOne.active && inpututil.IsKeyJustPressed(ebiten.KeyControlLeft) {
			g.resetPlayerOne()
			g.playerOne.active = true
		}
		if !g.playerTwo.active && inpututil.IsKeyJustPressed(ebiten.KeyQ) {
			g.resetPlayerTwo()
			g.playerTwo.active = true
		}
	}

	if g.state == StateGameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeyControlLeft) || inpututil.IsKeyJustPressed(ebiten.KeyQ) {
			g.reset()
			changed := g.ChangeState(StateTitle)
			if changed {
				audioPlayer.Close()
				soundThemeSource := audio.NewInfiniteLoop(audio1StageTheme, audio1StageTheme.Length()+1)
				audioPlayer, err = audio.NewPlayer(audioContext, soundThemeSource)
				if err != nil {
					log.Fatal(err)
				}
				audioPlayer.SetVolume(0.05)
				audioPlayer.Play()
			}
		}
	}
}

func (g *Game) ChangeState(newState State) bool {
	if time.Since(g.lastStateTransition) > STATECOOLDOWN {
		g.PreviousState = g.state
		g.state = newState
		g.lastStateTransition = time.Now()
		return true
	}
	return false
}
