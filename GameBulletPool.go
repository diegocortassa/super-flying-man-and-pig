package main

import (
	_ "embed"
	"fmt"
)

func initBulletPool(name string, etype EntityType, anim []int, animFPS float64, bulletPoolSize int, speed Vector, hitbox Box) []*Entity {
	pool := make([]*Entity, 0)
	for i := 0; i < bulletPoolSize; i++ {
		bul := NewEntity(name+fmt.Sprint(i), Vector{X: 0, Y: 0})
		bul.HitBoxes = append(bul.HitBoxes, hitbox)
		bul.EntityType = etype

		sequences := map[string]*Sequence{
			"idle": NewSequence(anim, animFPS, true),
		}
		animator := NewAnimator(bul, sequences, "idle")
		bul.AddComponent(animator)

		mover := NewConstantMover(bul, speed)
		bul.AddComponent(mover)
		bul.Active = false
		pool = append(pool, bul)
	}
	return pool
}

func BullettFromPool(pool []*Entity) (*Entity, bool) {
	for _, bul := range pool {
		if !bul.Active {
			return bul, true
		}
	}
	return nil, false
}
