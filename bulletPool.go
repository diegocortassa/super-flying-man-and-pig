package main

import (
	_ "embed"
	"fmt"
)

func initBulletPool(name string, etype entityType, anim []int, bulletPoolSize int, speed Vector, hitbox Box) []*Entity {
	pool := make([]*Entity, 0)
	for i := 0; i < bulletPoolSize; i++ {
		bul := newEntity(name+fmt.Sprint(i), Vector{x: 0, y: 0})
		bul.hitBoxes = append(bul.hitBoxes, hitbox)
		bul.entityType = etype

		sequences := map[string]*sequence{
			"idle": newSequence(anim, ANIM_FPS, true),
		}
		animator := newAnimator(bul, sequences, "idle")
		bul.addComponent(animator)

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
