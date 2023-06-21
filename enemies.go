package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

func cleanEnemyList(g *Game) {
	for i := 0; i < len(g.enemies); i++ {
		DebugPrintf("cleanEnemyList", len(g.enemies))
		if !g.enemies[i].active {
			g.enemies[i] = g.enemies[len(g.enemies)-1] // Copy last element to index i.
			g.enemies[len(g.enemies)-1] = nil          // Erase last element (write zero value).
			g.enemies = g.enemies[:len(g.enemies)-1]   // Truncate slice.
			DebugPrintf("cleanEnemyList", len(g.enemies))
		}
	}
}

func spawnBaloon(g *Game, x, y float64, speed Vector) {
	enemy := newEntity(
		"Baloon",
		Vector{x: x, y: y},
	)
	enemy.entityType = typeEnemy
	enemy.scoreValue = 20
	enemy.hitBoxes = append(enemy.hitBoxes, Box{5, 2, 15, 20})

	sequences := map[string]*sequence{
		"idle":    newSequence(animEnemyBaloon, animEnemyBaloonFPS, true),
		"destroy": newSequence(animEnemyBaloonDie, ANIM_FPS, false),
	}
	animator := newAnimator(enemy, sequences, "idle")
	enemy.addComponent(animator)

	sounds := map[Sound]*audio.Player{SoundDestroy: sfx_exp_odd1Player, SoundFire: sfx_wpn_laser8Player}
	soundPlayer := newSoundPlayer(enemy, sounds)
	enemy.addComponent(soundPlayer)

	cmover := NewConstantMover(enemy, Vector{x: 0.2, y: 1})
	enemy.addComponent(cmover)
	g.enemies = append(g.enemies, enemy)
}

func spawnThing(g *Game, x, y float64, speed Vector) {
	enemy := newEntity(
		"Thing",
		Vector{x: x, y: y},
	)
	enemy.entityType = typeEnemy
	enemy.scoreValue = 50
	enemy.hitBoxes = append(enemy.hitBoxes, Box{5, 2, 15, 20})

	sequences := map[string]*sequence{
		"idle":    newSequence(animEnemyThing, ANIM_FPS, true),
		"destroy": newSequence(animExplosion, ANIM_FPS, false),
	}
	animator := newAnimator(enemy, sequences, "idle")
	enemy.addComponent(animator)

	sounds := map[Sound]*audio.Player{SoundDestroy: sfx_exp_odd1Player, SoundFire: sfx_wpn_laser8Player}
	soundPlayer := newSoundPlayer(enemy, sounds)
	enemy.addComponent(soundPlayer)

	cmover := NewConstantMover(enemy, Vector{x: 0.2, y: 1})
	enemy.addComponent(cmover)
	g.enemies = append(g.enemies, enemy)
}

func spawnFlyingMan1(g *Game, x, y float64, speed Vector) {
	enemy := newEntity(
		"FlyingMan1",
		Vector{x: x, y: y},
	)
	enemy.entityType = typeEnemy
	enemy.scoreValue = 100
	enemy.hitBoxes = append(enemy.hitBoxes, Box{5, 2, 15, 20})

	sequences := map[string]*sequence{
		"idle":    newSequence(animEnemyFlyingMan1, ANIM_FPS, true),
		"destroy": newSequence(animSuperFlyingManDie, ANIM_FPS, false),
	}
	animator := newAnimator(enemy, sequences, "idle")
	enemy.addComponent(animator)

	sounds := map[Sound]*audio.Player{SoundDestroy: sfx_exp_odd1Player, SoundFire: sfx_wpn_laser8Player}
	soundPlayer := newSoundPlayer(enemy, sounds)
	enemy.addComponent(soundPlayer)

	cmover := NewConstantMover(enemy, speed)
	enemy.addComponent(cmover)

	cshooter := NewAimedShooter(
		enemy,
		time.Millisecond*900,
		g.enemiesBullettPool,
		g.playerOne,
		g.playerTwo,
	)
	enemy.addComponent(cshooter)

	g.enemies = append(g.enemies, enemy)
}

func spawnCat(g *Game, x, y float64, speed Vector) {
	enemy := newEntity(
		"Cat",
		Vector{x: x, y: y},
	)
	enemy.entityType = typeEnemy
	enemy.scoreValue = 200
	enemy.hitBoxes = append(enemy.hitBoxes, Box{5, 2, 15, 20})

	sequences := map[string]*sequence{
		"idle":    newSequence(animEnemyCat, ANIM_FPS, true),
		"destroy": newSequence(animExplosion, ANIM_FPS, false),
	}
	animator := newAnimator(enemy, sequences, "idle")
	enemy.addComponent(animator)

	sounds := map[Sound]*audio.Player{SoundDestroy: sfx_exp_odd1Player, SoundFire: sfx_wpn_laser8Player}
	soundPlayer := newSoundPlayer(enemy, sounds)
	enemy.addComponent(soundPlayer)

	cmover := NewConstantMover(enemy, speed)
	enemy.addComponent(cmover)

	cshooter := NewConstantShooter(
		enemy,
		time.Millisecond*600,
		g.enemiesBullettPool,
	)
	enemy.addComponent(cshooter)

	g.enemies = append(g.enemies, enemy)
}

func spawnEnemies(g *Game) {

	x := rand.Float64() * float64(screenWidth)
	speed := Vector{rand.Float64()*2.0 - 1.0, rand.Float64()}

	// Enemies
	if g.position%96 == 0 {
		spawnBaloon(g, x, -spriteSize, speed)
	}
	if g.position%108 == 0 {
		spawnThing(g, x, -spriteSize, speed)

	}
	if g.position%141 == 0 {
		// Enemies
		spawnFlyingMan1(g, x, -spriteSize, speed)

	}
	if g.position%303 == 0 {
		spawnCat(g, x, -spriteSize, speed)
	}
}
