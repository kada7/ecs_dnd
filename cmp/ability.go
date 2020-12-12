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
	typ           ability.Type
	score         attr.Attr
	scoreModifier int
}

func NewAbility(p ability.Type, baseScore int) (*Ability, error) {
	if err := MustInScoreRange(baseScore); err != nil {
		return nil, fmt.Errorf("cmp.NewAbility: %w", err)
	}
	v, err := attr.New(MIN_ABILITY_SCORE, MAX_ABILITY_SCORE)
	if err != nil {
		return nil, fmt.Errorf("cmp.NewAbility: %w", err)
	}
	v.SetBase(baseScore)
	a := &Ability{}
	return a, nil
}

func (a *Ability) ScoreModifier() int {
	return a.scoreModifier
}

func (a *Ability) SetScoreModifier(scoreModifier int) {
	a.scoreModifier = scoreModifier
}

func (a *Ability) Score() attr.Attr {
	return a.score
}

func (a *Ability) SetScore(score attr.Attr) {
	a.score = score
}

func (a *Ability) resetScoreModifier() {

}

func (a *Ability) Typ() ability.Type {
	return a.typ
}

func (a *Ability) SetTyp(typ ability.Type) {
	a.typ = typ
}

func (a Ability) Clone() Ability {
	a.score = a.score.Clone()
	return a
}

func MustInScoreRange(abilityScore int) error {
	if abilityScore < MIN_ABILITY_SCORE || abilityScore > MAX_ABILITY_SCORE {
		return errors.New("abilityScore MUST in range [MIN_ABILITY_SCORE, MAX_ABILITY_SCORE]")
	}
	return nil
}
