package main

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

// Check if any collision for all hitboxes in entity
func CollideBoxes(c1, c2 *Entity) bool {
	for _, b1 := range c1.hitBoxes {
		for _, b2 := range c2.hitBoxes {

			// range gives copies of item we con modify them in place
			// to get world position without worring to modify entity
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

// check is two squares collide
func isCollisionSquare(c1, c2 Box) bool {
	return c1.x < c2.x+c2.w &&
		c1.x+c1.w > c2.x &&
		c1.y < c2.y+c2.h &&
		c1.y+c1.h > c2.y
}

// we got a collision, whe handle it
func HandleCollision(e1, e2 *Entity, g *Game) {
	// playerOne bullet hit an enemy
	if e1.entityType == typePlayerOneBullet {
		if !e2.exploding { // bullet pass trough enemy explosion
			e1.active = false // bullet destroyed
		}
		// e2.active = false
		e2.hit = true
		e2.exploding = true
		// e2an := e1.getComponent(&animator{}).(*animator)
		// e2an.lastFrameChange = time.Now()
		// e2an.setSequence("destroy")
		// if e2an.currentSeq != "destroy" {
		// 	fmt.Println("**", e2an.current)
		// 	e2an.currentSeq = "destroy"

		g.playerOne.scores += e2.scoreValue
		if g.playerOne.scores > g.hiScores {
			g.hiScores = g.playerOne.scores
		}
		return
	}
	// playerTwo bullet hit an enemy
	if e1.entityType == typePlayerTwoBullet {
		if !e2.exploding { // bullet pass trough enemy explosion
			e1.active = false // bullet destroyed
		}
		// e2.active = false
		e2.hit = true
		e2.exploding = true
		g.playerTwo.scores += e2.scoreValue
		if g.playerTwo.scores > g.hiScores {
			g.hiScores = g.playerTwo.scores
		}
		return
	}

	// player hit by a bullet or rammed by an enemy
	if e1.entityType == typeEnemyBullet || e1.entityType == typeEnemy {

		if e1.entityType == typeEnemyBullet {
			e1.active = false // bullet ha no idle anim, just disappears
		} else {
			e1.hit = true // trigger idle anim, animator will take care of setting active to false when anim finishes
		}

		// TODO hit logic is patially shared with animator because it takes care of playing the destroy animation
		// TODO once an enemy is hit should not collide but adding && !e1.hit here gives superpowers to player
		if !e2.invulnerable && !e2.hit && !e1.exploding {
			// animator will reset hit and invulnerable flags
			e2.invulnerable = true
			e2.hit = true
			// animator will set active to false if no more lives
			e2.lives -= 1
			// e2an := e1.getComponent(&animator{}).(*animator)
			// e2an.lastFrameChange = time.Now()
			return
		}
	}
}
