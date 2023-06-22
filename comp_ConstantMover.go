package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Moves an entity at constant velocity
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

	// entity out of screen
	if mover.container.position.x > screenWidth || mover.container.position.x+spriteSize < 0 ||
		mover.container.position.y > screenHeight || mover.container.position.y+spriteSize < 0 {
		mover.container.active = false
	}
}

func (k *ConstantMover) Draw(screen *ebiten.Image) {
	return
}
