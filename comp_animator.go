package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
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

	// If container was hit play explosion
	if an.container.hit && an.currentSeq == "idle" {
		// in no sequence destroy exists enemy just disappear
		if _, ok := an.sequences["destroy"]; ok {
			an.sequences["destroy"].currentFrame = 0
			an.currentSeq = "destroy"
		} else {
			an.container.active = false
		}
		sp := an.container.getComponent(&SoundPlayer{}).(*SoundPlayer)
		if sp != nil {
			sp.PlaySound(SoundDestroy)

		}
	}

	// TODO Collision logic should not be here !!
	if an.container.hit && an.finished {
		an.container.invulnerable = true
		an.container.invulnerableSetTime = time.Now()
		an.container.hit = false
		an.currentSeq = "idle"
		if an.container.lives < 1 {
			an.container.active = false
		}
	}
	if an.container.invulnerable && time.Since(an.container.invulnerableSetTime) > time.Second*3 {
		an.container.invulnerable = false
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

	if an.container.invulnerable && time.Now().UnixMilli()%2 == 0 {
		op := &colorm.DrawImageOptions{}
		op.GeoM.Translate(an.container.position.x, an.container.position.y)
		var cm colorm.ColorM
		cm.Scale(0, 0, 0, 1)
		cm.Translate(0xff/0xff, 0xff/0xff, 0x00/0xff, 0)
		colorm.DrawImage(screen, getSprite(frameIndex).(*ebiten.Image), cm, op)
	}

	if debug {
		for _, b := range an.container.hitBoxes {
			vector.DrawFilledRect(screen, float32(an.container.position.x+b.x), float32(an.container.position.y+b.y), float32(b.w), float32(b.h), color.RGBA{0xff, 0, 0, 0xff}, false)
		}
	}

}

// change animation sequence is exists
func (an *animator) setSequence(name string) {
	if _, ok := an.sequences[name]; ok {
		an.currentSeq = name
		an.lastFrameChange = time.Now()
	}

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
