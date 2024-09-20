package main

import (
	"time"

	"github.com/dcortassa/super-flying-man-and-pig/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type ShooterConstant struct {
	active    bool
	container *Entity
	trigger   time.Duration
	pool      []*Entity
	lastShot  time.Time
}

func NewShooterConstant(container *Entity, trigger time.Duration, pool []*Entity) *ShooterConstant {
	return &ShooterConstant{
		active:    true,
		pool:      pool,
		container: container,
		trigger:   trigger,
		lastShot:  time.Now(),
	}
}

func (shooter *ShooterConstant) Update() {
	if time.Since(shooter.lastShot) >= shooter.trigger && !shooter.container.Hit {
		shooter.shoot(shooter.container.Position.X+25, shooter.container.Position.Y-20)
		shooter.lastShot = time.Now()
	}
}

func (shooter *ShooterConstant) Draw(screen *ebiten.Image) {
	// shooter doesn't need to be drawn
}

// Shoot bullet from pool starting at position x,y
func (shooter *ShooterConstant) shoot(x, y float64) {
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
		bul.Position.Y = shooter.container.Position.Y + assets.SpriteSize/2
		// reset the bullet speed/direction to Vector{0, 2}
		// TODO the speed/direction could be changed with:
		mover := bul.GetComponent(&MoverConstant{}).(*MoverConstant)
		mover.speed = Vector{0, 2}
		bul.Active = true
		shooter.lastShot = time.Now()
	}
}
