package main

import (
	"image"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type animator struct {
	container       *Entity
	sequences       map[string]*sequence
	current         string
	lastFrameChange time.Time
	finished        bool
}

func newAnimator(container *Entity, sequences map[string]*sequence, defaultSequence string) *animator {
	var an animator

	an.container = container
	an.sequences = sequences
	an.current = defaultSequence
	an.lastFrameChange = time.Now()

	return &an
}

func (an *animator) Update() {
	sequence := an.sequences[an.current]

	frameInterval := float64(time.Second) / sequence.sampleRate

	if time.Since(an.lastFrameChange) >= time.Duration(frameInterval) {
		an.finished = sequence.nextFrame()
		an.lastFrameChange = time.Now()
	}
}

func (an *animator) Draw(screen *ebiten.Image) {

	frameOffset := an.sequences[an.current].frames[an.sequences[an.current].frame] * spriteSize
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(an.container.position.x, an.container.position.y)
	screen.DrawImage(an.sequences[an.current].spriteSheet.SubImage(image.Rect(frameOffset, 0, frameOffset+spriteSize, spriteSize)).(*ebiten.Image), opts)

	if debug {
		for _, b := range an.container.hitBoxes {
			vector.DrawFilledRect(screen, float32(an.container.position.x+b.x), float32(an.container.position.y+b.y), float32(b.w), float32(b.h), color.RGBA{0xff, 0, 0, 0xff}, false)
		}
	}

}

type sequence struct {
	spriteSheet *ebiten.Image
	frames      []int
	frame       int
	sampleRate  float64
	loop        bool
}

func newSequence(spriteSheet *ebiten.Image, frames []int, sampleRate float64, loop bool) *sequence {

	var seq sequence
	seq.spriteSheet = spriteSheet
	seq.frames = frames
	seq.sampleRate = sampleRate
	seq.loop = loop

	return &seq
}

func (seq *sequence) nextFrame() bool {
	if seq.frame == len(seq.frames)-1 {
		if seq.loop {
			seq.frame = 0
		} else {
			return true
		}
	} else {
		seq.frame++
	}
	return false
}
