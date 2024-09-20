package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"

	"github.com/dcortassa/super-flying-man-and-pig/assets"
	"github.com/dcortassa/super-flying-man-and-pig/globals"
	"github.com/dcortassa/super-flying-man-and-pig/input"
)

func (g *Game) resetPlayerOne() {
	// player one
	g.playerOne = NewEntity(
		"P1",
		Vector{X: (globals.ScreenWidth - assets.SpriteSize) / 4, Y: globals.ScreenHeight - assets.SpriteSize - 20},
	)
	g.playerOne.Active = false
	g.playerOne.HitBoxes = append(g.playerOne.HitBoxes, Box{6, 3, 12, 17})

	sequences := map[string]*Sequence{
		"idle":    NewSequence(assets.AnimSuperFlyingMan, assets.AnimFPS, true),
		"destroy": NewSequence(assets.AnimSuperFlyingManDie, assets.AnimFPS, false),
	}
	animator := NewAnimator(g.playerOne, sequences, "idle")
	g.playerOne.AddComponent(animator)

	sounds := map[Sound]*audio.Player{SoundDestroy: assets.Sfx_exp_odd3Player, SoundFire: assets.Sfx_wpn_laser1Player}
	soundPlayer := NewSoundPlayer(g.playerOne, sounds)
	g.playerOne.AddComponent(soundPlayer)

	g.playerOne.Lives = g.startLives

	keyBinds := Keybinds{
		Up:    input.IsP1UpPressed,
		Down:  input.IsP1DownPressed,
		Left:  input.IsP1LeftPressed,
		Right: input.IsP1RightPressed,
		Fire:  input.IsP1FirePressed,
	}
	mover := NewKeyboardMover(g.playerOne, keyBinds, Vector{X: 1, Y: 1})
	g.playerOne.AddComponent(mover)

	g.playerOneBullettPool = initBulletPool("P1Bullet", TypePlayerOneBullet, assets.AnimSuperFlyingManPew, assets.AnimFPS, 5, Vector{X: 0, Y: -4}, Box{X: 8, Y: 2, W: 8, H: 8})
	shooter := NewKeyboardShooter(
		g.playerOne,
		input.IsP1FirePressed,
		g.playerOneBullettPool,
		time.Millisecond*250,
	)
	g.playerOne.AddComponent(shooter)

	gpMover := NewGamePadMover(g.playerOne, 0, Vector{X: 1, Y: 1}, g.playerOneBullettPool, time.Millisecond*250)
	g.playerOne.AddComponent(gpMover)
}

func (g *Game) resetPlayerTwo() {
	// player two
	g.playerTwo = NewEntity(
		"P2",
		Vector{X: (globals.ScreenWidth - assets.SpriteSize) / 4 * 3, Y: globals.ScreenHeight - assets.SpriteSize - 20},
	)
	g.playerTwo.Active = false
	g.playerTwo.HitBoxes = append(g.playerTwo.HitBoxes, Box{X: 5, Y: 3, W: 12, H: 17})

	sequences := map[string]*Sequence{
		"idle":    NewSequence(assets.AnimPig, assets.AnimFPS, true),
		"destroy": NewSequence(assets.AnimPigDie, assets.AnimFPS, false),
	}
	animator := NewAnimator(g.playerTwo, sequences, "idle")
	g.playerTwo.AddComponent(animator)

	sounds := map[Sound]*audio.Player{SoundDestroy: assets.Sfx_exp_odd3Player, SoundFire: assets.Sfx_wpn_laser1Player}
	soundPlayer := NewSoundPlayer(g.playerTwo, sounds)
	g.playerTwo.AddComponent(soundPlayer)

	g.playerTwo.Lives = g.startLives

	keyBinds := Keybinds{
		Up:    input.IsP2UpPressed,
		Down:  input.IsP2DownPressed,
		Left:  input.IsP2LeftPressed,
		Right: input.IsP2RightPressed,
		Fire:  input.IsP2FirePressed,
	}
	mover := NewKeyboardMover(g.playerTwo, keyBinds, Vector{X: 1, Y: 1})
	g.playerTwo.AddComponent(mover)

	g.playerTwoBullettPool = initBulletPool("P2Bullet", TypePlayerTwoBullet, assets.AnimPigPew, assets.AnimFPS, 5, Vector{X: 0, Y: -4}, Box{X: 8, Y: 2, W: 8, H: 8})
	shooter := NewKeyboardShooter(
		g.playerTwo,
		input.IsP2FirePressed,
		g.playerTwoBullettPool,
		time.Millisecond*250,
	)
	g.playerTwo.AddComponent(shooter)

	gpMover := NewGamePadMover(g.playerTwo, 1, Vector{X: 1, Y: 1}, g.playerTwoBullettPool, time.Millisecond*250)
	g.playerTwo.AddComponent(gpMover)
}
