package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/dcortassa/super-flying-man-and-pig/assets"
	"github.com/dcortassa/super-flying-man-and-pig/debug"
	"github.com/dcortassa/super-flying-man-and-pig/globals"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type SpawnCommand struct {
	Position int
	X        float64
	Command  string
}

func GetSpawnCommand(spawnHead int) SpawnCommand {
	min := 50.0
	max := globals.ScreenWidth - 50.0

	c := spawnScript[spawnHead]
	if c.X == -1 {
		c.X = rand.Float64() * (max - min)
	} else {
		c.X = c.X - assets.SpriteSize/2 // position by sprite center
	}

	return c
}

func SpawnEnemies(g *Game) {

	screenTopPosition := g.Position + globals.ScreenHeight + assets.SpriteSize/2

	// if X position is -1 we use a random X start position
	c := GetSpawnCommand(g.spawnHead)

	log.Println("*** top position", screenTopPosition, "Next at", c.Position)
	// If the game position is ahead of spawnHead skip passed positions
	for screenTopPosition > c.Position && g.spawnHead < len(spawnScript)-1 {
		log.Println("*** Skipping Command:", c, spawnScript[g.spawnHead], g.spawnHead)
		g.spawnHead++
		c = GetSpawnCommand(g.spawnHead)
	}

	for screenTopPosition == c.Position { // Position based spawn
		debug.DebugPrintf("*** Spawn Command:", c)
		log.Println("*** Spawn Command:", c, spawnScript[g.spawnHead], g.spawnHead)
		switch c.Command {
		case "BaloonA":
			spawnBaloonA(g, c.X, -assets.SpriteSize)
		case "BaloonB":
			spawnBaloonB(g, c.X, -assets.SpriteSize)
		case "Thing":
			spawnThing(g, c.X, -assets.SpriteSize)
		case "FlyingMan1":
			spawnFlyingMan1(g, c.X, -assets.SpriteSize)
		case "FlyingMan2A":
			spawnFlyingMan2A(g, c.X, -assets.SpriteSize)
		case "FlyingMan2B":
			spawnFlyingMan2B(g, c.X, -assets.SpriteSize)
		case "Cat":
			spawnCat(g, c.X, -assets.SpriteSize)
		case "Volcano":
			spawnVolcano(g, c.X, -assets.SpriteSize)
		case "Monkey":
			spawnMonkey(g, c.X, -assets.SpriteSize)
		}

		if g.spawnHead < len(spawnScript)-1 {
			g.spawnHead++
			c = GetSpawnCommand(g.spawnHead)
		} else {
			break
		}
	}

	// game looped, reset spawnHead
	if g.Position == 0 {
		g.spawnHead = 0
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

// BaloonA
func spawnBaloonA(g *Game, x, y float64) {
	enemy := NewEntity(
		"BaloonA",
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

	speed := Vector{X: -0.2, Y: 1}
	cmover := NewMoverConstant(enemy, speed)
	enemy.AddComponent(cmover)
	g.enemies = append(g.enemies, enemy)
}

// BaloonB
func spawnBaloonB(g *Game, x, y float64) {
	enemy := NewEntity(
		"BaloonB",
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

	speed := Vector{X: 0.2, Y: 1}
	mover := NewMoverConstant(enemy, speed)
	enemy.AddComponent(mover)
	g.enemies = append(g.enemies, enemy)
}

// Thing
func spawnThing(g *Game, x, y float64) {
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

	speed := Vector{X: 0.2, Y: 1}
	cmover := NewMoverConstant(enemy, speed)
	enemy.AddComponent(cmover)
	g.enemies = append(g.enemies, enemy)
}

// Flying man 1
func spawnFlyingMan1(g *Game, x, y float64) {

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

	mover := NewMoverScripted(enemy, enemyRotateAndGo)
	enemy.AddComponent(mover)

	shooter := NewShooterAimed(
		enemy,
		time.Millisecond*900,
		g.enemiesBullettPool,
		g.playerOne,
		g.playerTwo,
	)
	enemy.AddComponent(shooter)

	g.enemies = append(g.enemies, enemy)
}

// Flying man 2
func spawnFlyingMan2A(g *Game, x, y float64) {

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
		"FlyingMan2A",
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

	mover := NewMoverScripted(enemy, enemyRotateAndGo)
	enemy.AddComponent(mover)

	shooter := NewShooterAimed(
		enemy,
		time.Millisecond*900,
		g.enemiesBullettPool,
		g.playerOne,
		g.playerTwo,
	)
	enemy.AddComponent(shooter)

	g.enemies = append(g.enemies, enemy)
}

// Flying man 2
func spawnFlyingMan2B(g *Game, x, y float64) {

	var enemyRotateAndGo = []TimedCommand{
		//     270
		// 180     0
		//     90
		{time.Millisecond * 1, "rotate", "90"},
		{time.Millisecond * 1000, "speed", "2 2"},
		{time.Millisecond * 1500, "rotateAdd", "-22"},
		{time.Millisecond * 200, "rotateAdd", "-22"},
		{time.Millisecond * 200, "rotateAdd", "-22"},
		{time.Millisecond * 200, "rotateAdd", "-22"},
		{time.Millisecond * 200, "rotateAdd", "-22"},
		{time.Millisecond * 200, "rotateAdd", "-22"},
		{time.Millisecond * 200, "rotateAdd", "-22"},
		{time.Millisecond * 200, "rotateAdd", "-22"},
		{time.Millisecond * 200, "rotateAdd", "-22"},
		{time.Millisecond * 200, "rotateAdd", "-22"},
		{time.Millisecond * 200, "rotateAdd", "-22"},
		{time.Millisecond * 200, "rotateAdd", "-22"},
		{time.Millisecond * 200, "rotateAdd", "-22"},
		{time.Millisecond * 200, "rotateAdd", "-22"},
		{time.Millisecond * 200, "wait", ""},
	}

	enemy := NewEntity(
		"FlyingMan2B",
		Vector{X: globals.ScreenWidth/2 - assets.SpriteSize, Y: -assets.SpriteSize},
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

	mover := NewMoverScripted(enemy, enemyRotateAndGo)
	enemy.AddComponent(mover)

	shooter := NewShooterAimed(
		enemy,
		time.Millisecond*900,
		g.enemiesBullettPool,
		g.playerOne,
		g.playerTwo,
	)
	enemy.AddComponent(shooter)

	g.enemies = append(g.enemies, enemy)
}

// Cat
func spawnCat(g *Game, x, y float64) {
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

	// min = 0.1
	// max = 1.5
	// c.speed.X = rand.Float64() * (max - min)
	// c.speed.Y = rand.Float64() * (max - min)

	speed := Vector{X: rand.Float64()*2.0 - 1.0, Y: rand.Float64()}
	mover := NewMoverConstant(enemy, speed)
	enemy.AddComponent(mover)

	shooter := NewShooterConstant(
		enemy,
		time.Millisecond*600,
		g.enemiesBullettPool,
	)
	enemy.AddComponent(shooter)

	g.enemies = append(g.enemies, enemy)
}

// Volcano
func spawnVolcano(g *Game, x float64, y float64) {
	enemy := NewEntity(
		"Volcano",
		Vector{X: x, Y: y},
	)
	enemy.EntityType = TypeEnemy
	enemy.ScoreValue = 500
	// No hitbox, Volcanoes are indestructible
	// enemy.HitBoxes = append(enemy.HitBoxes, Box{X: 10, Y: 10, W: 0, H: 0})

	sequences := map[string]*Sequence{
		"idle":    NewSequence(assets.AnimEnemyVolcano, assets.AnimFPS, true),
		"destroy": NewSequence(assets.AnimExplosion, assets.AnimFPS, false),
	}
	animator := NewAnimator(enemy, sequences, "idle")
	enemy.AddComponent(animator)

	sounds := map[Sound]*audio.Player{SoundDestroy: assets.Sfx_exp_odd1Player, SoundFire: assets.Sfx_exp_short_hard2Player}
	soundPlayer := NewSoundPlayer(enemy, sounds)
	enemy.AddComponent(soundPlayer)

	speed := Vector{X: 0.0, Y: 0.498} // 4 va su, 6 va giu
	mover := NewMoverConstant(enemy, speed)
	enemy.AddComponent(mover)

	shooter := NewShooterRotative(
		enemy,
		time.Millisecond*300,
		180,
		25,
		g.enemiesBullettPool,
	)
	enemy.AddComponent(shooter)

	g.enemies = append(g.enemies, enemy)
}

// Monkey
func spawnMonkey(g *Game, x float64, y float64) {
	enemy := NewEntity(
		"Monkey",
		Vector{X: x, Y: y},
	)
	enemy.EntityType = TypeEnemy
	enemy.ScoreValue = 500
	// No hitbox, Volcanoes are indestructible
	// enemy.HitBoxes = append(enemy.HitBoxes, Box{X: 10, Y: 10, W: 0, H: 0})

	sequences := map[string]*Sequence{
		"idle":    NewSequence(assets.AnimEnemyMonkey, assets.AnimFPS, true),
		"destroy": NewSequence(assets.AnimExplosion, assets.AnimFPS, false),
	}
	animator := NewAnimator(enemy, sequences, "idle")
	enemy.AddComponent(animator)

	sounds := map[Sound]*audio.Player{SoundDestroy: assets.Sfx_exp_odd1Player, SoundFire: assets.Sfx_wpn_laser8Player}
	soundPlayer := NewSoundPlayer(enemy, sounds)
	enemy.AddComponent(soundPlayer)

	speed := Vector{X: 0.0, Y: 0.498} // 4 va su, 6 va giu
	mover := NewMoverConstant(enemy, speed)
	enemy.AddComponent(mover)

	shooter := NewShooterAimed(
		enemy,
		time.Millisecond*500,
		g.enemiesBullettPool,
		g.playerOne,
		g.playerTwo,
	)
	enemy.AddComponent(shooter)

	g.enemies = append(g.enemies, enemy)
}
