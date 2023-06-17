package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

// Draws text using it's center point as coordinates
func DrawTextByCenter(screen *ebiten.Image, msg string, font font.Face, cx, cy int, textColor color.Color) {
	bounds := text.BoundString(font, msg)
	x, y := cx-bounds.Min.X-bounds.Dx()/2, cy-bounds.Min.Y-bounds.Dy()/2
	text.Draw(screen, msg, font, x, y, textColor)
}

// Draws text using it's center point as coordinates
func DrawImageByCenter(screen *ebiten.Image, image *ebiten.Image, cx, cy int, op *ebiten.DrawImageOptions) {
	size := image.Bounds().Size()
	op.GeoM.Translate(float64(cx-(size.X/2)), float64(cy-(size.Y/2)))
	screen.DrawImage(image, op)
}
