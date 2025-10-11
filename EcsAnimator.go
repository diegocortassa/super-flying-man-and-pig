package main

import (
	"image/color"
	"time"

	"github.com/diegocortassa/super-flying-man-and-pig/assets"
	"github.com/diegocortassa/super-flying-man-and-pig/globals"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type animator struct {
	container       *Entity
	sequences       map[string]*Sequence
	currentSeq      string
	lastFrameChange time.Time
	finished        bool
	// mu              sync.Mutex
}

func NewAnimator(container *Entity, sequences map[string]*Sequence, defaultSequence string) *animator {
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
	if an.container.Hit && an.currentSeq == "idle" {
		// in no sequence destroy exists enemy just disappear
		if _, ok := an.sequences["destroy"]; ok {
			an.sequences["destroy"].currentFrame = 0
			an.currentSeq = "destroy"
		} else {
			an.container.Active = false
		}
		sp := an.container.GetComponent(&SoundPlayer{}).(*SoundPlayer)
		if sp != nil {
			sp.PlaySound(SoundDestroy)

		}
	}

	// TODO Collision logic should not be here !!
	if an.container.Hit && an.finished {
		an.container.Invulnerable = true
		an.container.InvulnerableSetTime = time.Now()
		an.container.Hit = false
		an.currentSeq = "idle"
		if an.container.Lives < 1 {
			an.container.Active = false
		}
	}
	if an.container.Invulnerable && time.Since(an.container.InvulnerableSetTime) > time.Second*3 {
		an.container.Invulnerable = false
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
	opts.GeoM.Translate(an.container.Position.X, an.container.Position.Y)
	screen.DrawImage(assets.GetSprite(frameIndex).(*ebiten.Image), opts)

	if an.container.Invulnerable && time.Now().UnixMilli()%2 == 0 {
		op := &colorm.DrawImageOptions{}
		op.GeoM.Translate(an.container.Position.X, an.container.Position.Y)
		var cm colorm.ColorM
		cm.Scale(0, 0, 0, 1)
		cm.Translate(0xff/0xff, 0xff/0xff, 0x00/0xff, 0)
		colorm.DrawImage(screen, assets.GetSprite(frameIndex).(*ebiten.Image), cm, op)
	}

	if globals.Debug {
		for _, b := range an.container.HitBoxes {
			vector.StrokeRect(screen, float32(an.container.Position.X+b.X), float32(an.container.Position.Y+b.Y), float32(b.W), float32(b.H), 1, color.RGBA{0xff, 0, 0, 0xff}, false)
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
