package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type Sound int

const (
	SoundFire Sound = iota
	SoundDestroy
	SoundIdle
	SoundGameOver
)

type SoundPlayer struct {
	container    *Entity
	sounds       map[Sound]*audio.Player
	currentSound Sound
}

func newSoundPlayer(container *Entity, sounds map[Sound]*audio.Player) *SoundPlayer {
	var sp SoundPlayer

	sp.container = container
	sp.sounds = sounds

	return &sp
}

func (sp *SoundPlayer) Update() {
	// no update necessary for SoundPlayer
}

func (an *SoundPlayer) Draw(screen *ebiten.Image) {
	// no draw necessary for SoundPlayer
}

func (sp *SoundPlayer) PlaySound(s Sound) {
	if ss, ok := sp.sounds[s]; ok {
		ss.Rewind()
		ss.Play()
	}
}
