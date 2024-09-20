package main

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type ShooterRotative struct {
	active        bool
	container     *Entity
	trigger       time.Duration
	pool          []*Entity
	lastShot      time.Time
	lastDirection int
	rotationSpeed int
}

func NewShooterRotative(container *Entity, trigger time.Duration, direction int, rotationSpeed int, pool []*Entity) *ShooterRotative {
	return &ShooterRotative{
		active:        true,
		pool:          pool,
		container:     container,
		trigger:       trigger,
		lastShot:      time.Now(),
		lastDirection: direction,
		rotationSpeed: rotationSpeed,
	}
}

func (shooter *ShooterRotative) Update() {
	if time.Since(shooter.lastShot) >= shooter.trigger && !shooter.container.Hit {
		shooter.shoot(shooter.container.Position.X+25, shooter.container.Position.Y-20)
		shooter.lastShot = time.Now()
	}
	return
}

func (shooter *ShooterRotative) Draw(screen *ebiten.Image) {
	// shooter doesn't need to be drawn
}

// Shoot bullet from pool starting at position x,y
func (shooter *ShooterRotative) shoot(x, y float64) {
	if bul, ok := BullettFromPool(shooter.pool); ok {
		// do not shoot while exploding
		if shooter.container.Exploding {
			return
		}
		// play fire sound
		sp := shooter.container.GetComponent(&SoundPlayer{}).(*SoundPlayer)
		sp.PlaySound(SoundFire)
		// shoot
		bul.Position.X = shooter.container.Position.X
		// This is used by volcanoes the bullet origins at sprite center
		// bul.Position.Y = shooter.container.Position.Y + assets.SpriteSize/2
		bul.Position.Y = shooter.container.Position.Y

		rotationRad := float64(shooter.lastDirection) * (math.Pi / 180.0)
		X := math.Cos(rotationRad)
		Y := math.Sin(rotationRad)
		shooter.lastDirection += shooter.rotationSpeed

		mover := bul.GetComponent(&MoverConstant{}).(*MoverConstant)
		mover.speed = Vector{X, Y}

		bul.Active = true
		shooter.lastShot = time.Now()
	}
}
