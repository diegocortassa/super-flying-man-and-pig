package main

import (
	"math/rand"
	"time"

	"github.com/dcortassa/superflyingmanandpig/assets"
	"github.com/dcortassa/superflyingmanandpig/debug"
	"github.com/dcortassa/superflyingmanandpig/globals"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

var spawnScript = []TapeCommand{
	{time.Millisecond * 1, "wait", ""},
	{time.Millisecond * 500, "Baloon", "Baloon"},
	{time.Millisecond * 500, "Baloon", "Baloon"},
	{time.Millisecond * 500, "Baloon", "Baloon"},
	{time.Millisecond * 500, "Baloon", "Baloon"},
	{time.Millisecond * 2000, "Thing", ""},
	{time.Millisecond * 1000, "Thing", ""},
	{time.Millisecond * 1000, "Thing", ""},
	{time.Millisecond * 1000, "Thing", ""},
	{time.Millisecond * 4000, "FlyingMan1", ""},
	{time.Millisecond * 300, "FlyingMan1", ""},
	{time.Millisecond * 300, "FlyingMan1", ""},
	{time.Millisecond * 300, "FlyingMan1", ""},
	{time.Millisecond * 300, "FlyingMan1", ""},
	{time.Millisecond * 300, "FlyingMan1", ""},
	{time.Millisecond * 5000, "Cat", ""},
	{time.Millisecond * 1000, "Cat", ""},
	{time.Millisecond * 1000, "Cat", ""},
	{time.Millisecond * 1000, "Cat", ""},
	{time.Millisecond * 1000, "Cat", ""},
	{time.Millisecond * 3000, "Baloon", "Baloon"},
	{time.Millisecond * 500, "Baloon", "Baloon"},
	{time.Millisecond * 500, "Baloon", "Baloon"},
	{time.Millisecond * 500, "Baloon", "Baloon"},
	{time.Millisecond * 500, "Baloon", "Baloon"},
	{time.Millisecond * 4000, "FlyingMan2", "FlyingMan2"},
	{time.Millisecond * 300, "FlyingMan2", "FlyingMan2"},
	{time.Millisecond * 300, "FlyingMan2", "FlyingMan2"},
	{time.Millisecond * 300, "FlyingMan2", "FlyingMan2"},
	{time.Millisecond * 300, "FlyingMan2", "FlyingMan2"},
	{time.Millisecond * 300, "FlyingMan2", "FlyingMan2"},
	{time.Millisecond * 5000, "rewind", ""},
}

var enemyRotateAndGo = []TapeCommand{
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

func spawnFlyingMan1(g *Game, x, y float64, speed Vector) {
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

func spawnFlyingMan2(g *Game, x, y float64, speed Vector) {
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

func SpawnEnemies(g *Game) {

	speed := Vector{X: rand.Float64()*2.0 - 1.0, Y: rand.Float64()}
	min := 50.0
	max := globals.ScreenWidth - 50.0
	x := rand.Float64()*(max-min) + min

	c := spawnScript[g.spawnHead]
	debug.DebugPrintf("time.Since(lastSpawn):", time.Since(g.lastSpawn), "spawnHead:", g.spawnHead)
	if time.Since(g.lastSpawn) > c.Time {
		debug.DebugPrintf("Spawn Command:", c)
		switch c.Command {
		case "Baloon":
			spawnBaloon(g, globals.ScreenWidth/2, -assets.SpriteSize, speed)
		case "Thing":
			spawnThing(g, x, -assets.SpriteSize, speed)
		case "FlyingMan1":
			spawnFlyingMan1(g, x, -assets.SpriteSize, speed)
		case "FlyingMan2":
			spawnFlyingMan2(g, x, -assets.SpriteSize, speed)
		case "Cat":
			spawnCat(g, x, -assets.SpriteSize, speed)
		case "wait":
		case "rewind":
			g.spawnHead = -1
		}

		g.lastSpawn = time.Now()
		if g.spawnHead < len(spawnScript)-1 {
			g.spawnHead++
		}
	}

	// Old pseudo random spawning
	// if g.Position%96 == 0 {
	// 	// spawnBaloon(g, x, -assets.SpriteSize, speed)
	// 	spawnBaloon(g, screenWidth/2, -assets.SpriteSize, speed)
	// }
	// if g.Position%108 == 0 {
	// 	spawnThing(g, x, -assets.SpriteSize, speed)

	// }
	// if g.Position%141 == 0 {
	// 	// Enemies
	// 	spawnFlyingMan1(g, x, -assets.SpriteSize, speed)

	// }
	// if g.Position%303 == 0 {
	// 	spawnCat(g, x, -assets.SpriteSize, speed)
	// }
}
