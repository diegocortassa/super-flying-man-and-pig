package main

import (
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/dcortassa/superflyingmanandpig/assets"
	"github.com/dcortassa/superflyingmanandpig/debug"
	"github.com/dcortassa/superflyingmanandpig/globals"
	"github.com/hajimehoshi/ebiten/v2"
)

type TimedCommand struct {
	Time      time.Duration
	Command   string
	Arguments string
}

// Moves an entity Following a given script
type ScriptedMover struct {
	active     bool
	container  *Entity
	speed      Vector
	lastUpdate time.Time
	head       int
	script     []TimedCommand
}

func NewScriptedMover(container *Entity, script []TimedCommand) *ScriptedMover {
	return &ScriptedMover{
		active:    true,
		container: container,
		speed:     Vector{0, 0},
		head:      0,
		script:    script,
	}
}

func (mover *ScriptedMover) Update() {
	if !mover.container.Active || mover.container.Hit {
		return
	}

	// play script
	c := mover.script[mover.head]
	if time.Since(mover.lastUpdate) > c.Time {
		debug.DebugPrintf(mover.container.Name)
		debug.DebugPrintf("Command:", c)

		switch c.Command {
		case "speed":
			splitField := strings.Fields(c.Arguments)
			x, _ := strconv.ParseFloat(splitField[0], 64)
			y, _ := strconv.ParseFloat(splitField[1], 64)
			mover.speed = Vector{x, y}
		case "rotate":
			rot, _ := strconv.ParseFloat(c.Arguments, 64)
			mover.container.Rotation = rot
		case "rotateAdd":
			rot, _ := strconv.ParseFloat(c.Arguments, 64)
			mover.container.Rotation += rot
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
	rotationRad := mover.container.Rotation * (math.Pi / 180.0)
	mover.container.Position.X += mover.speed.X * math.Cos(rotationRad)
	mover.container.Position.Y += mover.speed.Y * math.Sin(rotationRad)

	// entity out of screen
	if mover.container.Position.X > globals.ScreenWidth || mover.container.Position.X+assets.SpriteSize < 0 ||
		mover.container.Position.Y > globals.ScreenHeight || mover.container.Position.Y+assets.SpriteSize < 0 {
		mover.container.Active = false
	}
}

func (k *ScriptedMover) Draw(screen *ebiten.Image) {
	return
}
