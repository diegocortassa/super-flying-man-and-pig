package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ConstantMover struct {
	active    bool
	container *Entity
	speed     Vector
}

func NewConstantMover(container *Entity, speed Vector) *ConstantMover {
	return &ConstantMover{
		active:    true,
		container: container,
		speed:     speed,
	}
}

func (mover *ConstantMover) Update() {
	if !mover.active || mover.container.hit {
		return
	}
	mover.container.position.x += mover.speed.x
	mover.container.position.y += mover.speed.y

	// if we want to use rotation to aim movement
	// mover.container.position.x += mover.speed.x * math.Cos(mover.container.rotation)
	// mover.container.position.y += mover.speed.y * math.Sin(mover.container.rotation)

	// entity out fo screen
	if mover.container.position.x > screenWidth+spriteSize || mover.container.position.x+spriteSize < 0 ||
		mover.container.position.y > screenHeight+spriteSize || mover.container.position.y+spriteSize < 0 {
		mover.container.active = false
	}
}

func (k *ConstantMover) Draw(screen *ebiten.Image) {
	return
}
