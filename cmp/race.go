/*
 * @author: Haoyuan Liu
 * @date: 2020/12/11
 */

package cmp

import "esc_dnd/enum/ability"

type RaceComponent struct {
	Name                 string
	AbilityScoreIncrease map[ability.Type]int
}
