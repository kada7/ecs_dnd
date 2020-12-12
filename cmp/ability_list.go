/*
 * @author: Haoyuan Liu
 * @date: 2020/12/11
 */

package cmp

import (
	"esc_dnd/enum/ability"
)

type AbilityListComponent struct {
	abilityList [ability.TYPE_COUNT]Ability
}

func NewAbilityListComponent() AbilityListComponent {
	return AbilityListComponent{abilityList: [ability.TYPE_COUNT]Ability{}}
}

func (c AbilityListComponent) Ability(p ability.Type) (Ability, bool) {
	for _, a := range c.abilityList {
		if a.typ == p {
			return a, true
		}
	}
	return Ability{}, false
}

func (c AbilityListComponent) Clone() AbilityListComponent {
	return c
}
