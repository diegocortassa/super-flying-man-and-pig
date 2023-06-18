package main

import (
	"image/color"

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

// Draws image using it's center point as coordinates
func DrawImageByCenter(screen *ebiten.Image, image *ebiten.Image, cx, cy int, op *ebiten.DrawImageOptions) {
	size := image.Bounds().Size()
	op.GeoM.Translate(float64(cx-(size.X/2)), float64(cy-(size.Y/2)))
	screen.DrawImage(image, op)
}

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
	return inpututil.IsKeyJustPressed(ebiten.KeyControlLeft)
}

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
	return inpututil.IsKeyJustPressed(ebiten.KeyQ)
}

func IsExitJustPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyEscape)
}

func IsResetJustPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyR)
}

func IsFullScreenJustPressed() bool {
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
	return inpututil.IsKeyJustPressed(ebiten.KeyF)
}
