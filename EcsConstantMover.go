package main

import (
	"github.com/dcortassa/super-flying-man-and-pig/assets"
	"github.com/dcortassa/super-flying-man-and-pig/globals"
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
	if !mover.active || mover.container.Hit {
		return
	}
	mover.container.Position.X += mover.speed.X
	mover.container.Position.Y += mover.speed.Y

	// entity out of screen
	if mover.container.Position.X > globals.ScreenWidth || mover.container.Position.X+assets.SpriteSize < 0 ||
		mover.container.Position.Y > globals.ScreenHeight || mover.container.Position.Y+assets.SpriteSize < 0 {
		mover.container.Active = false
	}
}

func (k *ConstantMover) Draw(screen *ebiten.Image) {
	return
}
