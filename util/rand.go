package util

import "math/rand"

// RandRange return a pseudo-random number in [min, max]
func RandRange(r *rand.Rand, min, max int) int {
	if max == min {
		return max
	}
	return r.Intn(max-min) + min + 1
}

// RandRange return a pseudo-random number in [min, max]
func RandRange2(min, max int) int {
	if max == min {
		return max
	}
	return rand.Intn(max-min) + min + 1
}
