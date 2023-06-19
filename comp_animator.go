package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type animator struct {
	container       *Entity
	sequences       map[string]*sequence
	currentSeq      string
	lastFrameChange time.Time
	finished        bool
	// mu              sync.Mutex
}

func newAnimator(container *Entity, sequences map[string]*sequence, defaultSequence string) *animator {
	var an animator

	an.container = container
	an.sequences = sequences
	an.currentSeq = defaultSequence
	an.lastFrameChange = time.Now()

	return &an
}

func (an *animator) Update() {
	sequence := an.sequences[an.currentSeq]

	// TODO Collision logic should not be here !!
	if an.container.hit && an.currentSeq == "idle" {
		an.sequences["destroy"].currentFrame = 0
		an.currentSeq = "destroy"
		sfx_exp_odd1Player.Rewind()
		sfx_exp_odd1Player.Play()
	}
	if an.container.hit && an.finished {
		an.container.hit = false
		an.container.invulnerable = false
		an.currentSeq = "idle"
		if an.container.lives < 1 {
			an.container.active = false
		}

	}

	frameInterval := float64(time.Second) / sequence.fps

	if time.Since(an.lastFrameChange) >= time.Duration(frameInterval) {
		an.finished = sequence.nextFrame()
		an.lastFrameChange = time.Now()
	}
}

func (an *animator) Draw(screen *ebiten.Image) {

	frameIndex := an.sequences[an.currentSeq].frames[an.sequences[an.currentSeq].currentFrame]
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(an.container.position.x, an.container.position.y)
	screen.DrawImage(getSprite(frameIndex).(*ebiten.Image), opts)

	if debug {
		for _, b := range an.container.hitBoxes {
			vector.DrawFilledRect(screen, float32(an.container.position.x+b.x), float32(an.container.position.y+b.y), float32(b.w), float32(b.h), color.RGBA{0xff, 0, 0, 0xff}, false)
		}
	}

}

func (an *animator) setSequence(name string) {
	an.currentSeq = name
	an.lastFrameChange = time.Now()
}

type sequence struct {
	frames       []int
	currentFrame int
	fps          float64
	loop         bool
}

func newSequence(frames []int, fps float64, loop bool) *sequence {

	var seq sequence
	seq.frames = frames
	seq.fps = fps
	seq.loop = loop

	return &seq
}

func (seq *sequence) nextFrame() bool {
	if seq.currentFrame == len(seq.frames)-1 {
		if seq.loop {
			seq.currentFrame = 0
		} else {
			return true
		}
	} else {
		seq.currentFrame++
	}
	return false
}
