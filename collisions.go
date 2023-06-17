package main

func isCollisionSquare(c1, c2 Box) bool {
	return c1.x < c2.x+c2.w &&
		c1.x+c1.w > c2.x &&
		c1.y < c2.y+c2.h &&
		c1.y+c1.h > c2.y
}

// Check if any collision for all hitboxes in entity
func CollideBoxes(c1, c2 *Entity) bool {
	for _, b1 := range c1.hitBoxes {
		for _, b2 := range c2.hitBoxes {

			// range gives copies of item we con modify them in place
			// to get world position without worries to modify entity
			b1.x += c1.position.x
			b1.y += c1.position.y
			b2.x += c2.position.x
			b2.y += c2.position.y

			if isCollisionSquare(b1, b2) {
				return true
			}
		}
	}
	return false
}

func Collide(l1, l2 []*Entity, g *Game) {
	for i := 0; i < len(l1); i++ {
		for j := 0; j < len(l2); j++ {
			if !l1[i].active || !l2[j].active {
				continue
			}
			DebugPrintf("Collision Check:", l1[i].name, i, "<>", l2[j].name, j)
			if CollideBoxes(l1[i], l2[j]) {
				DebugPrintf("Collision:", l1[i].name, "<>", l2[j].name)
				HandleCollision(l1[i], l2[j], g)
			}
		}
	}
}

func CheckCollisions(g *Game) {
	Collide(g.playerOneBullettPool, g.enemies, g)
	Collide(g.playerTwoBullettPool, g.enemies, g)
	Collide(g.enemiesBullettPool, []*Entity{g.playerOne}, g)
	Collide(g.enemiesBullettPool, []*Entity{g.playerTwo}, g)
	Collide(g.enemies, []*Entity{g.playerOne}, g)
	Collide(g.enemies, []*Entity{g.playerTwo}, g)

}

func HandleCollision(e1, e2 *Entity, g *Game) {
	if e1.entityType == typePlayerOneBullet {
		e1.active = false
		e2.active = false
		g.playerOne.scores += e2.scoreValue
		if g.playerOne.scores > g.hiScores {
			g.hiScores = g.playerOne.scores
		}
	}
	if e1.entityType == typePlayerTwoBullet {
		e1.active = false
		e2.active = false
		g.playerTwo.scores += e2.scoreValue
		if g.playerTwo.scores > g.hiScores {
			g.hiScores = g.playerTwo.scores
		}
	}
	if e1.entityType == typeEnemyBullet || e1.entityType == typeEnemy {
		e1.active = false
		e2.lives -= 1
		if e2.lives < 1 {
			e2.active = false
		}
	}
}
