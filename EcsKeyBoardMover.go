package main

import (
	"github.com/dcortassa/superflyingmanandpig/assets"
	"github.com/dcortassa/superflyingmanandpig/globals"
	"github.com/hajimehoshi/ebiten/v2"
)

type IsKeyPressed func() bool
type Keybinds struct {
	Up, Down, Left, Right, Fire IsKeyPressed
}

type KeyBoardMover struct {
	active    bool
	container *Entity
	Keybinds  Keybinds
	speed     Vector
}

func NewKeyboardMover(container *Entity, Keybinds Keybinds, speed Vector) *KeyBoardMover {
	return &KeyBoardMover{
		active:    true,
		container: container,
		Keybinds:  Keybinds,
		speed:     speed,
	}
}

func (k *KeyBoardMover) Update() {
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

func (k *KeyBoardMover) Draw(screen *ebiten.Image) {
	return
}
