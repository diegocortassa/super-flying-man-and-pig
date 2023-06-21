package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type FollowMover struct {
	active        bool
	container     *Entity
	lead          *Entity
	delay         time.Duration // delay in milliseconds
	speed         Vector
	lastUpdate    time.Time
	lastLeadSpeed Vector
}

func NewFollowMover(container *Entity, lead *Entity) *FollowMover {
	leadMover := lead.getComponent(&ConstantMover{}).(*ConstantMover)
	leadSpeed := leadMover.speed
	return &FollowMover{
		active:        true,
		container:     container,
		lead:          lead,
		lastUpdate:    time.Now(),
		lastLeadSpeed: leadSpeed,
	}
}

func (mover *FollowMover) Update() {
	if !mover.active || mover.container.hit {
		return
	}

	if time.Since(mover.lastUpdate) > time.Millisecond*500 && mover.lead.active {
		leadMover := mover.lead.getComponent(&ConstantMover{}).(*ConstantMover)
		mover.speed = leadMover.speed
	}
	mover.container.position.x += mover.speed.x
	mover.container.position.y += mover.speed.y

	// entity out fo screen
	if mover.container.position.x > screenWidth+spriteSize || mover.container.position.x+spriteSize < 0 ||
		mover.container.position.y > screenHeight+spriteSize || mover.container.position.y+spriteSize < 0 {
		mover.container.active = false
	}
}

func (k *FollowMover) Draw(screen *ebiten.Image) {
	return
}
