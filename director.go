package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	STATECOOLDOWN       = time.Millisecond * 1000
	attractRotationTime = 5 // Seconds before rotating attract screens
)

func (g *Game) UpdateDirector() {

	// g.state = StateHiscores

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	DebugPrintf(fmt.Sprintf("\tAlloc = %v MiB", m.Alloc/1024/1024))
	DebugPrintf(fmt.Sprintf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024))
	DebugPrintf(fmt.Sprintf("\tSys = %v MiB", m.Sys/1024/1024))
	DebugPrintf(fmt.Sprintf("\tNumGC = %v\n", m.NumGC))

	// In all States
	if IsExitJustPressed() {
		os.Exit(0)
	}

	if IsFullScreenJustPressed() {
		if ebiten.IsFullscreen() {
			ebiten.SetFullscreen(false)
		} else {
			ebiten.SetFullscreen(true)
		}
	}

	if IsResetJustPressed() {
		g.reset()
		changed := g.ChangeState(StateTitle)
		if changed {
			PlayTheme(Theme1StagePlayer)
		}
	}

	// At first init
	if g.state == StateInit {
		g.state = StateTitle
		PlayTheme(Theme1StagePlayer)
	}

	// *STATE* Game
	if g.state == StateGame {
		if !g.playerOne.active && !g.playerTwo.active {
			g.reset()
			changed := g.ChangeState(StateGameOver)
			if changed {
				PlayTheme(ThemeStageSelectPlayer)
			}
		}
		if !g.playerOne.active && IsP1FireJustPressed() {
			g.resetPlayerOne()
			g.playerOne.active = true
		}
		if !g.playerTwo.active && IsP2FireJustPressed() {
			g.resetPlayerTwo()
			g.playerTwo.active = true
		}
		return
	}

	// *STATE* Title
	if g.state == StateTitle {
		g.CheckStartPressed()
		if time.Since(g.lastStateTransition) > time.Second*attractRotationTime {
			g.reset()
			g.position = 400
			_ = g.ChangeState(StateAttract)
		}
		return
	}

	// *STATE* GameOver
	if g.state == StateGameOver {
		if IsP1FireJustPressed() || IsP2FireJustPressed() || time.Since(g.lastStateTransition) > time.Second*10 {
			changed := g.ChangeState(StateTitle)
			if changed {
				PlayTheme(Theme1StagePlayer)
			}
		}
		return
	}

	// *STATE* Attract
	if g.state == StateAttract {
		g.CheckStartPressed()
		if time.Since(g.lastStateTransition) > time.Second*attractRotationTime {
			_ = g.ChangeState(StateHiscores)
		}
		return
	}

	// *STATE* Highscores
	if g.state == StateHiscores {
		g.CheckStartPressed()
		if time.Since(g.lastStateTransition) > time.Second*attractRotationTime {
			_ = g.ChangeState(StateTitle)
		}
		return
	}
}

// Change game stare
func (g *Game) ChangeState(newState State) bool {
	if time.Since(g.lastStateTransition) > STATECOOLDOWN {
		g.PreviousState = g.state
		g.state = newState
		g.lastStateTransition = time.Now()
		return true
	}
	return false
}

func (g *Game) CheckStartPressed() {
	if IsP1FireJustPressed() {
		g.reset()
		g.resetPlayerOne()
		g.resetPlayerTwo()
		g.playerOne.active = true
		changed := g.ChangeState(StateGame)
		if changed {
			PlayTheme(Theme2StagePlayer)
		}
	}
	if IsP2FireJustPressed() {
		g.reset()
		g.resetPlayerOne()
		g.resetPlayerTwo()
		g.playerTwo.active = true
		changed := g.ChangeState(StateGame)
		if changed {
			PlayTheme(Theme2StagePlayer)
		}
	}
}
