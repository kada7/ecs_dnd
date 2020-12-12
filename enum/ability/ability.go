/*
 * @author: Haoyuan Liu
 * @date: 2020/12/11
 */

package ability

import (
	"esc_dnd/enum/skill"
	"fmt"
)

type Type uint8

const TYPE_COUNT = 6

const (
	Strength Type = iota
	Dexterity
	Constitution
	Intelligence
	Wisdom
	Charisma
)

var allType = [TYPE_COUNT]Type{
	Strength,
	Dexterity,
	Constitution,
	Intelligence,
	Wisdom,
	Charisma,
}

func AllTypes() [TYPE_COUNT]Type {
	return allType
}

func (t Type) IsValid() bool {
	for _, v := range allType {
		if v == t {
			return true
		}
	}
	return false
}

func (t Type) MustValid() error {
	if !t.IsValid() {
		return fmt.Errorf("Type[%d] is invalid", t)
	}
	return nil
}

func (t Type) SkillTypes() []skill.Type {
	skillTypes, ok := gAbilityToSkill[t]
	if !ok {
		skillTypes = []skill.Type{}
	}
	return skillTypes
}

func (t Type) String() string {
	var s = "unknown"
	switch t {
	case Strength:
		s = "strength"
	case Dexterity:
		s = "dexterity"
	case Constitution:
		s = "constitution"
	case Intelligence:
		s = "intelligence"
	case Wisdom:
		s = "wisdom"
	case Charisma:
		s = "charisma"
	}
	return s
}

func SkillToAbility(p skill.Type) Type {
	t, ok := gSkillToAbility[p]
	if !ok {
		panic(fmt.Sprintf("Type[%d] is invalid", p))
	}
	return t
}

var gAbilityToSkill = map[Type][]skill.Type{
	Strength: {
		skill.Athletics,
	},
	Dexterity: {
		skill.Acrobatics,
		skill.SleightOfHand,
		skill.Stealth,
	},
	Intelligence: {
		skill.Arcana,
		skill.History,
		skill.Medicine,
		skill.Religion,
	},
	Wisdom: {
		skill.AnimalHandling,
		skill.Insight,
		skill.Nature,
		skill.Perception,
		skill.Survival,
	},
	Charisma: {
		skill.Deception,
		skill.Intimidation,
		skill.Investigation,
		skill.Performance,
		skill.Persuasion,
	},
}

var gSkillToAbility map[skill.Type]Type

func init() {
	gSkillToAbility = map[skill.Type]Type{}
	for typ, list := range gAbilityToSkill {
		for _, skillTyp := range list {
			gSkillToAbility[skillTyp] = typ
		}
	}
}
