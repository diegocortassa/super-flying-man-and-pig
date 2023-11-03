package main

import (
	"math"
	"time"

	"github.com/dcortassa/superflyingmanandpig/assets"
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
	if time.Since(shooter.lastShot) >= shooter.trigger && !shooter.container.Hit {
		shooter.shoot(shooter.container.Position.X+25, shooter.container.Position.Y-20)
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
		if shooter.container.Exploding {
			return
		}

		mover := bul.GetComponent(&ConstantMover{}).(*ConstantMover)

		var px, py float64
		distP1 := distance(shooter.container, shooter.p1)
		distP2 := distance(shooter.container, shooter.p2)

		// find nearest player
		if !shooter.p2.Active || distP1 < distP2 {
			px = shooter.p1.Position.X
			py = shooter.p1.Position.Y
		} else {
			px = shooter.p2.Position.X
			py = shooter.p2.Position.Y
		}

		sx := shooter.container.Position.X
		sy := shooter.container.Position.Y

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
		sp := shooter.container.GetComponent(&SoundPlayer{}).(*SoundPlayer)
		sp.PlaySound(SoundFire)
		// shoot
		bul.Position.X = shooter.container.Position.X
		bul.Position.Y = shooter.container.Position.Y + assets.SpriteSize/2
		// the enemy bullet pool is initialized with  speed Vector{0, 2}
		// TODO the speed/direction could be changed with:
		// mover := bul.GetComponent(&aimedMover{}).(*aimedMover)
		// mover.speed = Vector{sx, sy}
		bul.Active = true
		shooter.lastShot = time.Now()
	}
	return

}

// Calculate distance between two Entities
func distance(e1, e2 *Entity) float64 {
	e1CenterX := e1.Position.X + assets.SpriteSize/2
	e1CenterY := e1.Position.Y + assets.SpriteSize/2
	e2CenterX := e1.Position.X + assets.SpriteSize/2
	e2CenterY := e1.Position.Y + assets.SpriteSize/2

	dist := math.Sqrt(math.Pow(e1CenterX-e2CenterX, 2) +
		math.Pow(e1CenterY-e2CenterY, 2))

	return dist
}
