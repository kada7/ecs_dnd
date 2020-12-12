/*
 * @author: Haoyuan Liu
 * @date: 2020/12/11
 */

package skill

import "fmt"

type Type uint8

const TYPE_COUNT = 17

const (
	Acrobatics Type = iota
	AnimalHandling
	Arcana
	Athletics
	Deception
	History
	Insight
	Intimidation
	Investigation
	Medicine
	Nature
	Perception
	Performance
	Persuasion
	Religion
	SleightOfHand
	Stealth
	Survival
)

var allType = [TYPE_COUNT]Type{
	Acrobatics,
	AnimalHandling,
	Arcana,
	Athletics,
	Deception,
	History,
	Insight,
	Intimidation,
	Investigation,
	Medicine,
	Nature,
	Perception,
	Performance,
	Persuasion,
	Religion,
	SleightOfHand,
	Stealth,
}

func AllType() [TYPE_COUNT]Type {
	return allType
}

func (t Type) IsValid() bool {
	for _, typ := range allType {
		if t == typ {
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

func (t Type) String() string {
	var s = "unknown"
	switch t {
	case Acrobatics:
		s = "acrobatics"
	case AnimalHandling:
		s = "animal_handling"
	case Arcana:
		s = "arcana"
	case Athletics:
		s = "athletics"
	case Deception:
		s = "deception"
	case History:
		s = "history"
	case Insight:
		s = "insight"
	case Intimidation:
		s = "intimidation"
	case Investigation:
		s = "investigation"
	case Medicine:
		s = "medicine"
	case Nature:
		s = "nature"
	case Perception:
		s = "perception"
	case Performance:
		s = "performance"
	case Persuasion:
		s = "persuasion"
	case Religion:
		s = "religion"
	case SleightOfHand:
		s = "sleight_of_hand"
	case Stealth:
		s = "stealth"
	case Survival:
		s = "survival"
	}
	return s
}
