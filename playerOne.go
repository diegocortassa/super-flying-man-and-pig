package main

import (
	"bytes"
	_ "embed"
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	playerOneSize         = 12
	playerOneSpeed        = 1
	playerOneAnimSpeed    = 4
	playerOneFramesNum    = 5
	playerOneFrameSize    = 24
	playerOneShotCoolDown = time.Millisecond * 250
)

type PlayerOne struct {
	x, y         float64
	animDir      int
	currentFrame int
	frameCounter int
	lastShoot    time.Time
}

var (
	SuperFlyingManImage *ebiten.Image
	//go:embed assets/sprites/SuperFlyingMan.png
	SuperFlyingMan_png []byte
)

func newPlayer() (p PlayerOne) {
	p.x = (screenWidth - playerOneFrameSize) / 4
	p.y = screenHeight - playerOneFrameSize - 20
	p.animDir = 1 // Start with forward direction

	// Decode spritesheet from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(SuperFlyingMan_png))
	if err != nil {
		log.Fatal(err)
	}
	SuperFlyingManImage = ebiten.NewImageFromImage(img)

	return p
}

func (p *PlayerOne) Update(g *Game) {

	if ebiten.IsKeyPressed(ebiten.KeyControlLeft) && time.Since(p.lastShoot) >= playerOneShotCoolDown {
		p.Shoot(g)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && p.y > 0 {
		p.y -= playerOneSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && p.y < screenHeight-playerOneFrameSize {
		p.y += playerOneSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && p.x > 0 {
		p.x -= playerOneSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && p.x < screenWidth-playerOneFrameSize {
		p.x += playerOneSpeed
	}

	// compute frame ping pong animation
	p.frameCounter++
	if p.frameCounter >= playerOneAnimSpeed {
		p.currentFrame += p.animDir
		if p.currentFrame >= playerOneAnimSpeed-1 || p.currentFrame <= 0 {
			p.animDir *= -1 // Reverse direction
		}
		p.frameCounter = 0
	}
}

func (p *PlayerOne) Draw(screen *ebiten.Image) {
	frameOffset := p.currentFrame * playerOneFrameSize
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(p.x, p.y)
	screen.DrawImage(SuperFlyingManImage.SubImage(image.Rect(frameOffset, 0, frameOffset+playerOneFrameSize, playerOneFrameSize)).(*ebiten.Image), opts)
}

func (p *PlayerOne) Shoot(g *Game) {
	if bul, ok := playerOneBullettFromPool(g); ok {
		bul.active = true
		bul.x = p.x
		bul.y = p.y
		p.lastShoot = time.Now()
	}
}
