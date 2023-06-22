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
	if g.CurrentState == StateInit {
		g.CurrentState = StateTitle
		PlayTheme(Theme1StagePlayer)
	}

	// *STATE* Game
	if g.CurrentState == StateGame {
		if !g.playerOne.active &&
			!g.playerTwo.active &&
			time.Since(g.playerOne.invulnerableSetTime) > STATECOOLDOWN &&
			time.Since(g.playerTwo.invulnerableSetTime) > STATECOOLDOWN {
			g.reset()
			changed := g.ChangeState(StateGameOver)
			if changed {
				PlayTheme(ThemeStageSelectPlayer)
			}
		}
		if !g.playerOne.active && IsP1FireJustPressed() {
			g.resetPlayerOne()
			g.playerOne.active = true
			g.playerOne.invulnerable = true
			g.playerOne.invulnerableSetTime = time.Now()
		}
		if !g.playerTwo.active && IsP2FireJustPressed() {
			g.resetPlayerTwo()
			g.playerTwo.active = true
			g.playerTwo.invulnerable = true
			g.playerTwo.invulnerableSetTime = time.Now()

		}
		return
	}

	// *STATE* Title
	if g.CurrentState == StateTitle {
		g.CheckStartPressed()
		if time.Since(g.lastStateTransition) > time.Second*attractRotationTime {
			g.reset()
			g.position = 400
			_ = g.ChangeState(StateAttract)
		}
		return
	}

	// *STATE* GameOver
	if g.CurrentState == StateGameOver {
		if IsP1FireJustPressed() || IsP2FireJustPressed() || time.Since(g.lastStateTransition) > time.Second*10 {
			changed := g.ChangeState(StateTitle)
			if changed {
				PlayTheme(Theme1StagePlayer)
			}
		}
		return
	}

	// *STATE* Attract
	if g.CurrentState == StateAttract {
		g.CheckStartPressed()
		if time.Since(g.lastStateTransition) > time.Second*attractRotationTime {
			_ = g.ChangeState(StateHiscores)
		}
		return
	}

	// *STATE* Highscores
	if g.CurrentState == StateHiscores {
		g.CheckStartPressed()
		if time.Since(g.lastStateTransition) > time.Second*attractRotationTime {
			_ = g.ChangeState(StateCredits)
		}
		return
	}

	// *STATE* Credits
	if g.CurrentState == StateCredits {
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
		g.PreviousState = g.CurrentState
		g.CurrentState = newState
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
