package numbers

import (
	"math/rand"
)

func Round(number float32) int {
	return int(number)
}

func RandomWithLowerZero(limit int) int {
	return generateRandom(limit)
}

func RandomWithLowerOne(limit int) int {
	return generateRandom(limit) + 1
}

func generateRandom(limit int) int {
	return rand.Intn(limit)
}
