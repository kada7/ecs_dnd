package attr

import "errors"

type Attr interface {
	Total() int
	Base() int
	Bonus() []Bonus
	Override() int
	IsOverride() bool
	Max() int
	Min() int
	SetBase(n int)
	SetOverride(n int)
	SetMin(n int)
	SetMax(n int)
	AddBonus(mods ...Bonus)
	CancelBonus(reason string)
	CancelOverride()
	Clone() Attr
}

type Bonus struct {
	Reason string
	Bonus  int
}

func New(min, max int) (Attr, error) {
	if err := checkMinAndMax(min, max); err != nil {
		return nil, err
	}
	return &attrImpl{
		min:   min,
		max:   max,
		bonus: make([]Bonus, 0),
	}, nil
}

type attrImpl struct {
	total      int
	base       int
	bonus      []Bonus
	override   int
	isOverride bool
	max        int
	min        int
}

func (v *attrImpl) Base() int {
	return v.base
}

func (v *attrImpl) Bonus() []Bonus {
	return v.bonus
}

func (v *attrImpl) Override() int {
	return v.override
}

func (v *attrImpl) IsOverride() bool {
	return v.isOverride
}

func (v *attrImpl) Max() int {
	return v.max
}

func (v *attrImpl) Min() int {
	return v.min
}

func (v *attrImpl) SetBase(n int) {
	v.base = n
	v.resetTotal()
}

func (v *attrImpl) SetOverride(n int) {
	v.isOverride = true
	n = limitInRange(v.Max(), v.Min(), n)
	v.override = n
	v.resetTotal()
}

func (v *attrImpl) CancelOverride() {
	v.isOverride = false
}

func (v *attrImpl) Total() int {
	return v.total
}

func (v *attrImpl) SetMin(n int) {
	if err := checkMinAndMax(n, v.Max()); err != nil {
		panic(err)
	}
	v.min = n
	v.resetTotal()
}

func (v *attrImpl) SetMax(n int) {
	if err := checkMinAndMax(v.Min(), n); err != nil {
		panic(err)
	}
	v.max = n
	v.resetTotal()
}

func (v *attrImpl) AddBonus(mods ...Bonus) {
	v.bonus = append(v.bonus, mods...)
	v.resetTotal()
}

func (v *attrImpl) CancelBonus(reason string) {
	mods := make([]Bonus, 0)
	for _, mod := range v.bonus {
		if mod.Reason == reason {
			continue
		}
		mods = append(mods, mod)
	}
	v.resetTotal()
}

func (v *attrImpl) resetTotal() {
	var n = v.Min()
	if v.IsOverride() {
		n = v.Override()
	} else {
		n = v.Base() + v.bonusTotal()
	}
	v.total = limitInRange(v.Min(), v.Max(), n)
}

func (v *attrImpl) bonusTotal() int {
	n := 0
	for _, b := range v.Bonus() {
		n += b.Bonus
	}
	return n
}

func (v attrImpl) Clone() Attr {
	bs := make([]Bonus, len(v.bonus))
	copy(bs, v.bonus)
	v.bonus = bs
	return &v
}

func checkMinAndMax(min, max int) error {
	if min > max {
		return errors.New("value min must less or equal max")
	}
	return nil
}

func limitInRange(min, max int, n int) int {
	if n > max {
		return n
	}
	if n < min {
		return min
	}
	return n
}
