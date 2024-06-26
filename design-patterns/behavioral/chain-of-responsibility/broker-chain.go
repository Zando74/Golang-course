package main

import (
	"fmt"
	"sync"
)

// cqs, mediator, cor

type Argument int

const (
	Attack Argument = iota
	Defense
)

type Query struct {
	Creature2Name string
	WhatToQuery   Argument
	Value         int
}

type Observer interface {
	Handle(*Query)
}

type Observable interface {
	Subscribe(o Observer)
	Unsubscribe(o Observer)
	Fire(q *Query)
}

type Game struct {
	observers sync.Map
}

func (g *Game) Subscribe(o Observer) {
	g.observers.Store(o, struct{}{})
	//                   ↑↑↑ empty anon struct
}

func (g *Game) Unsubscribe(o Observer) {
	g.observers.Delete(o)
}

func (g *Game) Fire(q *Query) {
	g.observers.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(Observer).Handle(q)
		return true
	})
}

type Creature2 struct {
	game            *Game
	Name            string
	attack, defense int // ← private!
}

func NewCreature2(game *Game, name string, attack int, defense int) *Creature2 {
	return &Creature2{game: game, Name: name, attack: attack, defense: defense}
}

func (c *Creature2) Attack() int {
	q := Query{c.Name, Attack, c.attack}
	c.game.Fire(&q)
	return q.Value
}

func (c *Creature2) Defense() int {
	q := Query{c.Name, Defense, c.defense}
	c.game.Fire(&q)
	return q.Value
}

func (c *Creature2) String() string {
	return fmt.Sprintf("%s (%d/%d)",
		c.Name, c.Attack(), c.Defense())
}

// data common to all modifiers
type Creature2Modifier struct {
	game      *Game
	creature2 *Creature2
}

func (c *Creature2Modifier) Handle(*Query) {
	// nothing here!
}

type DoubleAttackModifier2 struct {
	Creature2Modifier
}

func NewDoubleAttackModifier2(g *Game, c *Creature2) *DoubleAttackModifier2 {
	d := &DoubleAttackModifier2{Creature2Modifier{g, c}}
	g.Subscribe(d)
	return d
}

func (d *DoubleAttackModifier2) Handle(q *Query) {
	if q.Creature2Name == d.creature2.Name &&
		q.WhatToQuery == Attack {
		q.Value *= 2
	}
}

func (d *DoubleAttackModifier2) Close() error {
	d.game.Unsubscribe(d)
	return nil
}

func main2() {
	game := &Game{sync.Map{}}
	goblin := NewCreature2(game, "Strong Goblin", 2, 2)
	fmt.Println(goblin.String())

	{
		m := NewDoubleAttackModifier2(game, goblin)
		fmt.Println(goblin.String())
		m.Close()
	}

	fmt.Println(goblin.String())
}
