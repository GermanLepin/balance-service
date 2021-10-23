package random

import (
	"math/rand"
)

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func NumberForCards_db() string {
	number := ""
	for i := 0; i < 16; i++ {
		number += (string)(rand.Intn(10) + 48)
	}
	return number
}
