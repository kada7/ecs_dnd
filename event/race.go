/*
 * @author: Haoyuan Liu
 * @date: 2020/12/11
 */

package event

import (
	"esc_dnd/cmp"
	"github.com/EngoEngine/ecs"
)

const (
	RACE_CHANGED = "race_change"
)

type RaceChangedMsg struct {
	*ecs.BasicEntity
	NewRace cmp.RaceComponent
}

func (a RaceChangedMsg) Type() string {
	return RACE_CHANGED
}
