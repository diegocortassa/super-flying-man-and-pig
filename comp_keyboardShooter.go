package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type keyboardShooter struct {
	active      bool
	container   *Entity
	trigger     ebiten.Key
	bulletsPool []*Entity
	cooldown    time.Duration
	lastShot    time.Time
}

func NewKeyboardShooter(container *Entity, trigger ebiten.Key, bulletsPool []*Entity, cooldown time.Duration) *keyboardShooter {
	return &keyboardShooter{
		active:      true,
		bulletsPool: bulletsPool,
		container:   container,
		trigger:     trigger,
		cooldown:    cooldown,
		lastShot:    time.Now(),
	}
}

func (ks *keyboardShooter) Update() {
	if !ks.container.active || ks.container.hit {
		return
	}
	if ebiten.IsKeyPressed(ks.trigger) && time.Since(ks.lastShot) >= ks.cooldown {
		ks.shoot()
	}
	return
}

func (ks *keyboardShooter) Draw(screen *ebiten.Image) {
	return
}

func (ks *keyboardShooter) shoot() {
	if bul, ok := BullettFromPool(ks.bulletsPool); ok {
		bul.active = true
		bul.position.x = ks.container.position.x
		bul.position.y = ks.container.position.y
		// bul.rotation = 270 * (math.Pi / 180)
		ks.lastShot = time.Now()
		audiosfx_wpn_laser1Player.Rewind()
		audiosfx_wpn_laser1Player.Play()
	}
	return

}
