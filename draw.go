package main

import (
	"image/color"

	"github.com/dcortassa/superflyingmanandpig/assets"
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

// Draws text using it's center point as coordinates with a shadow
func DrawTextWithShadowByCenter(screen *ebiten.Image, msg string, font font.Face, cx, cy int, textColor color.Color) {
	bounds := text.BoundString(font, msg)
	x, y := cx-bounds.Min.X-bounds.Dx()/2, cy-bounds.Min.Y-bounds.Dy()/2
	text.Draw(screen, msg, assets.ArcadeFont, x+1, y+1, color.Black)
	text.Draw(screen, msg, assets.ArcadeFont, x, y, textColor)
}

// Draws text with a shadow
func DrawTextWithShadow(screen *ebiten.Image, msg string, font font.Face, x, y int, textColor color.Color) {
	text.Draw(screen, msg, assets.ArcadeFont, x+1, y+1, color.Black)
	text.Draw(screen, msg, assets.ArcadeFont, x, y, textColor)
}

// Draws image using it's center point as coordinates
func DrawImageByCenter(screen *ebiten.Image, image *ebiten.Image, cx, cy int, op *ebiten.DrawImageOptions) {
	size := image.Bounds().Size()
	op.GeoM.Translate(float64(cx-(size.X/2)), float64(cy-(size.Y/2)))
	screen.DrawImage(image, op)
}
