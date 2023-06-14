package main

import (
	_ "embed"
	"fmt"
	"image"
	"reflect"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	animSuperFlyingMan    = []int{0, 1, 2, 3, 4, 4, 3, 2, 1}                                  // SuperFlyingMan
	animSuperFlyingManPew = []int{5, 6, 7, 8, 7, 6, 5}                                        // SuperFlyingManPew
	animPigPew            = []int{14, 15, 16, 15}                                             // PigPew
	animPig               = []int{10, 11, 12, 13, 13, 12, 11, 10}                             // Pig
	animSuperFlyingManDie = []int{17, 18, 19, 19, 20, 20, 21, 21, 22, 23, 24, 24, 25}         // SuperFlyingManDie
	animPigDie            = []int{26, 27, 28, 28, 29, 29, 30, 30, 31, 31, 22, 23, 24, 24, 25} // PigDie
	animEnemyPew          = []int{47}                                                         // EnemyPew
	animEnemyBaloon       = []int{32, 33, 34, 35, 36, 36, 35, 34, 33, 32}                     // EnemyBaloon
	animEnemyBaloonDie    = []int{37, 38, 38, 39, 39, 40, 40, 46, 47, 48, 49, 49}             // EnemyBaloonDie
	animExplosion         = []int{46, 47, 48, 49, 48, 47, 46}                                 // Explosion
	animEnemyFlyingMan1   = []int{41, 42, 43, 44, 45, 45, 43, 42}                             // EnemyFlyingMan1
	animEnemyThing        = []int{50, 51, 52, 53, 52, 51, 50}                                 // EnemyThing
	animEnemyCat          = []int{54, 55, 56, 57, 56, 55, 54}                                 // EnemyCat
)

type Vector struct {
	x, y float64
}

type Sprite struct {
	image        *ebiten.Image
	animation    []int
	animSpeed    time.Duration
	currentFrame int
	lastFrame    time.Time
}

type Entity struct {
	sprite Sprite

	active   bool
	position Vector
	rotation float64

	components []Component
}

func newEntity(image *ebiten.Image, animation []int, position Vector) *Entity {
	return &Entity{
		sprite:   Sprite{image, animation, time.Millisecond * 67, 0, time.Now()},
		active:   true,
		position: Vector{x: position.x, y: position.y},
		rotation: 0,
	}
}

func (e *Entity) addComponent(c Component) {
	for _, existing := range e.components {
		if reflect.TypeOf(existing) == reflect.TypeOf(c) {
			panic(fmt.Sprintf("Component of type %v already exists", reflect.TypeOf(c)))
		}
	}
	e.components = append(e.components, c)
}

func (e *Entity) getComponent(ofType Component) Component {
	for _, existing := range e.components {
		if reflect.TypeOf(existing) == reflect.TypeOf(ofType) {
			return existing
		}
	}
	panic(fmt.Sprintf("Component of type %v not found", reflect.TypeOf(ofType)))
}

func (e *Entity) Update(g *Game) {
	if !e.active {
		return
	}
	if time.Since(e.sprite.lastFrame) >= e.sprite.animSpeed {
		e.sprite.lastFrame = time.Now()
		e.sprite.currentFrame++
		if e.sprite.currentFrame >= len(e.sprite.animation) {
			e.sprite.currentFrame = 0
		}
	}
	for _, c := range e.components {
		c.Update()
	}
}

func (e *Entity) Draw(screen *ebiten.Image) {
	if !e.active {
		return
	}
	frameOffset := e.sprite.animation[e.sprite.currentFrame] * spriteSize
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(e.position.x, e.position.y)
	screen.DrawImage(e.sprite.image.SubImage(image.Rect(frameOffset, 0, frameOffset+spriteSize, spriteSize)).(*ebiten.Image), opts)
}
