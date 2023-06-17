package main

import (
	_ "embed"
	"fmt"
)

func initBulletPool(name string, etype entityType, anim []int, bulletPoolSize int, speed Vector, hitbox Box) []*Entity {
	pool := make([]*Entity, 0)
	for i := 0; i < bulletPoolSize; i++ {
		bul := newEntity(SpriteSheetImage, anim, Vector{x: 0, y: 0})
		bul.hitBoxes = append(bul.hitBoxes, hitbox)
		bul.name = name + fmt.Sprint(i)
		bul.entityType = etype
		mover := NewConstantMover(bul, speed)
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
