package main

import (
	_ "embed"
	"fmt"
	"reflect"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vector struct {
	X, Y float64
}
type Box struct {
	X, Y, W, H float64
}

type EntityType int

const (
	TypePlayerOne EntityType = iota
	TypePlayerTwo
	TypePlayerOneBullet
	TypePlayerTwoBullet
	TypeEnemy
	TypeEnemyBullet
)

type Entity struct {
	Name                string
	Active              bool
	Position            Vector
	Rotation            float64
	Lives               int
	Scores              int
	ScoreValue          int
	HitBoxes            []Box
	components          []Component
	Parent              *Entity
	EntityType          EntityType
	Hit                 bool
	Exploding           bool
	Invulnerable        bool
	InvulnerableSetTime time.Time
}

func NewEntity(name string, position Vector) *Entity {
	return &Entity{
		Name:     name,
		Active:   true,
		Lives:    0,
		Position: Vector{X: position.X, Y: position.Y},
	}
}

func (e *Entity) AddComponent(c Component) {
	for _, existing := range e.components {
		if reflect.TypeOf(existing) == reflect.TypeOf(c) {
			panic(fmt.Sprintf("Component of type %v already exists", reflect.TypeOf(c)))
		}
	}
	e.components = append(e.components, c)
}

func (e *Entity) GetComponent(ofType Component) Component {
	for _, existing := range e.components {
		if reflect.TypeOf(existing) == reflect.TypeOf(ofType) {
			return existing
		}
	}
	panic(fmt.Sprintf("Component of type %v not found", reflect.TypeOf(ofType)))
}

func (e *Entity) Update() {
	if !e.Active {
		return
	}
	for _, c := range e.components {
		c.Update()
	}
}

func (e *Entity) Draw(screen *ebiten.Image) {

	if !e.Active {
		return
	}
	for _, c := range e.components {
		c.Draw(screen)
	}
}
