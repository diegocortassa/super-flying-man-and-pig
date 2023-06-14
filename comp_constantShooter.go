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

func (shooter *constantShooter) shoot(x, y float64) {
	if bul, ok := BullettFromPool(shooter.pool); ok {
		bul.active = true
		bul.position.x = shooter.container.position.x
		bul.position.y = shooter.container.position.y
		// mover := bul.getComponent(&ConstantMover{})
		// mover := bul.components[0]
		// mover.speed = Vector{1, 0}
		shooter.lastShot = time.Now()
	}
	return

}
