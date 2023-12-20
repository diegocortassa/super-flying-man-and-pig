package main

import (
	"math/rand"
	"time"

	"github.com/dcortassa/superflyingmanandpig/assets"
	"github.com/dcortassa/superflyingmanandpig/debug"
	"github.com/dcortassa/superflyingmanandpig/globals"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type SpawnCommand struct {
	Position int
	Command  string
	X        float64
}

func SpawnEnemies(g *Game) {

	speed := Vector{X: rand.Float64()*2.0 - 1.0, Y: rand.Float64()}
	min := 50.0
	max := globals.ScreenWidth - 50.0

	c := spawnScript[g.spawnHead]
	if c.X == -1 {
		c.X = rand.Float64()*(max-min) + min
	} else {
		c.X = c.X - assets.SpriteSize/2 // position by sprite center
	}

	// If the game position is ahead of spawnHead skip passed positions
	if g.Position > c.Position-globals.ScreenHeight-assets.SpriteSize/2 && g.spawnHead < len(spawnScript)-1 {
		g.spawnHead++
	}

	if g.Position == c.Position-globals.ScreenHeight-assets.SpriteSize/2 { // Position based spawn
		debug.DebugPrintf("*** Spawn Command:", c)
		switch c.Command {
		case "Baloon":
			spawnBaloon(g, globals.ScreenWidth/2, -assets.SpriteSize, speed)
		case "Thing":
			spawnThing(g, c.X, -assets.SpriteSize, speed)
		case "FlyingMan1":
			spawnFlyingMan1(g, c.X, -assets.SpriteSize, speed)
		case "FlyingMan2":
			spawnFlyingMan2(g, c.X, -assets.SpriteSize, speed)
		case "Cat":
			spawnCat(g, c.X, -assets.SpriteSize, speed)
		case "Vulcano":
			spawnVulcano(g, c.X, -assets.SpriteSize, Vector{X: 0.0, Y: 0.5})
		}

		if g.spawnHead < len(spawnScript)-1 {
			g.spawnHead++
		}
	}
}

func CleanEnemyList(g *Game) {
	for i := 0; i < len(g.enemies); i++ {
		debug.DebugPrintf("CleanEnemyList", len(g.enemies))
		if !g.enemies[i].Active {
			g.enemies[i] = g.enemies[len(g.enemies)-1] // Copy last element to index i.
			g.enemies[len(g.enemies)-1] = nil          // Erase last element (write zero value).
			g.enemies = g.enemies[:len(g.enemies)-1]   // Truncate slice.
			debug.DebugPrintf("CleanEnemyList", len(g.enemies))
		}
	}
}

// Specific enemy spawn functions

// Baloon
func spawnBaloon(g *Game, x, y float64, speed Vector) {
	enemy := NewEntity(
		"Baloon",
		Vector{X: x, Y: y},
	)
	enemy.EntityType = TypeEnemy
	enemy.ScoreValue = 20
	enemy.HitBoxes = append(enemy.HitBoxes, Box{X: 5, Y: 2, W: 15, H: 20})

	sequences := map[string]*Sequence{
		"idle":    NewSequence(assets.AnimEnemyBaloon, assets.AnimEnemyBaloonFPS, true),
		"destroy": NewSequence(assets.AnimEnemyBaloonDie, assets.AnimFPS, false),
	}
	animator := NewAnimator(enemy, sequences, "idle")
	enemy.AddComponent(animator)

	sounds := map[Sound]*audio.Player{SoundDestroy: assets.Sfx_exp_odd1Player, SoundFire: assets.Sfx_wpn_laser8Player}
	soundPlayer := NewSoundPlayer(enemy, sounds)
	enemy.AddComponent(soundPlayer)

	cmover := NewConstantMover(enemy, Vector{X: 0.2, Y: 1})
	enemy.AddComponent(cmover)
	g.enemies = append(g.enemies, enemy)
}

// Thing
func spawnThing(g *Game, x, y float64, speed Vector) {
	enemy := NewEntity(
		"Thing",
		Vector{X: x, Y: y},
	)
	enemy.EntityType = TypeEnemy
	enemy.ScoreValue = 50
	enemy.HitBoxes = append(enemy.HitBoxes, Box{X: 5, Y: 2, W: 15, H: 20})

	sequences := map[string]*Sequence{
		"idle":    NewSequence(assets.AnimEnemyThing, assets.AnimFPS, true),
		"destroy": NewSequence(assets.AnimExplosion, assets.AnimFPS, false),
	}
	animator := NewAnimator(enemy, sequences, "idle")
	enemy.AddComponent(animator)

	sounds := map[Sound]*audio.Player{SoundDestroy: assets.Sfx_exp_odd1Player, SoundFire: assets.Sfx_wpn_laser8Player}
	soundPlayer := NewSoundPlayer(enemy, sounds)
	enemy.AddComponent(soundPlayer)

	cmover := NewConstantMover(enemy, Vector{X: 0.2, Y: 1})
	enemy.AddComponent(cmover)
	g.enemies = append(g.enemies, enemy)
}

// Flying man 1
func spawnFlyingMan1(g *Game, x, y float64, speed Vector) {

	var enemyRotateAndGo = []TimedCommand{
		//     270
		// 180     0
		//     90
		{time.Millisecond * 1, "rotate", "90"},
		{time.Millisecond * 1000, "speed", "2 2"},
		{time.Millisecond * 1500, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "wait", ""},
	}

	enemy := NewEntity(
		"FlyingMan1",
		Vector{X: x, Y: y},
	)
	enemy.EntityType = TypeEnemy
	enemy.ScoreValue = 100
	enemy.HitBoxes = append(enemy.HitBoxes, Box{X: 5, Y: 2, W: 15, H: 20})

	sequences := map[string]*Sequence{
		"idle":    NewSequence(assets.AnimEnemyFlyingMan1, assets.AnimFPS, true),
		"destroy": NewSequence(assets.AnimSuperFlyingManDie, assets.AnimFPS, false),
	}
	animator := NewAnimator(enemy, sequences, "idle")
	enemy.AddComponent(animator)

	sounds := map[Sound]*audio.Player{SoundDestroy: assets.Sfx_exp_odd1Player, SoundFire: assets.Sfx_wpn_laser8Player}
	soundPlayer := NewSoundPlayer(enemy, sounds)
	enemy.AddComponent(soundPlayer)

	// cmover := NewConstantMover(enemy, speed)
	cmover := NewScriptedMover(enemy, enemyRotateAndGo)
	enemy.AddComponent(cmover)

	cshooter := NewAimedShooter(
		enemy,
		time.Millisecond*900,
		g.enemiesBullettPool,
		g.playerOne,
		g.playerTwo,
	)
	enemy.AddComponent(cshooter)

	g.enemies = append(g.enemies, enemy)
}

// Flying man 2
func spawnFlyingMan2(g *Game, x, y float64, speed Vector) {

	var enemyRotateAndGo = []TimedCommand{
		//     270
		// 180     0
		//     90
		{time.Millisecond * 1, "rotate", "90"},
		{time.Millisecond * 1000, "speed", "2 2"},
		{time.Millisecond * 1500, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "rotateAdd", "22"},
		{time.Millisecond * 200, "wait", ""},
	}

	enemy := NewEntity(
		"FlyingMan2",
		Vector{X: globals.ScreenWidth / 2, Y: -assets.SpriteSize},
	)
	enemy.EntityType = TypeEnemy
	enemy.ScoreValue = 100
	enemy.HitBoxes = append(enemy.HitBoxes, Box{X: 5, Y: 2, W: 15, H: 20})

	sequences := map[string]*Sequence{
		"idle":    NewSequence(assets.AnimEnemyFlyingMan2, assets.AnimFPS, true),
		"destroy": NewSequence(assets.AnimSuperFlyingManDie, assets.AnimFPS, false),
	}
	animator := NewAnimator(enemy, sequences, "idle")
	enemy.AddComponent(animator)

	sounds := map[Sound]*audio.Player{SoundDestroy: assets.Sfx_exp_odd1Player, SoundFire: assets.Sfx_wpn_laser8Player}
	soundPlayer := NewSoundPlayer(enemy, sounds)
	enemy.AddComponent(soundPlayer)

	// cmover := NewConstantMover(enemy, speed)
	cmover := NewScriptedMover(enemy, enemyRotateAndGo)
	enemy.AddComponent(cmover)

	cshooter := NewAimedShooter(
		enemy,
		time.Millisecond*900,
		g.enemiesBullettPool,
		g.playerOne,
		g.playerTwo,
	)
	enemy.AddComponent(cshooter)

	g.enemies = append(g.enemies, enemy)
}

// Cat
func spawnCat(g *Game, x, y float64, speed Vector) {
	enemy := NewEntity(
		"Cat",
		Vector{X: x, Y: y},
	)
	enemy.EntityType = TypeEnemy
	enemy.ScoreValue = 200
	enemy.HitBoxes = append(enemy.HitBoxes, Box{X: 5, Y: 2, W: 15, H: 20})

	sequences := map[string]*Sequence{
		"idle":    NewSequence(assets.AnimEnemyCat, assets.AnimFPS, true),
		"destroy": NewSequence(assets.AnimExplosion, assets.AnimFPS, false),
	}
	animator := NewAnimator(enemy, sequences, "idle")
	enemy.AddComponent(animator)

	sounds := map[Sound]*audio.Player{SoundDestroy: assets.Sfx_exp_odd1Player, SoundFire: assets.Sfx_wpn_laser8Player}
	soundPlayer := NewSoundPlayer(enemy, sounds)
	enemy.AddComponent(soundPlayer)

	cmover := NewConstantMover(enemy, speed)
	enemy.AddComponent(cmover)

	cshooter := NewConstantShooter(
		enemy,
		time.Millisecond*600,
		g.enemiesBullettPool,
	)
	enemy.AddComponent(cshooter)

	g.enemies = append(g.enemies, enemy)
}

// Vulcano
func spawnVulcano(g *Game, x float64, y float64, speed Vector) {
	enemy := NewEntity(
		"Vulcano",
		Vector{X: x, Y: y},
	)
	enemy.EntityType = TypeEnemy
	enemy.ScoreValue = 500
	// No hitbox, vulcanos are indestructible
	// enemy.HitBoxes = append(enemy.HitBoxes, Box{X: 10, Y: 10, W: 0, H: 0})

	sequences := map[string]*Sequence{
		"idle":    NewSequence(assets.AnimEnemyVulcano, assets.AnimFPS, true),
		"destroy": NewSequence(assets.AnimExplosion, assets.AnimFPS, false),
	}
	animator := NewAnimator(enemy, sequences, "idle")
	enemy.AddComponent(animator)

	sounds := map[Sound]*audio.Player{SoundDestroy: assets.Sfx_exp_odd1Player, SoundFire: assets.Sfx_exp_short_hard2Player}
	soundPlayer := NewSoundPlayer(enemy, sounds)
	enemy.AddComponent(soundPlayer)

	cmover := NewConstantMover(enemy, speed)
	enemy.AddComponent(cmover)

	rshooter := NewRotativeShooter(
		enemy,
		time.Millisecond*300,
		180,
		25,
		g.enemiesBullettPool,
	)
	enemy.AddComponent(rshooter)

	g.enemies = append(g.enemies, enemy)
}
