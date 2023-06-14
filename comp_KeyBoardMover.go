package main

import "github.com/hajimehoshi/ebiten/v2"

type Keybinds struct {
	Up, Down, Left, Right, Fire, Secondary ebiten.Key
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
	if !k.active {
		return
	}
	// if ebiten.IsKeyPressed(k.Keybinds.fire) && time.Since(k.container.lastShoot) >= k.container.shotCoolDown {
	// 	k.container.Shoot(g)
	// }
	if ebiten.IsKeyPressed(k.Keybinds.Up) && k.container.position.y > 0 {
		k.container.position.y -= k.speed.y
	}
	if ebiten.IsKeyPressed(k.Keybinds.Down) && k.container.position.y < screenHeight-spriteSize {
		k.container.position.y += k.speed.y
	}
	if ebiten.IsKeyPressed(k.Keybinds.Left) && k.container.position.x > 0 {
		k.container.position.x -= k.speed.x
	}
	if ebiten.IsKeyPressed(k.Keybinds.Right) && k.container.position.x < screenWidth-spriteSize {
		k.container.position.x += k.speed.x
	}
}

func (k *KeyBoardMover) Draw(screen *ebiten.Image) {
	return
}
