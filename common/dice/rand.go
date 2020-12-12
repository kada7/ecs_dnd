/*
 * @author: Haoyuan Liu
 * @date: 2020/10/20
 */

package dice

import (
	"math/rand"
	"time"
)

func newNowRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
