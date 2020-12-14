/*
 * @author: Haoyuan Liu
 * @date: 2020/12/11
 */

package cmp

import (
	"errors"
	"esc_dnd/common/attr"
	"esc_dnd/enum/ability"
	"fmt"
)

const (
	MIN_ABILITY_SCORE = 1
	MAX_ABILITY_SCORE = 30
)

type Ability struct {
	Type          ability.Type
	Score         attr.Attr
	ScoreModifier int
}

func NewAbility(p ability.Type, baseScore int) *Ability {
	if err := MustInScoreRange(baseScore); err != nil {
		panic(fmt.Errorf("cmp.NewAbility: %w", err))
	}
	v, err := attr.New(MIN_ABILITY_SCORE, MAX_ABILITY_SCORE)
	if err != nil {
		panic(fmt.Errorf("cmp.NewAbility: %w", err))
	}
	v.SetBase(baseScore)
	return a
}

func (a Ability) Clone() Ability {
	a.Score = a.Score.Clone()
	return a
}

func MustInScoreRange(abilityScore int) error {
	if abilityScore < MIN_ABILITY_SCORE || abilityScore > MAX_ABILITY_SCORE {
		return errors.New("abilityScore MUST in range [MIN_ABILITY_SCORE, MAX_ABILITY_SCORE]")
	}
	return nil
}
