package main

import (
	"time"

	"github.com/dcortassa/super-flying-man-and-pig/assets"
	"github.com/dcortassa/super-flying-man-and-pig/globals"
	"github.com/hajimehoshi/ebiten/v2"
)

type ButtonBinds struct {
	Up, Down, Left, Right, Fire, Secondary ebiten.Key
}

type MoverGamePad struct {
	active        bool
	container     *Entity
	buttonbinds   ButtonBinds
	speed         Vector
	gamePadNumber int
	gamepadIDsBuf []ebiten.GamepadID
	gamepadIDs    map[ebiten.GamepadID]struct{}

	bulletsPool []*Entity
	cooldown    time.Duration
	lastShot    time.Time
}

// func NewMoverGamePad(container *Entity, buttonBinds ButtonBinds, speed Vector) *MoverGamePad {
func NewMoverGamePad(container *Entity, gamePadNumber int, speed Vector, bulletsPool []*Entity, cooldown time.Duration) *MoverGamePad {
	return &MoverGamePad{
		active:        true,
		container:     container,
		gamePadNumber: gamePadNumber,
		speed:         speed,

		bulletsPool: bulletsPool,
		cooldown:    cooldown,
		lastShot:    time.Now(),
	}
}

func (gp *MoverGamePad) Update() {
	if !gp.container.Active || gp.container.Hit {
		return
	}

	// if gp.gamepadIDs == nil {
	// 	gp.gamepadIDs = map[ebiten.GamepadID]struct{}{}
	// }
	// gp.gamepadIDsBuf = inpututil.AppendJustConnectedGamepadIDs(gp.gamepadIDsBuf[:0])
	// for _, id := range gp.gamepadIDsBuf {
	// 	fmt.Printf("gamepad connected: id: %d, SDL ID: %s", id, ebiten.GamepadSDLID(id))
	// 	gp.gamepadIDs[id] = struct{}{}
	// }
	// for id := range gp.gamepadIDs {
	// 	if inpututil.IsGamepadJustDisconnected(id) {
	// 		fmt.Printf("gamepad disconnected: id: %d", id)
	// 		delete(gp.gamepadIDs, id)
	// 	}
	// }

	// gamepadIDs := gp.gamepadIDs
	gamepadIDs := ebiten.GamepadIDs()

	if len(gamepadIDs) >= gp.gamePadNumber+1 {
		id := gamepadIDs[gp.gamePadNumber]
		if ebiten.IsStandardGamepadLayoutAvailable(id) {
			for b := ebiten.StandardGamepadButton(0); b <= ebiten.StandardGamepadButtonMax; b++ {
				if ebiten.IsStandardGamepadButtonPressed(id, b) {
					switch b {
					case ebiten.StandardGamepadButtonLeftTop:
						if gp.container.Position.Y > 0 {
							gp.container.Position.Y -= gp.speed.Y
						}
					case ebiten.StandardGamepadButtonLeftLeft:
						if gp.container.Position.X > 0 {
							gp.container.Position.X -= gp.speed.X
						}
					case ebiten.StandardGamepadButtonLeftRight:
						if gp.container.Position.X < globals.ScreenWidth-assets.SpriteSize {
							gp.container.Position.X += gp.speed.X
						}
					case ebiten.StandardGamepadButtonLeftBottom:
						if gp.container.Position.Y < globals.ScreenHeight-assets.SpriteSize {
							gp.container.Position.Y += gp.speed.Y
						}
					case ebiten.StandardGamepadButtonRightTop,
						ebiten.StandardGamepadButtonRightLeft,
						ebiten.StandardGamepadButtonRightRight,
						ebiten.StandardGamepadButtonRightBottom:
						if time.Since(gp.lastShot) >= gp.cooldown {
							gp.shoot()
						}

					}

				}

			}
		}
	}

}

func (gp *MoverGamePad) Draw(screen *ebiten.Image) {
	// mover doesn't need to be drawn
}

func (gp *MoverGamePad) shoot() {
	if bul, ok := BullettFromPool(gp.bulletsPool); ok {
		bul.Active = true
		bul.Position.X = gp.container.Position.X
		bul.Position.Y = gp.container.Position.Y
		// bul.rotation = 270 * (math.Pi / 180)
		gp.lastShot = time.Now()
	}
}
