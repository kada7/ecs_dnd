/*
 * @author: Haoyuan Liu
 * @date: 2020/10/20
 */

package dice

import (
	"esc_dnd/util"
	"math/rand"
	"sync"
)

// Dice
type Dice struct {
	m        *sync.Mutex
	r        *rand.Rand
	min, max int
}

// New return a *Dice which rolling number in [min, max]
func New(min, max int) *Dice {
	if max < min {
		panic("max must equal or greater than min")
	}
	return &Dice{
		m:   &sync.Mutex{},
		r:   newNowRand(),
		min: min,
		max: max,
	}
}

// Roll return number in [min, max]
func (d Dice) Roll() int {
	d.m.Lock()
	defer d.m.Unlock()
	return util.RandRange(d.r, d.min, d.max)
}

// Max is Dice rolling min number
func (d Dice) Min() int {
	return d.min
}

// Max is Dice rolling max number
func (d Dice) Max() int {
	return d.max
}

func (d *Dice) SetSource(source rand.Source) {
	d.m.Lock()
	defer d.m.Unlock()
	d.r = rand.New(source)
}
