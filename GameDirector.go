package main

import (
	"os"
	"time"

	"github.com/dcortassa/super-flying-man-and-pig/assets"
	"github.com/dcortassa/super-flying-man-and-pig/input"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	StateCoolDown       = time.Millisecond * 1000
	attractRotationTime = 5 // Seconds before rotating attract screens
)

func (g *Game) UpdateDirector() {

	// g.state = StateHiscores

	// moved to hud debug
	// var m runtime.MemStats
	// runtime.ReadMemStats(&m)
	// debug.DebugPrintf(fmt.Sprintf("\tAlloc = %v MiB", m.Alloc/1024/1024))
	// debug.DebugPrintf(fmt.Sprintf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024))
	// debug.DebugPrintf(fmt.Sprintf("\tSys = %v MiB", m.Sys/1024/1024))
	// debug.DebugPrintf(fmt.Sprintf("\tNumGC = %v\n", m.NumGC))

	// In all States
	if input.IsExitJustPressed() {
		os.Exit(0)
	}

	if input.IsFullScreenJustPressed() {
		if ebiten.IsFullscreen() {
			ebiten.SetFullscreen(false)
		} else {
			ebiten.SetFullscreen(true)
		}
	}

	if input.IsResetJustPressed() {
		g.reset()
		changed := g.ChangeState(StateTitle)
		if changed {
			// assets.PlayTheme(assets.Theme1StagePlayer)
			assets.StopAudioPlayer()
		}
	}

	if input.IsPauseJustPressed() {
		if g.paused {
			assets.AudioPlayerPlaying.Play()
			g.paused = false
		} else {
			assets.AudioPlayerPlaying.Pause()
			g.paused = true
		}
	}

	if g.paused {
		return
	}

	// At first init
	if g.CurrentState == StateInit {
		g.CurrentState = StateTitle
		// assets.PlayTheme(assets.Theme1StagePlayer)
	}

	// *SCENE* Game
	if g.CurrentState == StateGame {
		if !g.playerOne.Active &&
			!g.playerTwo.Active &&
			time.Since(g.playerOne.InvulnerableSetTime) > StateCoolDown &&
			time.Since(g.playerTwo.InvulnerableSetTime) > StateCoolDown {
			g.reset()
			changed := g.ChangeState(StateGameOver)
			if changed {
				assets.PlayTheme(assets.ThemeStageSelectPlayer)
			}
		}
		if !g.playerOne.Active && input.IsP1FireJustPressed() {
			g.resetPlayerOne()
			g.playerOne.Active = true
			g.playerOne.Invulnerable = true
			g.playerOne.InvulnerableSetTime = time.Now()
		}
		if !g.playerTwo.Active && input.IsP2FireJustPressed() {
			g.resetPlayerTwo()
			g.playerTwo.Active = true
			g.playerTwo.Invulnerable = true
			g.playerTwo.InvulnerableSetTime = time.Now()

		}
		return
	}

	// *SCENE* Title
	if g.CurrentState == StateTitle {
		g.CheckStartPressed()
		if time.Since(g.lastStateTransition) > time.Second*attractRotationTime {
			g.reset()
			// start attract near the beach
			g.Position = 400
			_ = g.ChangeState(StateAttract)
		}
		return
	}

	// *SCENE* GameOver
	if g.CurrentState == StateGameOver {
		if input.IsP1FireJustPressed() || input.IsP2FireJustPressed() || time.Since(g.lastStateTransition) > time.Second*10 {
			changed := g.ChangeState(StateTitle)
			if changed {
				// assets.PlayTheme(assets.Theme1StagePlayer)
				assets.StopAudioPlayer()
			}
		}
		return
	}

	// *SCENE* Attract
	if g.CurrentState == StateAttract {
		g.CheckStartPressed()
		if time.Since(g.lastStateTransition) > time.Second*attractRotationTime*3 {
			_ = g.ChangeState(StateHiscores)
		}
		return
	}

	// *SCENE* Highscores
	if g.CurrentState == StateHiscores {
		g.CheckStartPressed()
		if time.Since(g.lastStateTransition) > time.Second*attractRotationTime {
			_ = g.ChangeState(StateCredits)
		}
		return
	}

	// *SCENE* Credits
	if g.CurrentState == StateCredits {
		g.CheckStartPressed()
		if time.Since(g.lastStateTransition) > time.Second*attractRotationTime {
			_ = g.ChangeState(StateTitle)
		}
		return
	}
}

// Change game state
func (g *Game) ChangeState(newState State) bool {
	if time.Since(g.lastStateTransition) > StateCoolDown {
		g.PreviousState = g.CurrentState
		g.CurrentState = newState
		g.lastStateTransition = time.Now()
		return true
	}
	return false
}

func (g *Game) CheckStartPressed() {
	if input.IsP1FireJustPressed() {
		g.reset()
		g.resetPlayerOne()
		g.resetPlayerTwo()
		g.playerOne.Active = true
		changed := g.ChangeState(StateGame)
		if changed {
			assets.PlayTheme(assets.Theme2StagePlayer)
		}
	}
	if input.IsP2FireJustPressed() {
		g.reset()
		g.resetPlayerOne()
		g.resetPlayerTwo()
		g.playerTwo.Active = true
		changed := g.ChangeState(StateGame)
		if changed {
			assets.PlayTheme(assets.Theme2StagePlayer)
		}
	}
}
