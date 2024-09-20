package main

import (
	"github.com/dcortassa/super-flying-man-and-pig/assets"
	"github.com/dcortassa/super-flying-man-and-pig/globals"
	"github.com/hajimehoshi/ebiten/v2"
)

type IsKeyPressed func() bool
type Keybinds struct {
	Up, Down, Left, Right, Fire IsKeyPressed
}

type MoverKeyBoard struct {
	active    bool
	container *Entity
	Keybinds  Keybinds
	speed     Vector
}

func NewMoverKeyboard(container *Entity, Keybinds Keybinds, speed Vector) *MoverKeyBoard {
	return &MoverKeyBoard{
		active:    true,
		container: container,
		Keybinds:  Keybinds,
		speed:     speed,
	}
}

func (k *MoverKeyBoard) Update() {
	if !k.container.Active || k.container.Hit {
		return
	}
	if k.Keybinds.Up() && k.container.Position.Y > 0 {
		k.container.Position.Y -= k.speed.Y
	}
	if k.Keybinds.Down() && k.container.Position.Y < globals.ScreenHeight-assets.SpriteSize {
		k.container.Position.Y += k.speed.Y
	}
	if k.Keybinds.Left() && k.container.Position.X > 0 {
		k.container.Position.X -= k.speed.X
	}
	if k.Keybinds.Right() && k.container.Position.X < globals.ScreenWidth-assets.SpriteSize {
		k.container.Position.X += k.speed.X
	}
}

func (k *MoverKeyBoard) Draw(screen *ebiten.Image) {
	// mover doesn't need to be drawn
}
