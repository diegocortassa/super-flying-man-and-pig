package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type keyboardShooter struct {
	active    bool
	container *Entity
	trigger   ebiten.Key
	pool      []*Entity
	cooldown  time.Duration
	lastShot  time.Time
}

func NewKeyboardShooter(container *Entity, trigger ebiten.Key, pool []*Entity, cooldown time.Duration) *keyboardShooter {
	return &keyboardShooter{
		active:    true,
		pool:      pool,
		container: container,
		trigger:   trigger,
		cooldown:  cooldown,
		lastShot:  time.Now(),
	}
}

func (ks *keyboardShooter) Update() {
	if ebiten.IsKeyPressed(ks.trigger) && time.Since(ks.lastShot) >= ks.cooldown {
		ks.shoot(ks.container.position.x+25, ks.container.position.y-20)
		ks.lastShot = time.Now()
	}
	return
}

func (ks *keyboardShooter) Draw(screen *ebiten.Image) {
	return
}

func (ks *keyboardShooter) shoot(x, y float64) {
	// if bul, ok := playerOneBullettFromPool(g); ok {
	if bul, ok := BullettFromPool(ks.pool); ok {
		bul.active = true
		bul.position.x = ks.container.position.x
		bul.position.y = ks.container.position.y
		// bul.rotation = 270 * (math.Pi / 180)
		ks.lastShot = time.Now()
	}
	return

}
