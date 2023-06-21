package main

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type AimedShooter struct {
	active    bool
	container *Entity
	p1        *Entity
	p2        *Entity
	trigger   time.Duration
	pool      []*Entity
	lastShot  time.Time
}

func NewAimedShooter(container *Entity, trigger time.Duration, pool []*Entity, p1, p2 *Entity) *AimedShooter {
	return &AimedShooter{
		active:    true,
		pool:      pool,
		container: container,
		p1:        p1,
		p2:        p2,
		trigger:   trigger,
		lastShot:  time.Now(),
	}
}

func (shooter *AimedShooter) Update() {
	if time.Since(shooter.lastShot) >= shooter.trigger {
		shooter.shoot(shooter.container.position.x+25, shooter.container.position.y-20)
		shooter.lastShot = time.Now()
	}
	return
}

func (shooter *AimedShooter) Draw(screen *ebiten.Image) {
	return
}

// Shoot bullet from pool starting at position x,y
func (shooter *AimedShooter) shoot(x, y float64) {
	if bul, ok := BullettFromPool(shooter.pool); ok {
		// do not shoot while exploding
		if shooter.container.exploding {
			return
		}

		mover := bul.getComponent(&ConstantMover{}).(*ConstantMover)

		var px, py float64
		distP1 := Distance(shooter.container, shooter.p1)
		distP2 := Distance(shooter.container, shooter.p2)

		// find nearest player
		if !shooter.p2.active || distP1 < distP2 {
			px = shooter.p1.position.x
			py = shooter.p1.position.y
		} else {
			px = shooter.p2.position.x
			py = shooter.p2.position.y
		}

		sx := shooter.container.position.x
		sy := shooter.container.position.y

		speed := 2.0

		mover.speed = Vector{0, 2}
		dx := px - sx
		dy := py - sy

		distance := math.Sqrt(dx*dx + dy*dy)
		// multiply the normalized result by desired speed
		speedX := (dx / distance) * speed
		speedY := (dy / distance) * speed

		mover.speed = Vector{speedX, speedY}

		// play fire sound
		sp := shooter.container.getComponent(&SoundPlayer{}).(*SoundPlayer)
		sp.PlaySound(SoundFire)
		// shoot
		bul.position.x = shooter.container.position.x
		bul.position.y = shooter.container.position.y + spriteSize/2
		// the enemy bullet pool is initialized with  speed Vector{0, 2}
		// TODO the speed/direction could be changed with:
		// mover := bul.getComponent(&aimedMover{}).(*aimedMover)
		// mover.speed = Vector{sx, sy}
		bul.active = true
		shooter.lastShot = time.Now()
	}
	return

}
