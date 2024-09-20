package main

import (
	"github.com/dcortassa/super-flying-man-and-pig/assets"
	"github.com/dcortassa/super-flying-man-and-pig/globals"
	"github.com/hajimehoshi/ebiten/v2"
)

// Moves an entity at constant velocity
type MoverConstant struct {
	active    bool
	container *Entity
	speed     Vector
}

func NewMoverConstant(container *Entity, speed Vector) *MoverConstant {
	return &MoverConstant{
		active:    true,
		container: container,
		speed:     speed,
	}
}

func (mover *MoverConstant) Update() {
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

func (k *MoverConstant) Draw(screen *ebiten.Image) {
	// mover doesn't need to be drawn
}
