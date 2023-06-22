package main

import (
	_ "embed"
	"fmt"
	"reflect"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vector struct {
	x, y float64
}
type Box struct {
	x, y, w, h float64
}

type entityType int

const (
	typePlayerOne entityType = iota
	typePlayerTwo
	typePlayerOneBullet
	typePlayerTwoBullet
	typeEnemy
	typeEnemyBullet
)

type Entity struct {
	name                string
	active              bool
	position            Vector
	rotation            float64
	lives               int
	scores              int
	scoreValue          int
	hitBoxes            []Box
	components          []Component
	parent              *Entity
	entityType          entityType
	hit                 bool
	exploding           bool
	invulnerable        bool
	invulnerableSetTime time.Time
}

func newEntity(name string, position Vector) *Entity {
	return &Entity{
		name:     name,
		active:   true,
		lives:    0,
		position: Vector{x: position.x, y: position.y},
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
	for _, c := range e.components {
		c.Update()
	}
}

func (e *Entity) Draw(screen *ebiten.Image) {

	if !e.active {
		return
	}
	for _, c := range e.components {
		c.Draw(screen)
	}
}
