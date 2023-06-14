package main

import (
	_ "embed"
)

const (
	bulletSpeed = -4
	// playerOneBulletAnimSpeed = 10.0
	bulletPoolSize = 5
)

func initBulletPool(anim []int) []*Entity {
	pool := make([]*Entity, 0)
	for i := 0; i < bulletPoolSize; i++ {
		bul := newEntity(SpriteSheetImage, anim, Vector{x: 0, y: 0})
		mover := NewConstantMover(bul, Vector{0, bulletSpeed})
		bul.addComponent(mover)
		bul.active = false
		pool = append(pool, bul)
	}
	return pool
}

func BullettFromPool(pool []*Entity) (*Entity, bool) {
	for _, bul := range pool {
		if !bul.active {
			return bul, true
		}
	}
	return nil, false
}
