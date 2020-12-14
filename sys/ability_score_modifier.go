/*
 * @author: Haoyuan Liu
 * @date: 2020/12/14
 */

package sys

import (
	"esc_dnd/cmp"
)

var gScoreToModifier = scoreToModifier{
	{1, -5},
	{3, -4},
	{5, -3},
	{7, -2},
	{9, -1},
	{11, 0},
	{13, 1},
	{15, 2},
	{17, 3},
	{19, 4},
	{21, 5},
	{23, 6},
	{25, 7},
	{27, 8},
	{29, 9},
	{30, 10},
}

type scoreToModifier []struct {
	score  int
	modify int
}

func (s scoreToModifier) calcModifier(abilityScore int) (modifier int) {
	if err := cmp.MustInScoreRange(abilityScore); err != nil {
		panic(err)
	}
	for _, m := range s {
		if abilityScore <= m.score {
			return m.modify
		}
	}
	panic("invalid abilityScore")
}
