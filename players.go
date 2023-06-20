package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
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

	g.playerOne.lives = 3
	mover := NewKeyboardMover(
		g.playerOne,
		Keybinds{
			Up:        ebiten.KeyArrowUp,
			Down:      ebiten.KeyArrowDown,
			Left:      ebiten.KeyArrowLeft,
			Right:     ebiten.KeyArrowRight,
			Fire:      ebiten.KeyControlLeft,
			Secondary: ebiten.KeyAltLeft,
		},
		Vector{1, 1},
	)
	g.playerOne.addComponent(mover)

	g.playerOneBullettPool = initBulletPool("P1Bullet", typePlayerOneBullet, animSuperFlyingManPew, ANIM_FPS, 5, Vector{0, -4}, Box{8, 2, 8, 8})
	shooter := NewKeyboardShooter(
		g.playerOne,
		ebiten.KeyControlLeft,
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

	g.playerTwo.lives = 3
	mover := NewKeyboardMover(
		g.playerTwo,
		Keybinds{
			Up:        ebiten.KeyW,
			Down:      ebiten.KeyS,
			Left:      ebiten.KeyA,
			Right:     ebiten.KeyD,
			Fire:      ebiten.KeyQ,
			Secondary: ebiten.KeyAltLeft,
		},
		Vector{1, 1},
	)
	g.playerTwo.addComponent(mover)

	g.playerTwoBullettPool = initBulletPool("P2Bullet", typePlayerTwoBullet, animPigPew, ANIM_FPS, 5, Vector{0, -4}, Box{8, 2, 8, 8})
	shooter := NewKeyboardShooter(
		g.playerTwo,
		ebiten.KeyQ,
		g.playerTwoBullettPool,
		time.Millisecond*250,
	)
	g.playerTwo.addComponent(shooter)

	gpMover := NewGamePadMover(g.playerTwo, 1, Vector{1, 1}, g.playerTwoBullettPool, time.Millisecond*250)
	g.playerTwo.addComponent(gpMover)
}
