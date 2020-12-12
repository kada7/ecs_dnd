/*
 * @author: Haoyuan Liu
 * @date: 2020/12/11
 */

package sys

import (
	"esc_dnd/cmp"
	"esc_dnd/common/attr"
	"esc_dnd/event"
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
)

type abilityEntity struct {
	*ecs.BasicEntity
	*cmp.AbilityListComponent
}

type AbilitySystem struct {
	m map[uint64]abilityEntity
}

func (a AbilitySystem) New(world *ecs.World) {
	engo.Mailbox.Listen(event.RACE_CHANGED, func(msg engo.Message) {
		m, ok := msg.(event.RaceChangedMsg)
		if !ok {
			return
		}
		e := a.m[m.ID()]
		for i, abi := range e.abilityList {
			incr, ok := m.NewRace.AbilityScoreIncrease[abi.typ]
			if !ok {
				continue
			}
			e.abilityList[i].score.AddBonus(attr.Bonus{
				Reason: "racial_bonus",
				Bonus:  incr,
			})
		}
	})
}

func (a AbilitySystem) Update(dt float32) {
	panic("implement me")
}

func (a AbilitySystem) Remove(e ecs.BasicEntity) {
	panic("implement me")
}

func (a *AbilitySystem) Add(b *ecs.BasicEntity, c *cmp.AbilityListComponent) {
	a.m[b.ID()] = abilityEntity{
		BasicEntity:          b,
		AbilityListComponent: c,
	}
}
