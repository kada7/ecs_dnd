/*
 * @author: Haoyuan Liu
 * @date: 2020/10/20
 */

package dice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDice(t *testing.T) {
	min := 1
	max := 6
	d := New(min, max)

	for i := 0; i < 10000; i++ {
		number := d.Roll()
		assert.LessOrEqual(t, min, number, "le min")
		assert.GreaterOrEqual(t, max, number, "ge max")
	}
}
