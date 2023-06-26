package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

// Draws text using it's center point as coordinates
func DrawTextByCenter(screen *ebiten.Image, msg string, font font.Face, cx, cy int, textColor color.Color) {
	bounds := text.BoundString(font, msg)
	x, y := cx-bounds.Min.X-bounds.Dx()/2, cy-bounds.Min.Y-bounds.Dy()/2
	text.Draw(screen, msg, font, x, y, textColor)
}

// Draws text using it's center point as coordinates with a shadow
func DrawTextWithShadowByCenter(screen *ebiten.Image, msg string, font font.Face, cx, cy int, textColor color.Color) {
	bounds := text.BoundString(font, msg)
	x, y := cx-bounds.Min.X-bounds.Dx()/2, cy-bounds.Min.Y-bounds.Dy()/2
	text.Draw(screen, msg, arcadeFont, x+1, y+1, color.Black)
	text.Draw(screen, msg, arcadeFont, x, y, textColor)
}

// Draws text with a shadow
func DrawTextWithShadow(screen *ebiten.Image, msg string, font font.Face, x, y int, textColor color.Color) {
	text.Draw(screen, msg, arcadeFont, x+1, y+1, color.Black)
	text.Draw(screen, msg, arcadeFont, x, y, textColor)
}

// Draws image using it's center point as coordinates
func DrawImageByCenter(screen *ebiten.Image, image *ebiten.Image, cx, cy int, op *ebiten.DrawImageOptions) {
	size := image.Bounds().Size()
	op.GeoM.Translate(float64(cx-(size.X/2)), float64(cy-(size.Y/2)))
	screen.DrawImage(image, op)
}

// Calculate distance between two Entities
func Distance(e1, e2 *Entity) float64 {
	e1CenterX := e1.position.x + spriteSize/2
	e1CenterY := e1.position.y + spriteSize/2
	e2CenterX := e1.position.x + spriteSize/2
	e2CenterY := e1.position.y + spriteSize/2

	dist := math.Sqrt(math.Pow(e1CenterX-e2CenterX, 2) +
		math.Pow(e1CenterY-e2CenterY, 2))

	return dist
}

// keys abstraction

func IsExitJustPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyEscape)
}

func IsResetJustPressed() bool {

	if MameKeys {
		return inpututil.IsKeyJustPressed(ebiten.KeyF1)
	}

	return inpututil.IsKeyJustPressed(ebiten.KeyR)
}

func IsFullScreenJustPressed() bool {
	if MameKeys {
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
	if MameKeys {
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
	if MameKeys {
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

	if MameKeys {
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

	if MameKeys {
		return ebiten.IsKeyPressed(ebiten.KeyA)
	}
	return ebiten.IsKeyPressed(ebiten.KeyQ)
}

func IsP2UpPressed() bool {
	if MameKeys {
		return ebiten.IsKeyPressed(ebiten.KeyR)
	}
	return ebiten.IsKeyPressed(ebiten.KeyW)
}

func IsP2DownPressed() bool {
	if MameKeys {
		return ebiten.IsKeyPressed(ebiten.KeyF)
	}
	return ebiten.IsKeyPressed(ebiten.KeyS)
}

func IsP2LeftPressed() bool {
	if MameKeys {
		return ebiten.IsKeyPressed(ebiten.KeyD)
	}
	return ebiten.IsKeyPressed(ebiten.KeyA)
}

func IsP2RightPressed() bool {
	if MameKeys {
		return ebiten.IsKeyPressed(ebiten.KeyG)
	}
	return ebiten.IsKeyPressed(ebiten.KeyD)
}
