/*
 * @author: Haoyuan Liu
 * @date: 2020/12/3
 */

package dice

import "fmt"

type DiceCount struct {
	Count int
	// Face is the face number of dice,
	// if Face == 0, dice always roll 0,
	Face  int
	Bonus int
}

func (count DiceCount) Roll() int {
	if count.Face == 0 {
		return 0
	}
	d := NewN(count.Face)
	result := 0
	for i := 0; i < count.Count; i++ {
		result += d.Roll()
	}
	result += count.Bonus
	return result
}

func (count DiceCount) String() string {
	if count.Bonus > 0 {
		return fmt.Sprintf("%dd%d+%d", count.Count, count.Face, count.Bonus)
	}
	return fmt.Sprintf("%dd%d", count.Count, count.Face)
}
