package main

import (
	"bytes"
	_ "embed"
	"image"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	enemyFlyingMan1Size = 12
	// EnemyFlyingMan1Speed     = 1
	EnemyFlyingMan1AnimSpeed = 4
	EnemyFlyingMan1FramesNum = 5
	enemyFlyingMan1FrameSize = 24
)

type EnemyFlyingMan1 struct {
	x, y                 float64
	animDir              int
	currentFrame         int
	frameCounter         int
	EnemyFlyingMan1Speed float64
}

var (
	EnemyFlyingMan1Image *ebiten.Image
	//go:embed assets/sprites/EnemyFlyingMan1.png
	EnemyFlyingMan1_png []byte
)

func newEnemyFlyingMan1(x, y float64) (p EnemyFlyingMan1) {
	p.x = x
	p.y = y
	// p.EnemyFlyingMan1Speed = rand.Float64()*2 - 1
	p.EnemyFlyingMan1Speed = rand.Float64()
	// p.x = (screenWidth - enemyFlyingMan1FrameSize) / 2
	// p.y = enemyFlyingMan1FrameSize + 20
	p.animDir = 1 // Start with forward direction

	// Decode spritesheet from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(EnemyFlyingMan1_png))
	if err != nil {
		log.Fatal(err)
	}
	EnemyFlyingMan1Image = ebiten.NewImageFromImage(img)

	return p
}

func (p *EnemyFlyingMan1) Update(g *Game) {
	p.y += p.EnemyFlyingMan1Speed
	// if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
	// 	p.y -= EnemyFlyingMan1Speed
	// }
	// if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
	// 	p.y += EnemyFlyingMan1Speed
	// }
	// if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
	// 	p.x -= EnemyFlyingMan1Speed
	// }
	// if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
	// 	p.x += EnemyFlyingMan1Speed
	// }

	// if p.x < 0 {
	// 	p.x = 0
	// } else if p.x > screenWidth-enemyFlyingMan1FrameSize {
	// 	p.x = screenWidth - enemyFlyingMan1FrameSize
	// }

	// if p.y < 0 {
	// 	p.y = 0
	// } else if p.y > screenHeight-enemyFlyingMan1FrameSize {
	// 	p.y = screenHeight - enemyFlyingMan1FrameSize
	// }

	// compute frame ping pong animation
	p.frameCounter++
	if p.frameCounter >= EnemyFlyingMan1AnimSpeed {
		p.currentFrame += p.animDir
		if p.currentFrame >= EnemyFlyingMan1FramesNum-1 || p.currentFrame <= 0 {
			p.animDir *= -1 // Reverse direction
		}
		p.frameCounter = 0
	}
}

func (p *EnemyFlyingMan1) Draw(screen *ebiten.Image) {
	frameOffset := p.currentFrame * enemyFlyingMan1FrameSize
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(p.x, p.y)
	screen.DrawImage(EnemyFlyingMan1Image.SubImage(image.Rect(frameOffset, 0, frameOffset+enemyFlyingMan1FrameSize, enemyFlyingMan1FrameSize)).(*ebiten.Image), opts)
}
