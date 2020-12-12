/*
 * @author: Haoyuan Liu
 * @date: 2020/11/30
 */

package alignment

import "fmt"

type Trend uint8

const (
	Neutral Trend = iota
	Good
	Evil
	Lawful
	Chaotic
)

func (t Trend) String() string {
	s := "unknown"
	switch t {
	case Neutral:
		s = "neutral"
	case Good:
		s = "good"
	case Evil:
		s = "evil"
	case Lawful:
		s = "lawful"
	case Chaotic:
		s = "chaotic"
	}
	return s
}

type Alignment struct {
	GoodOrEvil      Trend
	LawfulOrChaotic Trend
}

func NewAlignment(lawfulOrChaotic Trend, goodOrEvil Trend) Alignment {
	return Alignment{
		GoodOrEvil:      goodOrEvil,
		LawfulOrChaotic: lawfulOrChaotic,
	}
}

func (a Alignment) IsGood() bool {
	return a.GoodOrEvil == Good
}

func (a Alignment) IsEvil() bool {
	return a.GoodOrEvil == Evil
}

func (a Alignment) IsLawful() bool {
	return a.LawfulOrChaotic == Lawful
}

func (a Alignment) IsChaotic() bool {
	return a.LawfulOrChaotic == Chaotic
}

func (a Alignment) IsValid() bool {
	if !(a.GoodOrEvil == Neutral ||
		a.GoodOrEvil == Good ||
		a.GoodOrEvil == Evil) {

		return false
	}
	if !(a.LawfulOrChaotic == Neutral ||
		a.LawfulOrChaotic == Lawful ||
		a.LawfulOrChaotic == Chaotic) {

		return false
	}
	return true
}

func (a Alignment) String() string {
	return fmt.Sprintf("%s %s", a.LawfulOrChaotic, a.GoodOrEvil)
}
