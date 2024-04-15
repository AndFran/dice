package dice 

import (
	"math/rand"
	"time"
)


func Seed(n int64) {
	if n == 0 {
		rand.Seed(time.Now().UnixNano())
	}
	rand.Seed(n)
}

func Roll(sides int) int {
	return rand.Intn(sides) + 1
}
