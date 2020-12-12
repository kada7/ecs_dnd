/*
 * @author: Haoyuan Liu
 * @date: 2020/12/11
 */

package sys

import (
	"esc_dnd/cmp"
	"esc_dnd/event"
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
)

type raceEntity struct {
	*ecs.BasicEntity
	*cmp.RaceComponent
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

func (a *RaceSystem) Add(b *ecs.BasicEntity, c *cmp.RaceComponent) {
	a.m[b.ID()] = raceEntity{
		BasicEntity:   b,
		RaceComponent: c,
	}
}

func (a *RaceSystem) ChangeRace(id uint64, newRace *cmp.RaceComponent) {
	v := a.m[id]
	v.RaceComponent = newRace
	a.m[id] = v
	engo.Mailbox.Dispatch(event.RaceChangedMsg{
		BasicEntity: v.BasicEntity,
		NewRace:     *newRace,
	})
}
