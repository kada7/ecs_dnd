/*
 * @author: Haoyuan Liu
 * @date: 2020/10/20
 */

package dice

var (
	d2   = NewN(2)
	d4   = NewN(4)
	d6   = NewN(6)
	d8   = NewN(8)
	d12  = NewN(12)
	d20  = NewN(20)
	d100 = NewN(100)
)

// NewN return a *Dice which is in [1, n]
// if you want custom Dice, please use New
func NewN(n int) *Dice {
	return New(1, n)
}

func RollD2() int {
	return d2.Roll()
}

func RollD4() int {
	return d4.Roll()
}

func RollD6() int {
	return d6.Roll()
}

func RollD8() int {
	return d8.Roll()
}

func RollD12() int {
	return d12.Roll()
}

func RollD20() int {
	return d20.Roll()
}

func RollD100() int {
	return d100.Roll()
}

func RollAdvantage(dice *Dice) (advantage, disadvantage int) {
	a := dice.Roll()
	b := dice.Roll()
	if a > b {
		return a, b
	} else {
		return b, a
	}
}
