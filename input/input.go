package input

import (
	"github.com/dcortassa/superflyingmanandpig/globals"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func IsExitJustPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyEscape)
}

func IsResetJustPressed() bool {

	if globals.MameKeys {
		return inpututil.IsKeyJustPressed(ebiten.KeyF1)
	}

	return inpututil.IsKeyJustPressed(ebiten.KeyR)
}

func IsFullScreenJustPressed() bool {
	if globals.MameKeys {
		return inpututil.IsKeyJustPressed(ebiten.KeyF2)
	}
	return inpututil.IsKeyJustPressed(ebiten.KeyF)
}

func IsPauseJustPressed() bool {
	gamepadIDs := ebiten.GamepadIDs()

	for i := 0; i < len(gamepadIDs); i++ {
		// inpututil.IsGamepadButtonJustPressed
		for b := ebiten.StandardGamepadButton(0); b <= ebiten.StandardGamepadButtonMax; b++ {
			if inpututil.IsStandardGamepadButtonJustPressed(gamepadIDs[i], b) {
				if b == ebiten.StandardGamepadButtonCenterLeft {
					return true
				}
			}
		}
	}
	return inpututil.IsKeyJustPressed(ebiten.KeyP)
}

// - (not numeric keypad)
// Volume Down

// = (not numeric keypad)
// Volume Up

// P1

func IsP1FireJustPressed() bool {
	gamepadIDs := ebiten.GamepadIDs()
	if len(gamepadIDs) >= 1 {
		id := gamepadIDs[0]
		for b := ebiten.StandardGamepadButton(0); b <= ebiten.StandardGamepadButtonMax; b++ {
			// if inpututil.IsGamepadButtonJustPressed(id, b) {
			if inpututil.IsStandardGamepadButtonJustPressed(id, b) {
				if b == ebiten.StandardGamepadButtonRightTop ||
					b == ebiten.StandardGamepadButtonRightLeft ||
					b == ebiten.StandardGamepadButtonRightRight ||
					b == ebiten.StandardGamepadButtonRightBottom ||
					b == ebiten.StandardGamepadButtonCenterRight {
					return true
				}
			}
		}
	}
	if globals.MameKeys {
		return inpututil.IsKeyJustPressed(ebiten.KeyControlLeft)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyAltRight) ||
		inpututil.IsKeyJustPressed(ebiten.KeyControlRight) ||
		inpututil.IsKeyJustPressed(ebiten.KeySpace) ||
		inpututil.IsKeyJustPressed(ebiten.KeyAltLeft) ||
		inpututil.IsKeyJustPressed(ebiten.KeyControlLeft) {
		return true
	}
	return false
}

func IsP1FirePressed() bool {
	gamepadIDs := ebiten.GamepadIDs()
	if len(gamepadIDs) >= 1 {
		id := gamepadIDs[0]
		for b := ebiten.StandardGamepadButton(0); b <= ebiten.StandardGamepadButtonMax; b++ {
			// if inpututil.IsGamepadButtonJustPressed(id, b) {
			if inpututil.IsStandardGamepadButtonJustPressed(id, b) {
				if b == ebiten.StandardGamepadButtonRightTop ||
					b == ebiten.StandardGamepadButtonRightLeft ||
					b == ebiten.StandardGamepadButtonRightRight ||
					b == ebiten.StandardGamepadButtonRightBottom ||
					b == ebiten.StandardGamepadButtonCenterRight {
					return true
				}
			}
		}
	}
	if globals.MameKeys {
		return ebiten.IsKeyPressed(ebiten.KeyControlLeft)
	}
	if ebiten.IsKeyPressed(ebiten.KeyAltRight) ||
		ebiten.IsKeyPressed(ebiten.KeyControlRight) ||
		ebiten.IsKeyPressed(ebiten.KeySpace) ||
		ebiten.IsKeyPressed(ebiten.KeyAltLeft) ||
		ebiten.IsKeyPressed(ebiten.KeyControlLeft) {
		return true
	}
	return false

}

func IsP1UpPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyArrowUp)
}

func IsP1DownPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyArrowDown)
}

func IsP1LeftPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyArrowLeft)
}

func IsP1RightPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyArrowRight)
}

// P2
func IsP2FireJustPressed() bool {
	gamepadIDs := ebiten.GamepadIDs()
	if len(gamepadIDs) >= 2 {
		id := gamepadIDs[1]
		for b := ebiten.StandardGamepadButton(0); b <= ebiten.StandardGamepadButtonMax; b++ {
			if inpututil.IsStandardGamepadButtonJustPressed(id, b) {
				if b == ebiten.StandardGamepadButtonRightTop ||
					b == ebiten.StandardGamepadButtonRightLeft ||
					b == ebiten.StandardGamepadButtonRightRight ||
					b == ebiten.StandardGamepadButtonRightBottom ||
					b == ebiten.StandardGamepadButtonCenterRight {
					return true
				}
			}
		}
	}

	if globals.MameKeys {
		return inpututil.IsKeyJustPressed(ebiten.KeyA)
	}
	return inpututil.IsKeyJustPressed(ebiten.KeyQ)
}

func IsP2FirePressed() bool {
	gamepadIDs := ebiten.GamepadIDs()
	if len(gamepadIDs) >= 2 {
		id := gamepadIDs[1]
		for b := ebiten.StandardGamepadButton(0); b <= ebiten.StandardGamepadButtonMax; b++ {
			if inpututil.IsStandardGamepadButtonJustPressed(id, b) {
				if b == ebiten.StandardGamepadButtonRightTop ||
					b == ebiten.StandardGamepadButtonRightLeft ||
					b == ebiten.StandardGamepadButtonRightRight ||
					b == ebiten.StandardGamepadButtonRightBottom ||
					b == ebiten.StandardGamepadButtonCenterRight {
					return true
				}
			}
		}
	}

	if globals.MameKeys {
		return ebiten.IsKeyPressed(ebiten.KeyA)
	}
	return ebiten.IsKeyPressed(ebiten.KeyQ)
}

func IsP2UpPressed() bool {
	if globals.MameKeys {
		return ebiten.IsKeyPressed(ebiten.KeyR)
	}
	return ebiten.IsKeyPressed(ebiten.KeyW)
}

func IsP2DownPressed() bool {
	if globals.MameKeys {
		return ebiten.IsKeyPressed(ebiten.KeyF)
	}
	return ebiten.IsKeyPressed(ebiten.KeyS)
}

func IsP2LeftPressed() bool {
	if globals.MameKeys {
		return ebiten.IsKeyPressed(ebiten.KeyD)
	}
	return ebiten.IsKeyPressed(ebiten.KeyA)
}

func IsP2RightPressed() bool {
	if globals.MameKeys {
		return ebiten.IsKeyPressed(ebiten.KeyG)
	}
	return ebiten.IsKeyPressed(ebiten.KeyD)
}
