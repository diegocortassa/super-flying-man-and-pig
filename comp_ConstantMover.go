package main

import "github.com/hajimehoshi/ebiten/v2"

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
	if !mover.active {
		return
	}
	mover.container.position.x += mover.speed.x
	mover.container.position.y += mover.speed.y

	// mover.container.position.x += bulletSpeed * math.Cos(mover.container.rotation)
	// mover.container.position.y += bulletSpeed * math.Sin(mover.container.rotation)

	// entity out fo screen
	if mover.container.position.x > screenWidth+spriteSize || mover.container.position.x < 0 ||
		mover.container.position.y > screenHeight+spriteSize || mover.container.position.y < 0 {
		mover.container.active = false
	}
}

func (k *ConstantMover) Draw(screen *ebiten.Image) {
	return
}
