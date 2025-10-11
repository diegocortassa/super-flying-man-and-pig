package main

import (
	"os"
	"time"

	"github.com/diegocortassa/super-flying-man-and-pig/assets"
	"github.com/diegocortassa/super-flying-man-and-pig/input"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	StateCoolDown       = time.Millisecond * 1000
	attractRotationTime = 5 // Seconds before rotating attract screens
)

func (g *Game) UpdateDirector() {

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
			if assets.AudioPlayerPlaying != nil {
				assets.AudioPlayerPlaying.Play()
			}
			g.paused = false
		} else {
			if assets.AudioPlayerPlaying != nil {
				assets.AudioPlayerPlaying.Pause()
			}
			g.paused = true
		}
	}

	if g.paused {
		return
	}

	// At first init
	if g.CurrentState == StateInit {
		g.CurrentState = StateTitle
		// g.playerOne.Scores = 235
		// g.CurrentState = StateGameOver
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
		changed := false
		if g.playerOne.Scores >= g.HiScoresTable[4].Scores &&
			(input.IsP1FireJustPressed() ||
				input.IsP2FireJustPressed() ||
				time.Since(g.lastStateTransition) > time.Second*10) {

			// Prepare HiScores entry
			i := 0
			for i = 0; i < len(g.HiScoresTable); i++ {
				if g.playerOne.Scores >= g.HiScoresTable[i].Scores {
					g.HiScoresTable[i].Name = "AAA"
					g.HiScoresTable[i].Scores = g.playerOne.Scores
					break
				}
			}

			_ = g.ChangeState(StateHiscoresInsert)
			assets.PlayTheme(assets.ThemeBossFightPlayer)

		} else if input.IsP1FireJustPressed() || input.IsP2FireJustPressed() || time.Since(g.lastStateTransition) > time.Second*10 {
			changed = g.ChangeState(StateTitle)
		}
		if changed {
			// assets.PlayTheme(assets.Theme1StagePlayer)
			assets.StopAudioPlayer()
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

	// *SCENE* Hiscores
	if g.CurrentState == StateHiscores {
		g.CheckStartPressed()
		if time.Since(g.lastStateTransition) > time.Second*attractRotationTime {
			_ = g.ChangeState(StateCredits)
		}
		return
	}

	// *SCENE* HiscoresInsert
	if g.CurrentState == StateHiscoresInsert {
		if input.IsP1FireJustPressed() || input.IsP2FireJustPressed() || time.Since(g.lastStateTransition) > time.Second*28 {
			_ = g.ChangeState(StateHiscores)
			assets.StopAudioPlayer()
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
