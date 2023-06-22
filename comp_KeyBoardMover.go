package main

import "github.com/hajimehoshi/ebiten/v2"

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
	if !k.container.active || k.container.hit {
		return
	}
	if k.Keybinds.Up() && k.container.position.y > 0 {
		k.container.position.y -= k.speed.y
	}
	if k.Keybinds.Down() && k.container.position.y < screenHeight-spriteSize {
		k.container.position.y += k.speed.y
	}
	if k.Keybinds.Left() && k.container.position.x > 0 {
		k.container.position.x -= k.speed.x
	}
	if k.Keybinds.Right() && k.container.position.x < screenWidth-spriteSize {
		k.container.position.x += k.speed.x
	}
}

func (k *KeyBoardMover) Draw(screen *ebiten.Image) {
	return
}
