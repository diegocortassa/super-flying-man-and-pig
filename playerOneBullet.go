package main

import (
	"bytes"
	_ "embed"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	playerOneBulletSize      = 12
	playerOneBulletSpeed     = -4
	playerOneBulletAnimSpeed = 10.0
	playerOneBulletFramesNum = 5
	playerOneBulletFrameSize = 24
	playerOneBulletPoolSize  = 5
)

type PlayerOneBullet struct {
	x, y         float64
	animDir      int
	currentFrame int
	frameCounter int
	active       bool
}

var (
	playerOneBulletBulletImage *ebiten.Image
	//go:embed assets/sprites/SuperFlyingManPew.png
	playerOneBulletBullet_png []byte
)

func newPlayerOneBullet(x, y float64) (p PlayerOneBullet) {
	p.x = x
	p.y = y
	p.animDir = 1 // Start with forward direction
	p.active = false

	// Decode spritesheet from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(playerOneBulletBullet_png))
	if err != nil {
		log.Fatal(err)
	}
	playerOneBulletBulletImage = ebiten.NewImageFromImage(img)

	return p
}

func (p *PlayerOneBullet) Update(g *Game) {
	if !p.active {
		return
	}

	p.y += playerOneBulletSpeed

	if p.y+float64(playerOneBulletFrameSize) < 0 { // bullet out fo screen
		p.y = 0
		p.active = false
	} else if p.y > screenHeight-playerOneBulletFrameSize {
		p.y = screenHeight - playerOneBulletFrameSize
	}

	// compute frame ping pong animation
	p.frameCounter++
	if p.frameCounter >= playerOneBulletAnimSpeed {
		p.currentFrame += p.animDir
		if p.currentFrame >= playerOneBulletFramesNum-1 || p.currentFrame <= 0 {
			p.animDir *= -1 // Reverse direction
		}
		p.frameCounter = 0
	}
}

func (p *PlayerOneBullet) Draw(screen *ebiten.Image) {
	if !p.active {
		return
	}

	frameOffset := p.currentFrame * playerOneBulletFrameSize
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(p.x, p.y)
	screen.DrawImage(playerOneBulletBulletImage.SubImage(image.Rect(frameOffset, 0, frameOffset+playerOneBulletFrameSize, playerOneBulletFrameSize)).(*ebiten.Image), opts)
}

func initPlayerOneBullettPool(g *Game) {
	for i := 0; i < playerOneBulletPoolSize; i++ {
		bul := newPlayerOneBullet(0, 0)
		g.playerOneBullettPool = append(g.playerOneBullettPool, &bul)
	}
}

func playerOneBullettFromPool(g *Game) (*PlayerOneBullet, bool) {
	for _, bul := range g.playerOneBullettPool {
		if !bul.active {
			return bul, true
		}
	}
	return nil, false
}
