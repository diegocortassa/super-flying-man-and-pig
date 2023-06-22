package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type KeyboardShooter struct {
	active      bool
	container   *Entity
	trigger     IsKeyPressed
	bulletsPool []*Entity
	cooldown    time.Duration
	lastShot    time.Time
}

func NewKeyboardShooter(container *Entity, trigger IsKeyPressed, bulletsPool []*Entity, cooldown time.Duration) *KeyboardShooter {
	return &KeyboardShooter{
		active:      true,
		bulletsPool: bulletsPool,
		container:   container,
		trigger:     trigger,
		cooldown:    cooldown,
		lastShot:    time.Now(),
	}
}

func (ks *KeyboardShooter) Update() {
	if !ks.container.active || ks.container.hit {
		return
	}
	if ks.trigger() && time.Since(ks.lastShot) >= ks.cooldown {
		ks.shoot()
	}
	return
}

func (ks *KeyboardShooter) Draw(screen *ebiten.Image) {
	return
}

func (ks *KeyboardShooter) shoot() {
	if bul, ok := BullettFromPool(ks.bulletsPool); ok {
		// Play fire sound
		sp := ks.container.getComponent(&SoundPlayer{}).(*SoundPlayer)
		sp.PlaySound(SoundFire)
		// shoot
		bul.position.x = ks.container.position.x
		bul.position.y = ks.container.position.y
		bul.active = true
		// bul.rotation = 270 * (math.Pi / 180)
		ks.lastShot = time.Now()
	}
	return

}
