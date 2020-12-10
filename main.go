package main

import (
	"fmt"
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
)

const (
	abilityScoreTotalUpdated = "ability_score_total_updated"
)

type abilityScoreTotalUpdatedMessage struct {
	*ecs.BasicEntity
	Before int
	After  int
}

func (a abilityScoreTotalUpdatedMessage) Type() string {
	return abilityScoreTotalUpdated
}

const (
	raceChanged = "raceChange"
)

type raceChangedData struct {
	*ecs.BasicEntity
	After *RaceComponent
}

func (a raceChangedData) Type() string {
	return raceChanged
}

func main() {
	world := &ecs.World{}
	engo.Mailbox = &engo.MessageManager{}

	absys := &AbilitySystem{map[uint64]abilityEntity{}}
	rcsys := &RaceSystem{map[uint64]raceEntity{}}
	world.AddSystem(absys)
	world.AddSystem(rcsys)
	ch := CharacterEntity{
		BasicEntity: ecs.NewBasic(),
		AbilityComponent: AbilityComponent{
			AbilityScoreTotal: 10,
			AbilityScoreBase:  10,
			AbilityScoreBonus: map[string]Bonus{
				"race_bonus": {
					Bonus:  1,
					Reason: "race_bonus",
				},
			},
			AbilityScoreModifier: 5,
		},
		RaceComponent: RaceComponent{
			Name:                 "fighter",
			AbilityScoreIncrease: 1,
		},
	}

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *AbilitySystem:
			sys.Add(&ch.BasicEntity, &ch.AbilityComponent)
		case *RaceSystem:
			sys.Add(&ch.BasicEntity, &ch.RaceComponent)
		}
	}

	absys.New(nil)
	fmt.Printf("%+v\n", ch.AbilityComponent)
	ch.RaceComponent = RaceComponent{
		Name:                 "cer",
		AbilityScoreIncrease: 2,
	}
	rcsys.ChangeRace(ch.ID(), &ch.RaceComponent)
	fmt.Printf("%+v\n", ch.AbilityComponent)
}

type CharacterEntity struct {
	ecs.BasicEntity
	//CreatureComponent
	//HealthComponent
	//MoveComponent
	AbilityComponent
	RaceComponent
}

type AbilityComponent struct {
	AbilityScoreTotal    int
	AbilityScoreBase     int
	AbilityScoreBonus    map[string]Bonus
	AbilityScoreModifier int
}

type Bonus struct {
	Bonus  int
	Reason string
}

type RaceComponent struct {
	Name                 string
	AbilityScoreIncrease int
}

type abilityEntity struct {
	*ecs.BasicEntity
	*AbilityComponent
}

type AbilitySystem struct {
	m map[uint64]abilityEntity
}

func (a AbilitySystem) New(world *ecs.World) {
	engo.Mailbox.Listen(raceChanged, func(msg engo.Message) {
		m, ok := msg.(raceChangedData)
		if !ok {
			return
		}
		e := a.m[m.ID()]
		newRace := m.After
		bonus := e.AbilityScoreBonus
		bonus["race_bonus"] = Bonus{newRace.AbilityScoreIncrease, "race_bonus"}
	})
}

func (a AbilitySystem) Update(dt float32) {
	panic("implement me")
}

func (a AbilitySystem) Remove(e ecs.BasicEntity) {
	panic("implement me")
}

func (a *AbilitySystem) Add(b *ecs.BasicEntity, c *AbilityComponent) {
	a.m[b.ID()] = abilityEntity{
		BasicEntity:      b,
		AbilityComponent: c,
	}
}

type raceEntity struct {
	*ecs.BasicEntity
	*RaceComponent
}

type RaceSystem struct {
	m map[uint64]raceEntity
}

func (a RaceSystem) Update(dt float32) {
	panic("implement me")
}

func (a RaceSystem) Remove(e ecs.BasicEntity) {
	panic("implement me")
}

func (a *RaceSystem) Add(b *ecs.BasicEntity, c *RaceComponent) {
	a.m[b.ID()] = raceEntity{
		BasicEntity:   b,
		RaceComponent: c,
	}
}

func (a *RaceSystem) ChangeRace(id uint64, newRace *RaceComponent) {
	v := a.m[id]
	v.RaceComponent = newRace
	a.m[id] = v
	engo.Mailbox.Dispatch(raceChangedData{
		BasicEntity: v.BasicEntity,
		After:       newRace,
	})
}
