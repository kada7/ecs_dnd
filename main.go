package main

import (
	"fmt"
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
)

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

type Bonus struct {
	Bonus  int
	Reason string
}
