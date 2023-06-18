package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type ButtonBinds struct {
	Up, Down, Left, Right, Fire, Secondary ebiten.Key
}

type GamePadMover struct {
	active        bool
	container     *Entity
	buttonbinds   ButtonBinds
	speed         Vector
	gamePadNumber int

	bulletsPool []*Entity
	cooldown    time.Duration
	lastShot    time.Time
}

// func NewGamePadMover(container *Entity, buttonBinds ButtonBinds, speed Vector) *GamePadMover {
func NewGamePadMover(container *Entity, gamePadNumber int, speed Vector, bulletsPool []*Entity, cooldown time.Duration) *GamePadMover {
	return &GamePadMover{
		active:        true,
		container:     container,
		gamePadNumber: gamePadNumber,
		speed:         speed,

		bulletsPool: bulletsPool,
		cooldown:    cooldown,
		lastShot:    time.Now(),
	}
}

func (gp *GamePadMover) Update() {
	if !gp.container.active || gp.container.hit {
		return
	}

	// if gp.gamepadIDs == nil {
	// 	gp.gamepadIDs = map[ebiten.GamepadID]struct{}{}
	// }
	// // fmt.Println("GamePadMover update")
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

	gamepadIDs := ebiten.GamepadIDs()

	if len(gamepadIDs) >= gp.gamePadNumber+1 {
		id := gamepadIDs[gp.gamePadNumber]
		if ebiten.IsStandardGamepadLayoutAvailable(id) {
			for b := ebiten.StandardGamepadButton(0); b <= ebiten.StandardGamepadButtonMax; b++ {
				if ebiten.IsStandardGamepadButtonPressed(id, b) {
					switch b {
					case ebiten.StandardGamepadButtonLeftTop:
						if gp.container.position.y > 0 {
							gp.container.position.y -= gp.speed.y
						}
					case ebiten.StandardGamepadButtonLeftLeft:
						if gp.container.position.x > 0 {
							gp.container.position.x -= gp.speed.x
						}
					case ebiten.StandardGamepadButtonLeftRight:
						if gp.container.position.x < screenWidth-spriteSize {
							gp.container.position.x += gp.speed.x
						}
					case ebiten.StandardGamepadButtonLeftBottom:
						if gp.container.position.y < screenHeight-spriteSize {
							gp.container.position.y += gp.speed.y
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

func (gp *GamePadMover) Draw(screen *ebiten.Image) {
	return
}

func (gp *GamePadMover) shoot() {
	if bul, ok := BullettFromPool(gp.bulletsPool); ok {
		bul.active = true
		bul.position.x = gp.container.position.x
		bul.position.y = gp.container.position.y
		// bul.rotation = 270 * (math.Pi / 180)
		gp.lastShot = time.Now()
	}
	return

}
