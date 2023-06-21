package main

import (
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type TapeCommand struct {
	time      time.Duration
	command   string
	arguments string
}

type ScriptedMover struct {
	active     bool
	container  *Entity
	speed      Vector
	lastUpdate time.Time
	head       int
	script     []TapeCommand
}

func NewScriptedMover(container *Entity, script []TapeCommand) *ScriptedMover {
	return &ScriptedMover{
		active:    true,
		container: container,
		speed:     Vector{0, 0},
		head:      0,
		script:    script,
	}
}

func (mover *ScriptedMover) Update() {
	if !mover.container.active || mover.container.hit {
		return
	}

	// play script
	c := mover.script[mover.head]
	if time.Since(mover.lastUpdate) > c.time {
		DebugPrintf(mover.container.name)
		DebugPrintf("Command:", c)

		switch c.command {
		case "speed":
			splitted := strings.Fields(c.arguments)
			x, _ := strconv.ParseFloat(splitted[0], 64)
			y, _ := strconv.ParseFloat(splitted[1], 64)
			mover.speed = Vector{x, y}
		case "rotate":
			rot, _ := strconv.ParseFloat(c.arguments, 64)
			mover.container.rotation = rot
		case "rotateAdd":
			rot, _ := strconv.ParseFloat(c.arguments, 64)
			mover.container.rotation += rot
		case "rewind":
			mover.head = -1
		case "wait":
		}
		mover.lastUpdate = time.Now()
		if mover.head < len(mover.script)-1 {
			mover.head++
		}

	}

	// move using rotation
	rotationRad := mover.container.rotation * (math.Pi / 180.0)
	mover.container.position.x += mover.speed.x * math.Cos(rotationRad)
	mover.container.position.y += mover.speed.y * math.Sin(rotationRad)

	// entity out of screen
	if mover.container.position.x > screenWidth+spriteSize || mover.container.position.x+spriteSize < 0 ||
		mover.container.position.y > screenHeight+spriteSize || mover.container.position.y+spriteSize < 0 {
		mover.container.active = false
	}
}

func (k *ScriptedMover) Draw(screen *ebiten.Image) {
	return
}
