/*
 * @author: Haoyuan Liu
 * @date: 2020/12/11
 */

package cmp

import (
	"esc_dnd/enum/ability"
)

type AbilityListComponent struct {
	abilityList [ability.TYPE_COUNT]*Ability
}

func NewAbilityListComponent(a [ability.TYPE_COUNT]*Ability) *AbilityListComponent {
	return &AbilityListComponent{abilityList: a}
}

func (c AbilityListComponent) Ability(p ability.Type) (*Ability, bool) {
	for _, a := range c.abilityList {
		if a.Type == p {
			return a, true
		}
	}
	return nil, false
}

func (c AbilityListComponent) AllAbility() [ability.TYPE_COUNT]*Ability {
	return c.abilityList
}

func (c AbilityListComponent) Clone() AbilityListComponent {
	return c
}
