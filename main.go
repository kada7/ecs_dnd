package main

import (
	"esc_dnd/cmp"
	"esc_dnd/enum/ability"
	"esc_dnd/sys"
	"fmt"
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
)

func main() {
	world := &ecs.World{}
	engo.Mailbox = &engo.MessageManager{}

	absys := sys.NewAbilitySystem()
	rcsys := sys.NewRaceSystem()
	world.AddSystem(absys)
	world.AddSystem(rcsys)
	ch := CharacterEntity{
		BasicEntity: ecs.NewBasic(),
		AbilityListComponent: cmp.NewAbilityListComponent([ability.TYPE_COUNT]*cmp.Ability{
			cmp.NewAbility(ability.Strength, 10),
			cmp.NewAbility(ability.Charisma, 10),
			cmp.NewAbility(ability.Intelligence, 10),
			cmp.NewAbility(ability.Constitution, 10),
			cmp.NewAbility(ability.Dexterity, 10),
			cmp.NewAbility(ability.Wisdom, 10),
		}),
		RaceComponent: &cmp.RaceComponent{
			Name: "fighter",
			AbilityScoreIncrease: map[ability.Type]int{
				ability.Strength: 1,
			},
		},
	}

	for _, system := range world.Systems() {
		switch s := system.(type) {
		case *sys.AbilitySystem:
			s.Add(&ch.BasicEntity, ch.AbilityListComponent)
		case *sys.RaceSystem:
			s.Add(&ch.BasicEntity, ch.RaceComponent)
		}
	}

	absys.New(nil)
	fmt.Printf("%+v\n", ch.AbilityListComponent)
	ch.RaceComponent = &cmp.RaceComponent{
		Name: "cer",
		AbilityScoreIncrease: map[ability.Type]int{
			ability.Charisma: 1,
		},
	}
	rcsys.ChangeRace(ch.ID(), ch.RaceComponent)
	fmt.Printf("%+v\n", ch.AbilityListComponent)
}

type CharacterEntity struct {
	ecs.BasicEntity
	*cmp.AbilityListComponent
	*cmp.RaceComponent
}
