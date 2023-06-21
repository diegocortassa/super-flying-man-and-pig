package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type constantShooter struct {
	active    bool
	container *Entity
	trigger   time.Duration
	pool      []*Entity
	lastShot  time.Time
}

func NewConstantShooter(container *Entity, trigger time.Duration, pool []*Entity) *constantShooter {
	return &constantShooter{
		active:    true,
		pool:      pool,
		container: container,
		trigger:   trigger,
		lastShot:  time.Now(),
	}
}

func (shooter *constantShooter) Update() {
	if time.Since(shooter.lastShot) >= shooter.trigger {
		shooter.shoot(shooter.container.position.x+25, shooter.container.position.y-20)
		shooter.lastShot = time.Now()
	}
	return
}

func (shooter *constantShooter) Draw(screen *ebiten.Image) {
	return
}

// Shoot bullet from pool starting at position x,y
func (shooter *constantShooter) shoot(x, y float64) {
	if bul, ok := BullettFromPool(shooter.pool); ok {
		// do not shoot while exploding
		if shooter.container.exploding {
			return
		}
		// play fire sound
		sp := shooter.container.getComponent(&SoundPlayer{}).(*SoundPlayer)
		sp.PlaySound(SoundFire)
		// shoot
		bul.position.x = shooter.container.position.x
		bul.position.y = shooter.container.position.y + spriteSize/2
		// reset the bullet speed/direction to Vector{0, 2}
		// TODO the speed/direction could be changed with:
		mover := bul.getComponent(&ConstantMover{}).(*ConstantMover)
		mover.speed = Vector{0, 2}
		bul.active = true
		shooter.lastShot = time.Now()
	}
	return

}
