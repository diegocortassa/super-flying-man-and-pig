package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

func (g *Game) resetPlayerOne() {
	// player one
	g.playerOne = newEntity(
		"P1",
		Vector{x: (screenWidth - spriteSize) / 4, y: screenHeight - spriteSize - 20},
	)
	g.playerOne.active = false
	g.playerOne.hitBoxes = append(g.playerOne.hitBoxes, Box{6, 3, 12, 17})

	sequences := map[string]*sequence{
		"idle":    newSequence(animSuperFlyingMan, ANIM_FPS, true),
		"destroy": newSequence(animSuperFlyingManDie, ANIM_FPS, false),
	}
	animator := newAnimator(g.playerOne, sequences, "idle")
	g.playerOne.addComponent(animator)

	sounds := map[Sound]*audio.Player{SoundDestroy: sfx_exp_odd3Player, SoundFire: sfx_wpn_laser1Player}
	soundPlayer := newSoundPlayer(g.playerOne, sounds)
	g.playerOne.addComponent(soundPlayer)

	g.playerOne.lives = 3

	keyBinds := Keybinds{
		Up:    ebiten.KeyArrowUp,
		Down:  ebiten.KeyArrowDown,
		Left:  ebiten.KeyArrowLeft,
		Right: ebiten.KeyArrowRight,
		Fire:  ebiten.KeyAltRight,
	}
	if MameKeys {
		keyBinds.Fire = ebiten.KeyControlLeft
	}
	mover := NewKeyboardMover(g.playerOne, keyBinds, Vector{1, 1})
	g.playerOne.addComponent(mover)

	g.playerOneBullettPool = initBulletPool("P1Bullet", typePlayerOneBullet, animSuperFlyingManPew, ANIM_FPS, 5, Vector{0, -4}, Box{8, 2, 8, 8})
	shooter := NewKeyboardShooter(
		g.playerOne,
		keyBinds.Fire,
		g.playerOneBullettPool,
		time.Millisecond*250,
	)
	g.playerOne.addComponent(shooter)

	gpMover := NewGamePadMover(g.playerOne, 0, Vector{1, 1}, g.playerOneBullettPool, time.Millisecond*250)
	g.playerOne.addComponent(gpMover)
}

func (g *Game) resetPlayerTwo() {
	// player two
	g.playerTwo = newEntity(
		"P2",
		Vector{x: (screenWidth - spriteSize) / 4 * 3, y: screenHeight - spriteSize - 20},
	)
	g.playerTwo.active = false
	g.playerTwo.hitBoxes = append(g.playerTwo.hitBoxes, Box{5, 3, 12, 17})

	sequences := map[string]*sequence{
		"idle":    newSequence(animPig, ANIM_FPS, true),
		"destroy": newSequence(animPigDie, ANIM_FPS, false),
	}
	animator := newAnimator(g.playerTwo, sequences, "idle")
	g.playerTwo.addComponent(animator)

	sounds := map[Sound]*audio.Player{SoundDestroy: sfx_exp_odd3Player, SoundFire: sfx_wpn_laser1Player}
	soundPlayer := newSoundPlayer(g.playerTwo, sounds)
	g.playerTwo.addComponent(soundPlayer)

	g.playerTwo.lives = 3

	var keyBinds = Keybinds{}
	if MameKeys {
		keyBinds = Keybinds{
			Up:    ebiten.KeyR,
			Down:  ebiten.KeyF,
			Left:  ebiten.KeyD,
			Right: ebiten.KeyG,
			Fire:  ebiten.KeyA,
		}
	} else {
		keyBinds = Keybinds{
			Up:    ebiten.KeyW,
			Down:  ebiten.KeyS,
			Left:  ebiten.KeyA,
			Right: ebiten.KeyD,
			Fire:  ebiten.KeyQ,
		}
	}
	mover := NewKeyboardMover(g.playerTwo, keyBinds, Vector{1, 1})
	g.playerTwo.addComponent(mover)

	g.playerTwoBullettPool = initBulletPool("P2Bullet", typePlayerTwoBullet, animPigPew, ANIM_FPS, 5, Vector{0, -4}, Box{8, 2, 8, 8})
	shooter := NewKeyboardShooter(
		g.playerTwo,
		keyBinds.Fire,
		g.playerTwoBullettPool,
		time.Millisecond*250,
	)
	g.playerTwo.addComponent(shooter)

	gpMover := NewGamePadMover(g.playerTwo, 1, Vector{1, 1}, g.playerTwoBullettPool, time.Millisecond*250)
	g.playerTwo.addComponent(gpMover)
}
