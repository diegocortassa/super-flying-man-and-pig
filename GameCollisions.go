package main

import (
	"github.com/dcortassa/super-flying-man-and-pig/debug"
)

// TODO move collision detection to a component ***********

func CheckCollisions(g *Game) {
	Collide(g.playerOneBullettPool, g.enemies, g)
	Collide(g.playerTwoBullettPool, g.enemies, g)
	Collide(g.enemiesBullettPool, []*Entity{g.playerOne}, g)
	Collide(g.enemiesBullettPool, []*Entity{g.playerTwo}, g)
	Collide(g.enemies, []*Entity{g.playerOne}, g)
	Collide(g.enemies, []*Entity{g.playerTwo}, g)

}

func Collide(l1, l2 []*Entity, g *Game) {
	for i := 0; i < len(l1); i++ {
		for j := 0; j < len(l2); j++ {
			if !l1[i].Active || !l2[j].Active {
				continue
			}
			debug.DebugPrintf("Collision Check:", l1[i].Name, i, "<>", l2[j].Name, j)
			if CollideBoxes(l1[i], l2[j]) {
				debug.DebugPrintf("Collision:", l1[i].Name, "<>", l2[j].Name)
				HandleCollision(l1[i], l2[j], g)
			}
		}
	}
}

// Check if any collision for all hitboxes in entity
func CollideBoxes(c1, c2 *Entity) bool {
	for _, b1 := range c1.HitBoxes {
		for _, b2 := range c2.HitBoxes {

			// range gives copies of item we con modify them in place
			// to get world position without worring to modify entity
			b1.X += c1.Position.X
			b1.Y += c1.Position.Y
			b2.X += c2.Position.X
			b2.Y += c2.Position.Y

			if isCollisionSquare(b1, b2) {
				return true
			}
		}
	}
	return false
}

// check is two squares collide
func isCollisionSquare(c1, c2 Box) bool {
	return c1.X < c2.X+c2.W &&
		c1.X+c1.W > c2.X &&
		c1.Y < c2.Y+c2.H &&
		c1.Y+c1.H > c2.Y
}

// we got a collision, whe handle it
func HandleCollision(e1, e2 *Entity, g *Game) {
	// playerOne bullet hit an enemy
	if e1.EntityType == TypePlayerOneBullet {
		if !e2.Exploding {
			e1.Active = false                   // bullet destroyed
			g.playerOne.Scores += e2.ScoreValue // increment scores
			if g.playerOne.Scores > g.HiScores {
				g.HiScores = g.playerOne.Scores
			}
		}
		e2.Hit = true
		e2.Exploding = true
		return
	}
	// playerTwo bullet hit an enemy
	if e1.EntityType == TypePlayerTwoBullet {
		if !e2.Exploding {
			e1.Active = false                   // bullet destroyed
			g.playerTwo.Scores += e2.ScoreValue // increment scores
			if g.playerTwo.Scores > g.HiScores {
				g.HiScores = g.playerTwo.Scores
			}
		}
		e2.Hit = true
		e2.Exploding = true
		return
	}

	// player hit by a bullet or rammed by an enemy
	if e1.EntityType == TypeEnemyBullet || e1.EntityType == TypeEnemy {

		if e1.EntityType == TypeEnemyBullet {
			e1.Active = false // bullet destroyed no explosion anim
		} else if !e2.Exploding {
			e1.Hit = true // animator will take care of playing explosion setting active to false when anim finishes
		}

		// TODO hit logic is partially shared with animator because it takes care of playing the destroy animation
		// TODO once an enemy is hit should not collide but adding && !e1.hit here gives superpowers to player
		if !e2.Invulnerable && !e2.Hit && !e1.Exploding {
			// animator will reset hit and invulnerable flags
			e2.Hit = true
			// animator will set active to false if no more lives
			e2.Lives -= 1
			return
		}
	}
}
