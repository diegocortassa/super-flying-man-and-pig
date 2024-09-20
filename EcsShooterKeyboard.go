package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type ShooterKeyboard struct {
	active      bool
	container   *Entity
	trigger     IsKeyPressed
	bulletsPool []*Entity
	cooldown    time.Duration
	lastShot    time.Time
}

func NewShooterKeyboard(container *Entity, trigger IsKeyPressed, bulletsPool []*Entity, cooldown time.Duration) *ShooterKeyboard {
	return &ShooterKeyboard{
		active:      true,
		bulletsPool: bulletsPool,
		container:   container,
		trigger:     trigger,
		cooldown:    cooldown,
		lastShot:    time.Now(),
	}
}

func (ks *ShooterKeyboard) Update() {
	if !ks.container.Active || ks.container.Hit {
		return
	}
	if ks.trigger() && time.Since(ks.lastShot) >= ks.cooldown {
		ks.shoot()
	}
}

func (ks *ShooterKeyboard) Draw(screen *ebiten.Image) {
	// shooter doesn't need to be drawn
}

func (ks *ShooterKeyboard) shoot() {
	if bul, ok := BullettFromPool(ks.bulletsPool); ok {
		// Play fire sound
		sp := ks.container.GetComponent(&SoundPlayer{}).(*SoundPlayer)
		sp.PlaySound(SoundFire)
		// shoot
		bul.Position.X = ks.container.Position.X
		bul.Position.Y = ks.container.Position.Y
		bul.Active = true
		// bul.rotation = 270 * (math.Pi / 180)
		ks.lastShot = time.Now()
	}
}
